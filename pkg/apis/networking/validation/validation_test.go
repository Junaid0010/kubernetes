/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validation

import (
	"fmt"
	"strings"
	"testing"
	"time"

	apimachineryvalidation "k8s.io/apimachinery/pkg/api/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/networking"
	utilpointer "k8s.io/utils/pointer"
)

func makeValidNetworkPolicy() *networking.NetworkPolicy {
	return &networking.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
		Spec: networking.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{"a": "b"},
			},
		},
	}
}

type netpolTweak func(networkPolicy *networking.NetworkPolicy)

func makeNetworkPolicyCustom(tweaks ...netpolTweak) *networking.NetworkPolicy {
	networkPolicy := makeValidNetworkPolicy()
	for _, fn := range tweaks {
		fn(networkPolicy)
	}
	return networkPolicy
}

func makePort(proto *api.Protocol, port intstr.IntOrString, endPort int32) networking.NetworkPolicyPort {
	r := networking.NetworkPolicyPort{
		Protocol: proto,
		Port:     nil,
	}
	if port != intstr.FromInt(0) && port != intstr.FromString("") && port != intstr.FromString("0") {
		r.Port = &port
	}
	if endPort != 0 {
		r.EndPort = utilpointer.Int32Ptr(endPort)
	}
	return r
}

func TestValidateNetworkPolicy(t *testing.T) {
	protocolTCP := api.ProtocolTCP
	protocolUDP := api.ProtocolUDP
	protocolICMP := api.Protocol("ICMP")
	protocolSCTP := api.ProtocolSCTP

	// Tweaks used below.
	setIngressEmptyFirstElement := func(networkPolicy *networking.NetworkPolicy) {
		networkPolicy.Spec.Ingress = []networking.NetworkPolicyIngressRule{{}}
	}

	setIngressFromEmptyFirstElement := func(networkPolicy *networking.NetworkPolicy) {
		if networkPolicy.Spec.Ingress == nil {
			setIngressEmptyFirstElement(networkPolicy)
		}
		networkPolicy.Spec.Ingress[0].From = []networking.NetworkPolicyPeer{{}}
	}

	setIngressFromIfEmpty := func(networkPolicy *networking.NetworkPolicy) {
		if networkPolicy.Spec.Ingress == nil {
			setIngressEmptyFirstElement(networkPolicy)
		}
		if networkPolicy.Spec.Ingress[0].From == nil {
			setIngressFromEmptyFirstElement(networkPolicy)
		}
	}

	setIngressEmptyPorts := func(networkPolicy *networking.NetworkPolicy) {
		networkPolicy.Spec.Ingress = []networking.NetworkPolicyIngressRule{{
			Ports: []networking.NetworkPolicyPort{{}},
		}}
	}

	setIngressPorts := func(ports ...networking.NetworkPolicyPort) netpolTweak {
		return func(np *networking.NetworkPolicy) {
			if np.Spec.Ingress == nil {
				setIngressEmptyFirstElement(np)
			}
			np.Spec.Ingress[0].Ports = make([]networking.NetworkPolicyPort, len(ports))
			for i, p := range ports {
				np.Spec.Ingress[0].Ports[i] = p
			}
		}
	}

	setIngressFromPodSelector := func(k, v string) func(*networking.NetworkPolicy) {
		return func(networkPolicy *networking.NetworkPolicy) {
			setIngressFromIfEmpty(networkPolicy)
			networkPolicy.Spec.Ingress[0].From[0].PodSelector = &metav1.LabelSelector{
				MatchLabels: map[string]string{k: v},
			}
		}
	}

	setIngressFromNamespaceSelector := func(networkPolicy *networking.NetworkPolicy) {
		setIngressFromIfEmpty(networkPolicy)
		networkPolicy.Spec.Ingress[0].From[0].NamespaceSelector = &metav1.LabelSelector{
			MatchLabels: map[string]string{"c": "d"},
		}
	}

	setIngressFromIPBlockIPV4 := func(networkPolicy *networking.NetworkPolicy) {
		setIngressFromIfEmpty(networkPolicy)
		networkPolicy.Spec.Ingress[0].From[0].IPBlock = &networking.IPBlock{
			CIDR:   "192.168.0.0/16",
			Except: []string{"192.168.3.0/24", "192.168.4.0/24"},
		}
	}

	setIngressFromIPBlockIPV6 := func(networkPolicy *networking.NetworkPolicy) {
		setIngressFromIfEmpty(networkPolicy)
		networkPolicy.Spec.Ingress[0].From[0].IPBlock = &networking.IPBlock{
			CIDR:   "fd00:192:168::/48",
			Except: []string{"fd00:192:168:3::/64", "fd00:192:168:4::/64"},
		}
	}

	setEgressEmptyFirstElement := func(networkPolicy *networking.NetworkPolicy) {
		networkPolicy.Spec.Egress = []networking.NetworkPolicyEgressRule{{}}
	}

	setEgressToEmptyFirstElement := func(networkPolicy *networking.NetworkPolicy) {
		if networkPolicy.Spec.Egress == nil {
			setEgressEmptyFirstElement(networkPolicy)
		}
		networkPolicy.Spec.Egress[0].To = []networking.NetworkPolicyPeer{{}}
	}

	setEgressToIfEmpty := func(networkPolicy *networking.NetworkPolicy) {
		if networkPolicy.Spec.Egress == nil {
			setEgressEmptyFirstElement(networkPolicy)
		}
		if networkPolicy.Spec.Egress[0].To == nil {
			setEgressToEmptyFirstElement(networkPolicy)
		}
	}

	setEgressToNamespaceSelector := func(networkPolicy *networking.NetworkPolicy) {
		setEgressToIfEmpty(networkPolicy)
		networkPolicy.Spec.Egress[0].To[0].NamespaceSelector = &metav1.LabelSelector{
			MatchLabels: map[string]string{"c": "d"},
		}
	}

	setEgressToPodSelector := func(networkPolicy *networking.NetworkPolicy) {
		setEgressToIfEmpty(networkPolicy)
		networkPolicy.Spec.Egress[0].To[0].PodSelector = &metav1.LabelSelector{
			MatchLabels: map[string]string{"c": "d"},
		}
	}

	setEgressToIPBlockIPV4 := func(networkPolicy *networking.NetworkPolicy) {
		setEgressToIfEmpty(networkPolicy)
		networkPolicy.Spec.Egress[0].To[0].IPBlock = &networking.IPBlock{
			CIDR:   "192.168.0.0/16",
			Except: []string{"192.168.3.0/24", "192.168.4.0/24"},
		}
	}

	setEgressToIPBlockIPV6 := func(networkPolicy *networking.NetworkPolicy) {
		setEgressToIfEmpty(networkPolicy)
		networkPolicy.Spec.Egress[0].To[0].IPBlock = &networking.IPBlock{
			CIDR:   "fd00:192:168::/48",
			Except: []string{"fd00:192:168:3::/64", "fd00:192:168:4::/64"},
		}
	}

	setEgressPorts := func(ports ...networking.NetworkPolicyPort) netpolTweak {
		return func(np *networking.NetworkPolicy) {
			if np.Spec.Egress == nil {
				setEgressEmptyFirstElement(np)
			}
			np.Spec.Egress[0].Ports = make([]networking.NetworkPolicyPort, len(ports))
			for i, p := range ports {
				np.Spec.Egress[0].Ports[i] = p
			}
		}
	}

	setPolicyTypesEgress := func(networkPolicy *networking.NetworkPolicy) {
		networkPolicy.Spec.PolicyTypes = []networking.PolicyType{networking.PolicyTypeEgress}
	}

	setPolicyTypesIngressEgress := func(networkPolicy *networking.NetworkPolicy) {
		networkPolicy.Spec.PolicyTypes = []networking.PolicyType{networking.PolicyTypeIngress, networking.PolicyTypeEgress}
	}

	successCases := []*networking.NetworkPolicy{
		makeNetworkPolicyCustom(setIngressEmptyFirstElement),
		makeNetworkPolicyCustom(setIngressFromEmptyFirstElement, setIngressEmptyPorts),
		makeNetworkPolicyCustom(setIngressPorts(
			makePort(nil, intstr.FromInt(80), 0),
			makePort(&protocolTCP, intstr.FromInt(0), 0),
			makePort(&protocolTCP, intstr.FromInt(443), 0),
			makePort(&protocolUDP, intstr.FromString("dns"), 0),
			makePort(&protocolSCTP, intstr.FromInt(7777), 0),
		)),
		makeNetworkPolicyCustom(setIngressFromPodSelector("c", "d")),
		makeNetworkPolicyCustom(setIngressFromNamespaceSelector),
		makeNetworkPolicyCustom(setIngressFromPodSelector("e", "f"), setIngressFromNamespaceSelector),
		makeNetworkPolicyCustom(setEgressToNamespaceSelector, setIngressFromIPBlockIPV4),
		makeNetworkPolicyCustom(setIngressFromIPBlockIPV4),
		makeNetworkPolicyCustom(setEgressToIPBlockIPV4, setPolicyTypesEgress),
		makeNetworkPolicyCustom(setEgressToIPBlockIPV4, setPolicyTypesIngressEgress),
		makeNetworkPolicyCustom(setEgressPorts(
			makePort(nil, intstr.FromInt(80), 0),
			makePort(&protocolTCP, intstr.FromInt(0), 0),
			makePort(&protocolTCP, intstr.FromInt(443), 0),
			makePort(&protocolUDP, intstr.FromString("dns"), 0),
			makePort(&protocolSCTP, intstr.FromInt(7777), 0),
		)),
		makeNetworkPolicyCustom(setEgressToNamespaceSelector, setIngressFromIPBlockIPV6),
		makeNetworkPolicyCustom(setIngressFromIPBlockIPV6),
		makeNetworkPolicyCustom(setEgressToIPBlockIPV6, setPolicyTypesEgress),
		makeNetworkPolicyCustom(setEgressToIPBlockIPV6, setPolicyTypesIngressEgress),
		makeNetworkPolicyCustom(setEgressPorts(makePort(nil, intstr.FromInt(32000), 32768), makePort(&protocolUDP, intstr.FromString("dns"), 0))),
		makeNetworkPolicyCustom(
			setEgressToNamespaceSelector,
			setEgressPorts(
				makePort(nil, intstr.FromInt(30000), 32768),
				makePort(nil, intstr.FromInt(32000), 32768),
			),
			setIngressFromPodSelector("e", "f"),
			setIngressPorts(makePort(&protocolTCP, intstr.FromInt(32768), 32768))),
	}

	// Success cases are expected to pass validation.

	for k, v := range successCases {
		if errs := ValidateNetworkPolicy(v); len(errs) != 0 {
			t.Errorf("Expected success for the success validation test number %d, got %v", k, errs)
		}
	}

	invalidSelector := map[string]string{"NoUppercaseOrSpecialCharsLike=Equals": "b"}

	errorCases := map[string]*networking.NetworkPolicy{
		"namespaceSelector and ipBlock": makeNetworkPolicyCustom(setIngressFromNamespaceSelector, setIngressFromIPBlockIPV4),
		"podSelector and ipBlock":       makeNetworkPolicyCustom(setEgressToPodSelector, setEgressToIPBlockIPV4),
		"missing from and to type":      makeNetworkPolicyCustom(setIngressFromEmptyFirstElement, setEgressToEmptyFirstElement),
		"invalid spec.podSelector": makeNetworkPolicyCustom(setIngressFromNamespaceSelector, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec = networking.NetworkPolicySpec{
				PodSelector: metav1.LabelSelector{
					MatchLabels: invalidSelector,
				},
			}
		}),
		"invalid ingress.ports.protocol":   makeNetworkPolicyCustom(setIngressPorts(makePort(&protocolICMP, intstr.FromInt(80), 0))),
		"invalid ingress.ports.port (int)": makeNetworkPolicyCustom(setIngressPorts(makePort(&protocolTCP, intstr.FromInt(123456789), 0))),
		"invalid ingress.ports.port (str)": makeNetworkPolicyCustom(
			setIngressPorts(makePort(&protocolTCP, intstr.FromString("!@#$"), 0))),
		"invalid ingress.from.podSelector": makeNetworkPolicyCustom(setIngressFromEmptyFirstElement, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].PodSelector = &metav1.LabelSelector{
				MatchLabels: invalidSelector,
			}
		}),
		"invalid egress.to.podSelector": makeNetworkPolicyCustom(setEgressToEmptyFirstElement, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Egress[0].To[0].PodSelector = &metav1.LabelSelector{
				MatchLabels: invalidSelector,
			}
		}),
		"invalid egress.ports.protocol":   makeNetworkPolicyCustom(setEgressPorts(makePort(&protocolICMP, intstr.FromInt(80), 0))),
		"invalid egress.ports.port (int)": makeNetworkPolicyCustom(setEgressPorts(makePort(&protocolTCP, intstr.FromInt(123456789), 0))),
		"invalid egress.ports.port (str)": makeNetworkPolicyCustom(setEgressPorts(makePort(&protocolTCP, intstr.FromString("!@#$"), 0))),
		"invalid ingress.from.namespaceSelector": makeNetworkPolicyCustom(setIngressFromEmptyFirstElement, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].NamespaceSelector = &metav1.LabelSelector{
				MatchLabels: invalidSelector,
			}
		}),
		"missing cidr field": makeNetworkPolicyCustom(setIngressFromIPBlockIPV4, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock.CIDR = ""
		}),
		"invalid cidr format": makeNetworkPolicyCustom(setIngressFromIPBlockIPV4, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock.CIDR = "192.168.5.6"
		}),
		"invalid ipv6 cidr format": makeNetworkPolicyCustom(setIngressFromIPBlockIPV6, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock.CIDR = "fd00:192:168::"
		}),
		"except field is an empty string": makeNetworkPolicyCustom(setIngressFromIPBlockIPV4, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock.Except = []string{""}
		}),
		"except field is an space string": makeNetworkPolicyCustom(setIngressFromIPBlockIPV4, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock.Except = []string{" "}
		}),
		"except field is an invalid ip": makeNetworkPolicyCustom(setIngressFromIPBlockIPV4, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock.Except = []string{"300.300.300.300"}
		}),
		"except IP is outside of CIDR range": makeNetworkPolicyCustom(setIngressFromEmptyFirstElement, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock = &networking.IPBlock{
				CIDR:   "192.168.8.0/24",
				Except: []string{"192.168.9.1/24"},
			}
		}),
		"except IP is not strictly within CIDR range": makeNetworkPolicyCustom(setIngressFromEmptyFirstElement, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock = &networking.IPBlock{
				CIDR:   "192.168.0.0/24",
				Except: []string{"192.168.0.0/24"},
			}
		}),
		"except IPv6 is outside of CIDR range": makeNetworkPolicyCustom(setIngressFromEmptyFirstElement, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.Ingress[0].From[0].IPBlock = &networking.IPBlock{
				CIDR:   "fd00:192:168:1::/64",
				Except: []string{"fd00:192:168:2::/64"},
			}
		}),
		"invalid policyTypes": makeNetworkPolicyCustom(setEgressToIPBlockIPV4, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.PolicyTypes = []networking.PolicyType{"foo", "bar"}
		}),
		"too many policyTypes": makeNetworkPolicyCustom(setEgressToIPBlockIPV4, func(networkPolicy *networking.NetworkPolicy) {
			networkPolicy.Spec.PolicyTypes = []networking.PolicyType{"foo", "bar", "baz"}
		}),
		"multiple ports defined, one port range is invalid": makeNetworkPolicyCustom(
			setEgressToNamespaceSelector,
			setEgressPorts(
				makePort(&protocolUDP, intstr.FromInt(35000), 32768),
				makePort(nil, intstr.FromInt(32000), 32768),
			),
		),
		"endPort defined with named/string port": makeNetworkPolicyCustom(
			setEgressToNamespaceSelector,
			setEgressPorts(
				makePort(&protocolUDP, intstr.FromString("dns"), 32768),
				makePort(nil, intstr.FromInt(32000), 32768),
			),
		),
		"endPort defined without port defined": makeNetworkPolicyCustom(
			setEgressToNamespaceSelector,
			setEgressPorts(makePort(&protocolTCP, intstr.FromInt(0), 32768)),
		),
		"port is greater than endPort": makeNetworkPolicyCustom(
			setEgressToNamespaceSelector,
			setEgressPorts(makePort(&protocolSCTP, intstr.FromInt(35000), 32768)),
		),
		"multiple invalid port ranges defined": makeNetworkPolicyCustom(
			setEgressToNamespaceSelector,
			setEgressPorts(
				makePort(&protocolUDP, intstr.FromInt(35000), 32768),
				makePort(&protocolTCP, intstr.FromInt(0), 32768),
				makePort(&protocolTCP, intstr.FromString("https"), 32768),
			),
		),
		"invalid endport range defined": makeNetworkPolicyCustom(setEgressToNamespaceSelector, setEgressPorts(makePort(&protocolTCP, intstr.FromInt(30000), 65537))),
	}

	// Error cases are not expected to pass validation.
	for testName, networkPolicy := range errorCases {
		if errs := ValidateNetworkPolicy(networkPolicy); len(errs) == 0 {
			t.Errorf("Expected failure for test: %s", testName)
		}
	}
}

func TestValidateNetworkPolicyUpdate(t *testing.T) {
	type npUpdateTest struct {
		old    networking.NetworkPolicy
		update networking.NetworkPolicy
	}
	successCases := map[string]npUpdateTest{
		"no change": {
			old: networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PodSelector: metav1.LabelSelector{
						MatchLabels: map[string]string{"a": "b"},
					},
					Ingress: []networking.NetworkPolicyIngressRule{},
				},
			},
			update: networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PodSelector: metav1.LabelSelector{
						MatchLabels: map[string]string{"a": "b"},
					},
					Ingress: []networking.NetworkPolicyIngressRule{},
				},
			},
		},
		"change spec": {
			old: networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PodSelector: metav1.LabelSelector{},
					Ingress:     []networking.NetworkPolicyIngressRule{},
				},
			},
			update: networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PodSelector: metav1.LabelSelector{
						MatchLabels: map[string]string{"a": "b"},
					},
					Ingress: []networking.NetworkPolicyIngressRule{},
				},
			},
		},
	}

	for testName, successCase := range successCases {
		successCase.old.ObjectMeta.ResourceVersion = "1"
		successCase.update.ObjectMeta.ResourceVersion = "1"
		if errs := ValidateNetworkPolicyUpdate(&successCase.update, &successCase.old); len(errs) != 0 {
			t.Errorf("expected success (%s): %v", testName, errs)
		}
	}

	errorCases := map[string]npUpdateTest{
		"change name": {
			old: networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PodSelector: metav1.LabelSelector{},
					Ingress:     []networking.NetworkPolicyIngressRule{},
				},
			},
			update: networking.NetworkPolicy{
				ObjectMeta: metav1.ObjectMeta{Name: "baz", Namespace: "bar"},
				Spec: networking.NetworkPolicySpec{
					PodSelector: metav1.LabelSelector{},
					Ingress:     []networking.NetworkPolicyIngressRule{},
				},
			},
		},
	}

	for testName, errorCase := range errorCases {
		errorCase.old.ObjectMeta.ResourceVersion = "1"
		errorCase.update.ObjectMeta.ResourceVersion = "1"
		if errs := ValidateNetworkPolicyUpdate(&errorCase.update, &errorCase.old); len(errs) == 0 {
			t.Errorf("expected failure: %s", testName)
		}
	}
}

func TestValidateNetworkPolicyStatusUpdate(t *testing.T) {

	type netpolStatusCases struct {
		obj          networking.NetworkPolicyStatus
		expectedErrs field.ErrorList
	}

	testCases := map[string]netpolStatusCases{
		"valid conditions": {
			obj: networking.NetworkPolicyStatus{
				Conditions: []metav1.Condition{
					{
						Type:   string(networking.NetworkPolicyConditionStatusAccepted),
						Status: metav1.ConditionTrue,
						LastTransitionTime: metav1.Time{
							Time: time.Now().Add(-5 * time.Minute),
						},
						Reason:             "RuleApplied",
						Message:            "rule was successfully applied",
						ObservedGeneration: 2,
					},
					{
						Type:   string(networking.NetworkPolicyConditionStatusFailure),
						Status: metav1.ConditionFalse,
						LastTransitionTime: metav1.Time{
							Time: time.Now().Add(-5 * time.Minute),
						},
						Reason:             "RuleApplied",
						Message:            "no error was found",
						ObservedGeneration: 2,
					},
				},
			},
			expectedErrs: field.ErrorList{},
		},
		"duplicate type": {
			obj: networking.NetworkPolicyStatus{
				Conditions: []metav1.Condition{
					{
						Type:   string(networking.NetworkPolicyConditionStatusAccepted),
						Status: metav1.ConditionTrue,
						LastTransitionTime: metav1.Time{
							Time: time.Now().Add(-5 * time.Minute),
						},
						Reason:             "RuleApplied",
						Message:            "rule was successfully applied",
						ObservedGeneration: 2,
					},
					{
						Type:   string(networking.NetworkPolicyConditionStatusAccepted),
						Status: metav1.ConditionFalse,
						LastTransitionTime: metav1.Time{
							Time: time.Now().Add(-5 * time.Minute),
						},
						Reason:             string(networking.NetworkPolicyConditionReasonFeatureNotSupported),
						Message:            "endport is not supported",
						ObservedGeneration: 2,
					},
				},
			},
			expectedErrs: field.ErrorList{field.Duplicate(field.NewPath("status").Child("conditions").Index(1).Child("type"),
				string(networking.NetworkPolicyConditionStatusAccepted))},
		},
		"invalid generation": {
			obj: networking.NetworkPolicyStatus{
				Conditions: []metav1.Condition{
					{
						Type:   string(networking.NetworkPolicyConditionStatusAccepted),
						Status: metav1.ConditionTrue,
						LastTransitionTime: metav1.Time{
							Time: time.Now().Add(-5 * time.Minute),
						},
						Reason:             "RuleApplied",
						Message:            "rule was successfully applied",
						ObservedGeneration: -1,
					},
				},
			},
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("status").Child("conditions").Index(0).Child("observedGeneration"),
				int64(-1), "must be greater than or equal to zero")},
		},
		"invalid null transition time": {
			obj: networking.NetworkPolicyStatus{
				Conditions: []metav1.Condition{
					{
						Type:               string(networking.NetworkPolicyConditionStatusAccepted),
						Status:             metav1.ConditionTrue,
						Reason:             "RuleApplied",
						Message:            "rule was successfully applied",
						ObservedGeneration: 3,
					},
				},
			},
			expectedErrs: field.ErrorList{field.Required(field.NewPath("status").Child("conditions").Index(0).Child("lastTransitionTime"),
				"must be set")},
		},
		"multiple condition errors": {
			obj: networking.NetworkPolicyStatus{
				Conditions: []metav1.Condition{
					{
						Type:               string(networking.NetworkPolicyConditionStatusAccepted),
						Status:             metav1.ConditionTrue,
						Reason:             "RuleApplied",
						Message:            "rule was successfully applied",
						ObservedGeneration: -1,
					},
				},
			},
			expectedErrs: field.ErrorList{
				field.Invalid(field.NewPath("status").Child("conditions").Index(0).Child("observedGeneration"),
					int64(-1), "must be greater than or equal to zero"),
				field.Required(field.NewPath("status").Child("conditions").Index(0).Child("lastTransitionTime"),
					"must be set"),
			},
		},
	}

	for testName, testCase := range testCases {
		errs := ValidateNetworkPolicyStatusUpdate(testCase.obj, networking.NetworkPolicyStatus{}, field.NewPath("status"))
		if len(errs) != len(testCase.expectedErrs) {
			t.Errorf("Test %s: Expected %d errors, got %d (%+v)", testName, len(testCase.expectedErrs), len(errs), errs)
		}

		for i, err := range errs {
			if err.Error() != testCase.expectedErrs[i].Error() {
				t.Errorf("Test %s: Expected error: %v, got %v", testName, testCase.expectedErrs[i], err)
			}
		}
	}

}

func TestValidateIngress(t *testing.T) {
	serviceBackend := &networking.IngressServiceBackend{
		Name: "defaultbackend",
		Port: networking.ServiceBackendPort{
			Name:   "",
			Number: 80,
		},
	}
	defaultBackend := networking.IngressBackend{
		Service: serviceBackend,
	}
	pathTypePrefix := networking.PathTypePrefix
	pathTypeImplementationSpecific := networking.PathTypeImplementationSpecific
	pathTypeFoo := networking.PathType("foo")

	baseIngress := networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: metav1.NamespaceDefault,
		},
		Spec: networking.IngressSpec{
			DefaultBackend: &defaultBackend,
			Rules: []networking.IngressRule{
				{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path:     "/foo",
									PathType: &pathTypeImplementationSpecific,
									Backend:  defaultBackend,
								},
							},
						},
					},
				},
			},
		},
		Status: networking.IngressStatus{
			LoadBalancer: api.LoadBalancerStatus{
				Ingress: []api.LoadBalancerIngress{
					{IP: "127.0.0.1"},
				},
			},
		},
	}

	testCases := map[string]struct {
		tweakIngress       func(ing *networking.Ingress)
		expectErrsOnFields []string
	}{
		"empty path (implementation specific)": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Path = ""
			},
			expectErrsOnFields: []string{},
		},
		"valid path": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Path = "/valid"
			},
			expectErrsOnFields: []string{},
		},
		// invalid use cases
		"backend with no service": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.DefaultBackend.Service.Name = ""
			},
			expectErrsOnFields: []string{
				"spec.defaultBackend.service.name",
			},
		},
		"invalid path type": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].PathType = &pathTypeFoo
			},
			expectErrsOnFields: []string{
				"spec.rules[0].http.paths[0].pathType",
			},
		},
		"empty path (prefix)": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Path = ""
				ing.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].PathType = &pathTypePrefix
			},
			expectErrsOnFields: []string{
				"spec.rules[0].http.paths[0].path",
			},
		},
		"no paths": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].IngressRuleValue.HTTP.Paths = []networking.HTTPIngressPath{}
			},
			expectErrsOnFields: []string{
				"spec.rules[0].http.paths",
			},
		},
		"invalid host (foobar:80)": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].Host = "foobar:80"
			},
			expectErrsOnFields: []string{
				"spec.rules[0].host",
			},
		},
		"invalid host (127.0.0.1)": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].Host = "127.0.0.1"
			},
			expectErrsOnFields: []string{
				"spec.rules[0].host",
			},
		},
		"valid wildcard host": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].Host = "*.bar.com"
			},
			expectErrsOnFields: []string{},
		},
		"invalid wildcard host (foo.*.bar.com)": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].Host = "foo.*.bar.com"
			},
			expectErrsOnFields: []string{
				"spec.rules[0].host",
			},
		},
		"invalid wildcard host (*)": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].Host = "*"
			},
			expectErrsOnFields: []string{
				"spec.rules[0].host",
			},
		},
		"path resource backend and service name are not allowed together": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].IngressRuleValue = networking.IngressRuleValue{
					HTTP: &networking.HTTPIngressRuleValue{
						Paths: []networking.HTTPIngressPath{
							{
								Path:     "/foo",
								PathType: &pathTypeImplementationSpecific,
								Backend: networking.IngressBackend{
									Service: serviceBackend,
									Resource: &api.TypedLocalObjectReference{
										APIGroup: utilpointer.StringPtr("example.com"),
										Kind:     "foo",
										Name:     "bar",
									},
								},
							},
						},
					},
				}
			},
			expectErrsOnFields: []string{
				"spec.rules[0].http.paths[0].backend",
			},
		},
		"path resource backend and service port are not allowed together": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.Rules[0].IngressRuleValue = networking.IngressRuleValue{
					HTTP: &networking.HTTPIngressRuleValue{
						Paths: []networking.HTTPIngressPath{
							{
								Path:     "/foo",
								PathType: &pathTypeImplementationSpecific,
								Backend: networking.IngressBackend{
									Service: serviceBackend,
									Resource: &api.TypedLocalObjectReference{
										APIGroup: utilpointer.StringPtr("example.com"),
										Kind:     "foo",
										Name:     "bar",
									},
								},
							},
						},
					},
				}
			},
			expectErrsOnFields: []string{
				"spec.rules[0].http.paths[0].backend",
			},
		},
		"spec.backend resource and service name are not allowed together": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.DefaultBackend = &networking.IngressBackend{
					Service: serviceBackend,
					Resource: &api.TypedLocalObjectReference{
						APIGroup: utilpointer.StringPtr("example.com"),
						Kind:     "foo",
						Name:     "bar",
					},
				}
			},
			expectErrsOnFields: []string{
				"spec.defaultBackend",
			},
		},
		"spec.backend resource and service port are not allowed together": {
			tweakIngress: func(ing *networking.Ingress) {
				ing.Spec.DefaultBackend = &networking.IngressBackend{
					Service: serviceBackend,
					Resource: &api.TypedLocalObjectReference{
						APIGroup: utilpointer.StringPtr("example.com"),
						Kind:     "foo",
						Name:     "bar",
					},
				}
			},
			expectErrsOnFields: []string{
				"spec.defaultBackend",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ingress := baseIngress.DeepCopy()
			testCase.tweakIngress(ingress)
			errs := validateIngress(ingress, IngressValidationOptions{})
			if len(testCase.expectErrsOnFields) != len(errs) {
				t.Fatalf("Expected %d errors, got %d errors: %v", len(testCase.expectErrsOnFields), len(errs), errs)
			}
			for i, err := range errs {
				if err.Field != testCase.expectErrsOnFields[i] {
					t.Errorf("Expected error on field: %s, got: %s", testCase.expectErrsOnFields[i], err.Error())
				}
			}
		})
	}
}

func TestValidateIngressRuleValue(t *testing.T) {
	serviceBackend := networking.IngressServiceBackend{
		Name: "defaultbackend",
		Port: networking.ServiceBackendPort{
			Name:   "",
			Number: 80,
		},
	}
	fldPath := field.NewPath("testing.http.paths[0].path")
	testCases := map[string]struct {
		pathType     networking.PathType
		path         string
		expectedErrs field.ErrorList
	}{
		"implementation specific: no leading slash": {
			pathType:     networking.PathTypeImplementationSpecific,
			path:         "foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "foo", "must be an absolute path")},
		},
		"implementation specific: leading slash": {
			pathType:     networking.PathTypeImplementationSpecific,
			path:         "/foo",
			expectedErrs: field.ErrorList{},
		},
		"implementation specific: many slashes": {
			pathType:     networking.PathTypeImplementationSpecific,
			path:         "/foo/bar/foo",
			expectedErrs: field.ErrorList{},
		},
		"implementation specific: repeating slashes": {
			pathType:     networking.PathTypeImplementationSpecific,
			path:         "/foo//bar/foo",
			expectedErrs: field.ErrorList{},
		},
		"prefix: no leading slash": {
			pathType:     networking.PathTypePrefix,
			path:         "foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "foo", "must be an absolute path")},
		},
		"prefix: leading slash": {
			pathType:     networking.PathTypePrefix,
			path:         "/foo",
			expectedErrs: field.ErrorList{},
		},
		"prefix: many slashes": {
			pathType:     networking.PathTypePrefix,
			path:         "/foo/bar/foo",
			expectedErrs: field.ErrorList{},
		},
		"prefix: repeating slashes": {
			pathType:     networking.PathTypePrefix,
			path:         "/foo//bar/foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "/foo//bar/foo", "must not contain '//'")},
		},
		"exact: no leading slash": {
			pathType:     networking.PathTypeExact,
			path:         "foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "foo", "must be an absolute path")},
		},
		"exact: leading slash": {
			pathType:     networking.PathTypeExact,
			path:         "/foo",
			expectedErrs: field.ErrorList{},
		},
		"exact: many slashes": {
			pathType:     networking.PathTypeExact,
			path:         "/foo/bar/foo",
			expectedErrs: field.ErrorList{},
		},
		"exact: repeating slashes": {
			pathType:     networking.PathTypeExact,
			path:         "/foo//bar/foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "/foo//bar/foo", "must not contain '//'")},
		},
		"prefix: with /./": {
			pathType:     networking.PathTypePrefix,
			path:         "/foo/./foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "/foo/./foo", "must not contain '/./'")},
		},
		"exact: with /../": {
			pathType:     networking.PathTypeExact,
			path:         "/foo/../foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "/foo/../foo", "must not contain '/../'")},
		},
		"prefix: with %2f": {
			pathType:     networking.PathTypePrefix,
			path:         "/foo/%2f/foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "/foo/%2f/foo", "must not contain '%2f'")},
		},
		"exact: with %2F": {
			pathType:     networking.PathTypeExact,
			path:         "/foo/%2F/foo",
			expectedErrs: field.ErrorList{field.Invalid(fldPath, "/foo/%2F/foo", "must not contain '%2F'")},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			irv := &networking.IngressRuleValue{
				HTTP: &networking.HTTPIngressRuleValue{
					Paths: []networking.HTTPIngressPath{
						{
							Path:     testCase.path,
							PathType: &testCase.pathType,
							Backend: networking.IngressBackend{
								Service: &serviceBackend,
							},
						},
					},
				},
			}
			errs := validateIngressRuleValue(irv, field.NewPath("testing"), IngressValidationOptions{})
			if len(errs) != len(testCase.expectedErrs) {
				t.Fatalf("Expected %d errors, got %d (%+v)", len(testCase.expectedErrs), len(errs), errs)
			}

			for i, err := range errs {
				if err.Error() != testCase.expectedErrs[i].Error() {
					t.Fatalf("Expected error: %v, got %v", testCase.expectedErrs[i], err)
				}
			}
		})
	}
}

func TestValidateIngressCreate(t *testing.T) {
	implementationPathType := networking.PathTypeImplementationSpecific
	exactPathType := networking.PathTypeExact
	serviceBackend := &networking.IngressServiceBackend{
		Name: "defaultbackend",
		Port: networking.ServiceBackendPort{
			Number: 80,
		},
	}
	defaultBackend := networking.IngressBackend{
		Service: serviceBackend,
	}
	resourceBackend := &api.TypedLocalObjectReference{
		APIGroup: utilpointer.StringPtr("example.com"),
		Kind:     "foo",
		Name:     "bar",
	}
	baseIngress := networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:            "test123",
			Namespace:       "test123",
			ResourceVersion: "1234",
		},
		Spec: networking.IngressSpec{
			DefaultBackend: &defaultBackend,
			Rules:          []networking.IngressRule{},
		},
	}

	testCases := map[string]struct {
		tweakIngress func(ingress *networking.Ingress)
		expectedErrs field.ErrorList
	}{
		"class field set": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.IngressClassName = utilpointer.StringPtr("bar")
			},
			expectedErrs: field.ErrorList{},
		},
		"class annotation set": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Annotations = map[string]string{annotationIngressClass: "foo"}
			},
			expectedErrs: field.ErrorList{},
		},
		"class field and annotation set": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.IngressClassName = utilpointer.StringPtr("bar")
				ingress.Annotations = map[string]string{annotationIngressClass: "foo"}
			},
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("annotations").Child(annotationIngressClass), "foo", "can not be set when the class field is also set")},
		},
		"valid regex path": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/([a-z0-9]*)",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"invalid regex path allowed (v1)": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/([a-z0-9]*)[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"Spec.Backend.Resource field allowed on create": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.DefaultBackend = &networking.IngressBackend{
					Resource: resourceBackend}
			},
			expectedErrs: field.ErrorList{},
		},
		"Paths.Backend.Resource field allowed on create": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/([a-z0-9]*)",
								PathType: &implementationPathType,
								Backend: networking.IngressBackend{
									Resource: resourceBackend},
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"valid secret": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.TLS = []networking.IngressTLS{{SecretName: "valid"}}
			},
		},
		"invalid secret": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.TLS = []networking.IngressTLS{{SecretName: "invalid name"}}
			},
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("spec").Child("tls").Index(0).Child("secretName"), "invalid name", `a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')`)},
		},
		"valid rules with wildcard host": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.TLS = []networking.IngressTLS{{Hosts: []string{"*.bar.com"}}}
				ingress.Spec.Rules = []networking.IngressRule{{
					Host: "*.foo.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/foo",
								PathType: &exactPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
		},
		"invalid rules with wildcard host": {
			tweakIngress: func(ingress *networking.Ingress) {
				ingress.Spec.TLS = []networking.IngressTLS{{Hosts: []string{"*.bar.com"}}}
				ingress.Spec.Rules = []networking.IngressRule{{
					Host: "*.foo.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "foo",
								PathType: &exactPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("spec").Child("rules").Index(0).Child("http").Child("paths").Index(0).Child("path"), "foo", `must be an absolute path`)},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			newIngress := baseIngress.DeepCopy()
			testCase.tweakIngress(newIngress)
			errs := ValidateIngressCreate(newIngress)
			if len(errs) != len(testCase.expectedErrs) {
				t.Fatalf("Expected %d errors, got %d (%+v)", len(testCase.expectedErrs), len(errs), errs)
			}

			for i, err := range errs {
				if err.Error() != testCase.expectedErrs[i].Error() {
					t.Fatalf("Expected error: %v, got %v", testCase.expectedErrs[i].Error(), err.Error())
				}
			}
		})
	}
}

func TestValidateIngressUpdate(t *testing.T) {
	implementationPathType := networking.PathTypeImplementationSpecific
	exactPathType := networking.PathTypeExact
	serviceBackend := &networking.IngressServiceBackend{
		Name: "defaultbackend",
		Port: networking.ServiceBackendPort{
			Number: 80,
		},
	}
	defaultBackend := networking.IngressBackend{
		Service: serviceBackend,
	}
	resourceBackend := &api.TypedLocalObjectReference{
		APIGroup: utilpointer.StringPtr("example.com"),
		Kind:     "foo",
		Name:     "bar",
	}
	baseIngress := networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:            "test123",
			Namespace:       "test123",
			ResourceVersion: "1234",
		},
		Spec: networking.IngressSpec{
			DefaultBackend: &defaultBackend,
		},
	}

	testCases := map[string]struct {
		tweakIngresses func(newIngress, oldIngress *networking.Ingress)
		expectedErrs   field.ErrorList
	}{
		"class field set": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				newIngress.Spec.IngressClassName = utilpointer.StringPtr("bar")
			},
			expectedErrs: field.ErrorList{},
		},
		"class annotation set": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				newIngress.Annotations = map[string]string{annotationIngressClass: "foo"}
			},
			expectedErrs: field.ErrorList{},
		},
		"class field and annotation set": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				newIngress.Spec.IngressClassName = utilpointer.StringPtr("bar")
				newIngress.Annotations = map[string]string{annotationIngressClass: "foo"}
			},
			expectedErrs: field.ErrorList{},
		},
		"valid regex path -> valid regex path": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/([a-z0-9]*)",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/([a-z0-9%]*)",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"valid regex path -> invalid regex path": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/([a-z0-9]*)",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/bar[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"invalid regex path -> valid regex path": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/bar[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/([a-z0-9]*)",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"invalid regex path -> invalid regex path": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/foo[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/bar[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"new Backend.Resource allowed on update": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.DefaultBackend = &defaultBackend
				newIngress.Spec.DefaultBackend = &networking.IngressBackend{
					Resource: resourceBackend}
			},
			expectedErrs: field.ErrorList{},
		},
		"old DefaultBackend.Resource allowed on update": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.DefaultBackend = &networking.IngressBackend{
					Resource: resourceBackend}
				newIngress.Spec.DefaultBackend = &networking.IngressBackend{
					Resource: resourceBackend}
			},
			expectedErrs: field.ErrorList{},
		},
		"changing spec.backend from resource -> no resource": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.DefaultBackend = &networking.IngressBackend{
					Resource: resourceBackend}
				newIngress.Spec.DefaultBackend = &defaultBackend
			},
			expectedErrs: field.ErrorList{},
		},
		"changing path backend from resource -> no resource": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/foo[",
								PathType: &implementationPathType,
								Backend: networking.IngressBackend{
									Resource: resourceBackend},
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/bar[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"changing path backend from resource -> resource": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/foo[",
								PathType: &implementationPathType,
								Backend: networking.IngressBackend{
									Resource: resourceBackend},
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/bar[",
								PathType: &implementationPathType,
								Backend: networking.IngressBackend{
									Resource: resourceBackend},
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"changing path backend from non-resource -> non-resource": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/foo[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/bar[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"changing path backend from non-resource -> resource": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/foo[",
								PathType: &implementationPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "foo.bar.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/bar[",
								PathType: &implementationPathType,
								Backend: networking.IngressBackend{
									Resource: resourceBackend},
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{},
		},
		"change valid secret -> invalid secret": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.TLS = []networking.IngressTLS{{SecretName: "valid"}}
				newIngress.Spec.TLS = []networking.IngressTLS{{SecretName: "invalid name"}}
			},
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("spec").Child("tls").Index(0).Child("secretName"), "invalid name", `a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')`)},
		},
		"change invalid secret -> invalid secret": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.TLS = []networking.IngressTLS{{SecretName: "invalid name 1"}}
				newIngress.Spec.TLS = []networking.IngressTLS{{SecretName: "invalid name 2"}}
			},
		},
		"change valid rules with wildcard host -> invalid rules": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.TLS = []networking.IngressTLS{{Hosts: []string{"*.bar.com"}}}
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "*.foo.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "/foo",
								PathType: &exactPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.TLS = []networking.IngressTLS{{Hosts: []string{"*.bar.com"}}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "*.foo.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "foo",
								PathType: &exactPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("spec").Child("rules").Index(0).Child("http").Child("paths").Index(0).Child("path"), "foo", `must be an absolute path`)},
		},
		"change invalid rules with wildcard host -> invalid rules": {
			tweakIngresses: func(newIngress, oldIngress *networking.Ingress) {
				oldIngress.Spec.TLS = []networking.IngressTLS{{Hosts: []string{"*.bar.com"}}}
				oldIngress.Spec.Rules = []networking.IngressRule{{
					Host: "*.foo.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "foo",
								PathType: &exactPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
				newIngress.Spec.TLS = []networking.IngressTLS{{Hosts: []string{"*.bar.com"}}}
				newIngress.Spec.Rules = []networking.IngressRule{{
					Host: "*.foo.com",
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{{
								Path:     "bar",
								PathType: &exactPathType,
								Backend:  defaultBackend,
							}},
						},
					},
				}}
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			newIngress := baseIngress.DeepCopy()
			oldIngress := baseIngress.DeepCopy()
			testCase.tweakIngresses(newIngress, oldIngress)

			errs := ValidateIngressUpdate(newIngress, oldIngress)

			if len(errs) != len(testCase.expectedErrs) {
				t.Fatalf("Expected %d errors, got %d (%+v)", len(testCase.expectedErrs), len(errs), errs)
			}

			for i, err := range errs {
				if err.Error() != testCase.expectedErrs[i].Error() {
					t.Fatalf("Expected error: %v, got %v", testCase.expectedErrs[i].Error(), err.Error())
				}
			}
		})
	}
}

type netIngressTweak func(ingressClass *networking.IngressClass)

func makeValidIngressClass(name, controller string, tweaks ...netIngressTweak) *networking.IngressClass {
	ingressClass := &networking.IngressClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: networking.IngressClassSpec{
			Controller: controller,
		},
	}

	for _, fn := range tweaks {
		fn(ingressClass)
	}
	return ingressClass
}

func makeIngressClassParams(apiGroup *string, kind, name string, scope, namespace *string) *networking.IngressClassParametersReference {
	return &networking.IngressClassParametersReference{
		APIGroup:  apiGroup,
		Kind:      kind,
		Name:      name,
		Scope:     scope,
		Namespace: namespace,
	}
}

func TestValidateIngressClass(t *testing.T) {
	setParams := func(params *networking.IngressClassParametersReference) netIngressTweak {
		return func(ingressClass *networking.IngressClass) {
			ingressClass.Spec.Parameters = params
		}
	}

	testCases := map[string]struct {
		ingressClass *networking.IngressClass
		expectedErrs field.ErrorList
	}{
		"valid name, valid controller": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar"),
			expectedErrs: field.ErrorList{},
		},
		"invalid name, valid controller": {
			ingressClass: makeValidIngressClass("test*123", "foo.co/bar"),
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("metadata.name"), "test*123", "a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')")},
		},
		"valid name, empty controller": {
			ingressClass: makeValidIngressClass("test123", ""),
			expectedErrs: field.ErrorList{field.Required(field.NewPath("spec.controller"), "")},
		},
		"valid name, controller max length": {
			ingressClass: makeValidIngressClass("test123", "foo.co/"+strings.Repeat("a", 243)),
			expectedErrs: field.ErrorList{},
		},
		"valid name, controller too long": {
			ingressClass: makeValidIngressClass("test123", "foo.co/"+strings.Repeat("a", 244)),
			expectedErrs: field.ErrorList{field.TooLong(field.NewPath("spec.controller"), "", 250)},
		},
		"valid name, valid controller, valid params": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(utilpointer.StringPtr("example.com"), "foo", "bar", utilpointer.StringPtr("Cluster"), nil)),
			),
			expectedErrs: field.ErrorList{},
		},
		"valid name, valid controller, invalid params (no kind)": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(utilpointer.StringPtr("example.com"), "", "bar", utilpointer.StringPtr("Cluster"), nil)),
			),
			expectedErrs: field.ErrorList{field.Required(field.NewPath("spec.parameters.kind"), "kind is required")},
		},
		"valid name, valid controller, invalid params (no name)": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(utilpointer.StringPtr("example.com"), "foo", "", utilpointer.StringPtr("Cluster"), nil)),
			),
			expectedErrs: field.ErrorList{field.Required(field.NewPath("spec.parameters.name"), "name is required")},
		},
		"valid name, valid controller, invalid params (bad kind)": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo/", "bar", utilpointer.StringPtr("Cluster"), nil)),
			),
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("spec.parameters.kind"), "foo/", "may not contain '/'")},
		},
		"valid name, valid controller, invalid params (bad scope)": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo", "bar", utilpointer.StringPtr("bad-scope"), nil)),
			),
			expectedErrs: field.ErrorList{field.NotSupported(field.NewPath("spec.parameters.scope"),
				"bad-scope", []string{"Cluster", "Namespace"})},
		},
		"valid name, valid controller, valid Namespace scope": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo", "bar", utilpointer.StringPtr("Namespace"), utilpointer.StringPtr("foo-ns"))),
			),
			expectedErrs: field.ErrorList{},
		},
		"valid name, valid controller, valid scope, invalid namespace": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo", "bar", utilpointer.StringPtr("Namespace"), utilpointer.StringPtr("foo_ns"))),
			),
			expectedErrs: field.ErrorList{field.Invalid(field.NewPath("spec.parameters.namespace"), "foo_ns",
				"a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-',"+
					" and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', "+
					"regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')")},
		},
		"valid name, valid controller, valid Cluster scope": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo", "bar", utilpointer.StringPtr("Cluster"), nil)),
			),
			expectedErrs: field.ErrorList{},
		},
		"namespace not set when scope is Namespace": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo", "bar", utilpointer.StringPtr("Namespace"), nil)),
			),
			expectedErrs: field.ErrorList{field.Required(field.NewPath("spec.parameters.namespace"),
				"`parameters.scope` is set to 'Namespace'")},
		},
		"namespace is forbidden when scope is Cluster": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo", "bar", utilpointer.StringPtr("Cluster"), utilpointer.StringPtr("foo-ns"))),
			),
			expectedErrs: field.ErrorList{field.Forbidden(field.NewPath("spec.parameters.namespace"),
				"`parameters.scope` is set to 'Cluster'")},
		},
		"empty namespace is forbidden when scope is Cluster": {
			ingressClass: makeValidIngressClass("test123", "foo.co/bar",
				setParams(makeIngressClassParams(nil, "foo", "bar", utilpointer.StringPtr("Cluster"), utilpointer.StringPtr(""))),
			),
			expectedErrs: field.ErrorList{field.Forbidden(field.NewPath("spec.parameters.namespace"),
				"`parameters.scope` is set to 'Cluster'")},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			errs := ValidateIngressClass(testCase.ingressClass)

			if len(errs) != len(testCase.expectedErrs) {
				t.Fatalf("Expected %d errors, got %d (%+v)", len(testCase.expectedErrs), len(errs), errs)
			}

			for i, err := range errs {
				if err.Error() != testCase.expectedErrs[i].Error() {
					t.Fatalf("Expected error: %v, got %v", testCase.expectedErrs[i].Error(), err.Error())
				}
			}
		})
	}
}

func TestValidateIngressClassUpdate(t *testing.T) {
	setResourceVersion := func(version string) netIngressTweak {
		return func(ingressClass *networking.IngressClass) {
			ingressClass.ObjectMeta.ResourceVersion = version
		}
	}

	setParams := func(params *networking.IngressClassParametersReference) netIngressTweak {
		return func(ingressClass *networking.IngressClass) {
			ingressClass.Spec.Parameters = params
		}
	}

	testCases := map[string]struct {
		newIngressClass *networking.IngressClass
		oldIngressClass *networking.IngressClass
		expectedErrs    field.ErrorList
	}{
		"name change": {
			newIngressClass: makeValidIngressClass("test123", "foo.co/bar", setResourceVersion("2")),
			oldIngressClass: makeValidIngressClass("test123", "foo.co/different"),
			expectedErrs:    field.ErrorList{field.Invalid(field.NewPath("spec").Child("controller"), "foo.co/bar", apimachineryvalidation.FieldImmutableErrorMsg)},
		},
		"parameters change": {
			newIngressClass: makeValidIngressClass("test123", "foo.co/bar",
				setResourceVersion("2"),
				setParams(
					makeIngressClassParams(utilpointer.StringPtr("v1"), "ConfigMap", "foo", utilpointer.StringPtr("Namespace"), utilpointer.StringPtr("bar")),
				),
			),
			oldIngressClass: makeValidIngressClass("test123", "foo.co/bar"),
			expectedErrs:    field.ErrorList{},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			errs := ValidateIngressClassUpdate(testCase.newIngressClass, testCase.oldIngressClass)

			if len(errs) != len(testCase.expectedErrs) {
				t.Fatalf("Expected %d errors, got %d (%+v)", len(testCase.expectedErrs), len(errs), errs)
			}

			for i, err := range errs {
				if err.Error() != testCase.expectedErrs[i].Error() {
					t.Fatalf("Expected error: %v, got %v", testCase.expectedErrs[i].Error(), err.Error())
				}
			}
		})
	}
}

func TestValidateIngressTLS(t *testing.T) {
	pathTypeImplementationSpecific := networking.PathTypeImplementationSpecific
	serviceBackend := &networking.IngressServiceBackend{
		Name: "defaultbackend",
		Port: networking.ServiceBackendPort{
			Number: 80,
		},
	}
	defaultBackend := networking.IngressBackend{
		Service: serviceBackend,
	}
	newValid := func() networking.Ingress {
		return networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: metav1.NamespaceDefault,
			},
			Spec: networking.IngressSpec{
				DefaultBackend: &defaultBackend,
				Rules: []networking.IngressRule{
					{
						Host: "foo.bar.com",
						IngressRuleValue: networking.IngressRuleValue{
							HTTP: &networking.HTTPIngressRuleValue{
								Paths: []networking.HTTPIngressPath{
									{
										Path:     "/foo",
										PathType: &pathTypeImplementationSpecific,
										Backend:  defaultBackend,
									},
								},
							},
						},
					},
				},
			},
			Status: networking.IngressStatus{
				LoadBalancer: api.LoadBalancerStatus{
					Ingress: []api.LoadBalancerIngress{
						{IP: "127.0.0.1"},
					},
				},
			},
		}
	}

	errorCases := map[string]networking.Ingress{}

	wildcardHost := "foo.*.bar.com"
	badWildcardTLS := newValid()
	badWildcardTLS.Spec.Rules[0].Host = "*.foo.bar.com"
	badWildcardTLS.Spec.TLS = []networking.IngressTLS{
		{
			Hosts: []string{wildcardHost},
		},
	}
	badWildcardTLSErr := fmt.Sprintf("spec.tls[0].hosts[0]: Invalid value: '%v'", wildcardHost)
	errorCases[badWildcardTLSErr] = badWildcardTLS

	for k, v := range errorCases {
		errs := validateIngress(&v, IngressValidationOptions{})
		if len(errs) == 0 {
			t.Errorf("expected failure for %q", k)
		} else {
			s := strings.Split(k, ":")
			err := errs[0]
			if err.Field != s[0] || !strings.Contains(err.Error(), s[1]) {
				t.Errorf("unexpected error: %q, expected: %q", err, k)
			}
		}
	}

	// Test for wildcard host and wildcard TLS
	validCases := map[string]networking.Ingress{}
	wildHost := "*.bar.com"
	goodWildcardTLS := newValid()
	goodWildcardTLS.Spec.Rules[0].Host = "*.bar.com"
	goodWildcardTLS.Spec.TLS = []networking.IngressTLS{
		{
			Hosts: []string{wildHost},
		},
	}
	validCases[fmt.Sprintf("spec.tls[0].hosts: Valid value: '%v'", wildHost)] = goodWildcardTLS
	for k, v := range validCases {
		errs := validateIngress(&v, IngressValidationOptions{})
		if len(errs) != 0 {
			t.Errorf("expected success for %q", k)
		}
	}
}

// TestValidateEmptyIngressTLS verifies that an empty TLS configuration can be
// specified, which ingress controllers may interpret to mean that TLS should be
// used with a default certificate that the ingress controller furnishes.
func TestValidateEmptyIngressTLS(t *testing.T) {
	pathTypeImplementationSpecific := networking.PathTypeImplementationSpecific
	serviceBackend := &networking.IngressServiceBackend{
		Name: "defaultbackend",
		Port: networking.ServiceBackendPort{
			Number: 443,
		},
	}
	defaultBackend := networking.IngressBackend{
		Service: serviceBackend,
	}
	newValid := func() networking.Ingress {
		return networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: metav1.NamespaceDefault,
			},
			Spec: networking.IngressSpec{
				Rules: []networking.IngressRule{
					{
						Host: "foo.bar.com",
						IngressRuleValue: networking.IngressRuleValue{
							HTTP: &networking.HTTPIngressRuleValue{
								Paths: []networking.HTTPIngressPath{
									{
										PathType: &pathTypeImplementationSpecific,
										Backend:  defaultBackend,
									},
								},
							},
						},
					},
				},
			},
		}
	}

	validCases := map[string]networking.Ingress{}
	goodEmptyTLS := newValid()
	goodEmptyTLS.Spec.TLS = []networking.IngressTLS{
		{},
	}
	validCases[fmt.Sprintf("spec.tls[0]: Valid value: %v", goodEmptyTLS.Spec.TLS[0])] = goodEmptyTLS
	goodEmptyHosts := newValid()
	goodEmptyHosts.Spec.TLS = []networking.IngressTLS{
		{
			Hosts: []string{},
		},
	}
	validCases[fmt.Sprintf("spec.tls[0]: Valid value: %v", goodEmptyHosts.Spec.TLS[0])] = goodEmptyHosts
	for k, v := range validCases {
		errs := validateIngress(&v, IngressValidationOptions{})
		if len(errs) != 0 {
			t.Errorf("expected success for %q", k)
		}
	}
}

func TestValidateIngressStatusUpdate(t *testing.T) {
	serviceBackend := &networking.IngressServiceBackend{
		Name: "defaultbackend",
		Port: networking.ServiceBackendPort{
			Number: 80,
		},
	}
	defaultBackend := networking.IngressBackend{
		Service: serviceBackend,
	}

	newValid := func() networking.Ingress {
		return networking.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:            "foo",
				Namespace:       metav1.NamespaceDefault,
				ResourceVersion: "9",
			},
			Spec: networking.IngressSpec{
				DefaultBackend: &defaultBackend,
				Rules: []networking.IngressRule{
					{
						Host: "foo.bar.com",
						IngressRuleValue: networking.IngressRuleValue{
							HTTP: &networking.HTTPIngressRuleValue{
								Paths: []networking.HTTPIngressPath{
									{
										Path:    "/foo",
										Backend: defaultBackend,
									},
								},
							},
						},
					},
				},
			},
			Status: networking.IngressStatus{
				LoadBalancer: api.LoadBalancerStatus{
					Ingress: []api.LoadBalancerIngress{
						{IP: "127.0.0.1", Hostname: "foo.bar.com"},
					},
				},
			},
		}
	}
	oldValue := newValid()
	newValue := newValid()
	newValue.Status = networking.IngressStatus{
		LoadBalancer: api.LoadBalancerStatus{
			Ingress: []api.LoadBalancerIngress{
				{IP: "127.0.0.2", Hostname: "foo.com"},
			},
		},
	}
	invalidIP := newValid()
	invalidIP.Status = networking.IngressStatus{
		LoadBalancer: api.LoadBalancerStatus{
			Ingress: []api.LoadBalancerIngress{
				{IP: "abcd", Hostname: "foo.com"},
			},
		},
	}
	invalidHostname := newValid()
	invalidHostname.Status = networking.IngressStatus{
		LoadBalancer: api.LoadBalancerStatus{
			Ingress: []api.LoadBalancerIngress{
				{IP: "127.0.0.1", Hostname: "127.0.0.1"},
			},
		},
	}

	errs := ValidateIngressStatusUpdate(&newValue, &oldValue)
	if len(errs) != 0 {
		t.Errorf("Unexpected error %v", errs)
	}

	errorCases := map[string]networking.Ingress{
		"status.loadBalancer.ingress[0].ip: Invalid value":       invalidIP,
		"status.loadBalancer.ingress[0].hostname: Invalid value": invalidHostname,
	}
	for k, v := range errorCases {
		errs := ValidateIngressStatusUpdate(&v, &oldValue)
		if len(errs) == 0 {
			t.Errorf("expected failure for %s", k)
		} else {
			s := strings.Split(k, ":")
			err := errs[0]
			if err.Field != s[0] || !strings.Contains(err.Error(), s[1]) {
				t.Errorf("unexpected error: %q, expected: %q", err, k)
			}
		}
	}
}

func makeValidClusterCIDRConfig() *networking.ClusterCIDRConfig {
	return &networking.ClusterCIDRConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:            "foo",
			ResourceVersion: "9",
		},
		Spec: networking.ClusterCIDRConfigSpec{
			PerNodeHostBits: int32(8),
			IPv4CIDR:        "10.1.0.0/16",
			IPv6CIDR:        "fd00:1:1::/64",
			NodeSelector: &api.NodeSelector{
				NodeSelectorTerms: []api.NodeSelectorTerm{
					{
						MatchExpressions: []api.NodeSelectorRequirement{
							{
								Key:      "foo",
								Operator: api.NodeSelectorOpIn,
								Values:   []string{"bar"},
							},
						},
					},
				},
			},
		},
	}
}

type cccTweak func(ccc *networking.ClusterCIDRConfig)

func makeClusterCIDRConfigCustom(tweaks ...cccTweak) *networking.ClusterCIDRConfig {
	ccc := makeValidClusterCIDRConfig()
	for _, fn := range tweaks {
		fn(ccc)
	}
	return ccc
}

func makeNodeSelector(key string, op api.NodeSelectorOperator, values []string) *api.NodeSelector {
	return &api.NodeSelector{
		NodeSelectorTerms: []api.NodeSelectorTerm{
			{
				MatchExpressions: []api.NodeSelectorRequirement{
					{
						Key:      key,
						Operator: op,
						Values:   values,
					},
				},
			},
		},
	}
}

func TestValidateClusterCIDRConfig(t *testing.T) {
	// Tweaks used below.
	setIPv4CIDR := func(perNodeHostBits int32, ipv4CIDR string) cccTweak {
		return func(ccc *networking.ClusterCIDRConfig) {
			ccc.Spec.IPv4CIDR = ipv4CIDR
			ccc.Spec.PerNodeHostBits = perNodeHostBits
		}
	}

	setIPv6CIDR := func(perNodeHostBits int32, ipv6CIDR string) cccTweak {
		return func(ccc *networking.ClusterCIDRConfig) {
			ccc.Spec.IPv6CIDR = ipv6CIDR
			ccc.Spec.PerNodeHostBits = perNodeHostBits
		}
	}

	setNodeSelector := func(nodeSelector *api.NodeSelector) cccTweak {
		return func(ccc *networking.ClusterCIDRConfig) {
			ccc.Spec.NodeSelector = nodeSelector
		}
	}

	validNodeSelector := makeNodeSelector("foo", api.NodeSelectorOpIn, []string{"bar"})

	successCases := map[string]*networking.ClusterCIDRConfig{
		"valid IPv6 only ClusterCIDRConfig":                      makeClusterCIDRConfigCustom(setIPv4CIDR(8, "")),
		"valid IPv4 only ClusterCIDRConfig":                      makeClusterCIDRConfigCustom(setIPv6CIDR(8, "")),
		"valid DualStack ClusterCIDRConfig with no NodeSelector": makeClusterCIDRConfigCustom(setNodeSelector(nil)),
		"valid NodeSelector":                                     makeClusterCIDRConfigCustom(setNodeSelector(validNodeSelector)),
	}

	// Success cases are expected to pass validation.

	for k, v := range successCases {
		if errs := ValidateClusterCIDRConfig(v); len(errs) != 0 {
			t.Errorf("Expected success for test '%s', got %v", k, errs)
		}
	}

	invalidNodeSelector := makeNodeSelector("NoUppercaseOrSpecialCharsLike=Equals", api.NodeSelectorOpIn, []string{"bar"})

	errorCases := map[string]*networking.ClusterCIDRConfig{
		// Config test.
		"empty spec.IPv4CIDR and spec.IPv6CIDR": makeClusterCIDRConfigCustom(
			setIPv4CIDR(8, ""), setIPv6CIDR(8, "")),
		"invalid spec.NodeSelector": makeClusterCIDRConfigCustom(
			setNodeSelector(invalidNodeSelector)),

		// IPv4 tests.
		"invalid spec.IPv4CIDR": makeClusterCIDRConfigCustom(
			setIPv4CIDR(8, "test")),
		"valid IPv6 CIDR in spec.IPv4CIDR": makeClusterCIDRConfigCustom(
			setIPv4CIDR(8, "fd00::/120")),
		"invalid spec.PerNodeHostBits with IPv4 CIDR": makeClusterCIDRConfigCustom(
			setIPv4CIDR(100, "10.2.0.0/16")),
		"invalid spec.IPv4.PerNodeHostBits > CIDR Host Bits": makeClusterCIDRConfigCustom(
			setIPv4CIDR(24, "10.2.0.0/16")),

		// IPv6 tests.
		"invalid spec.IPv6CIDR": makeClusterCIDRConfigCustom(
			setIPv6CIDR(8, "testv6")),
		"valid IPv4 CIDR in spec.IPv6CIDR": makeClusterCIDRConfigCustom(
			setIPv6CIDR(8, "10.2.0.0/16")),
		"invalid spec.PerNodeHostBits with IPv6 CIDR": makeClusterCIDRConfigCustom(
			setIPv6CIDR(1000, "fd00::/120")),
		"invalid spec.IPv6.PerNodeMaskSize < CIDR Mask": makeClusterCIDRConfigCustom(
			setIPv6CIDR(12, "fd00::/120")),
	}

	// Error cases are not expected to pass validation.
	for testName, ccc := range errorCases {
		if errs := ValidateClusterCIDRConfig(ccc); len(errs) == 0 {
			t.Errorf("Expected failure for test: %s", testName)
		}
	}
}

func TestValidateClusterConfigUpdate(t *testing.T) {
	oldCCC := makeValidClusterCIDRConfig()

	// Tweaks used below.
	setIPv4CIDR := func(perNodeHostBits int32, ipv4CIDR string) cccTweak {
		return func(ccc *networking.ClusterCIDRConfig) {
			ccc.Spec.IPv4CIDR = ipv4CIDR
			ccc.Spec.PerNodeHostBits = perNodeHostBits
		}
	}

	setIPv6CIDR := func(perNodeHostBits int32, ipv6CIDR string) cccTweak {
		return func(ccc *networking.ClusterCIDRConfig) {
			ccc.Spec.IPv6CIDR = ipv6CIDR
			ccc.Spec.PerNodeHostBits = perNodeHostBits
		}
	}

	setNodeSelector := func(nodeSelector *api.NodeSelector) cccTweak {
		return func(ccc *networking.ClusterCIDRConfig) {
			ccc.Spec.NodeSelector = nodeSelector
		}
	}

	updateNodeSelector := makeNodeSelector("foo", api.NodeSelectorOpIn, []string{"bar2"})

	successCases := map[string]*networking.ClusterCIDRConfig{
		"update with no tweaks": makeClusterCIDRConfigCustom(),
	}

	// Error cases are not expected to pass validation.
	for testName, ccc := range successCases {
		errs := ValidateClusterCIDRConfigUpdate(ccc, oldCCC)
		if len(errs) != 0 {
			t.Errorf("Expected success for test '%s', got %v", testName, errs)
		}
	}

	errorCases := map[string]*networking.ClusterCIDRConfig{
		"update spec.IPv4": makeClusterCIDRConfigCustom(setIPv4CIDR(8, "10.2.0.0/16")),
		"update spec.IPv6": makeClusterCIDRConfigCustom(setIPv6CIDR(8, "fd00:2:/112")),
		"update spec.NodeSelector": makeClusterCIDRConfigCustom(setNodeSelector(
			updateNodeSelector)),
	}

	// Error cases are not expected to pass validation.
	for testName, ccc := range errorCases {
		errs := ValidateClusterCIDRConfigUpdate(ccc, oldCCC)
		if len(errs) == 0 {
			t.Errorf("Expected failure for test: %s", testName)
		}
	}
}
