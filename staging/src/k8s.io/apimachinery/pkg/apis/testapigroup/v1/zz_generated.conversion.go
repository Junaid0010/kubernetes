// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by: _output/bin/conversion-gen --extra-peer-dirs k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/api/core/v1 --v 1 --logtostderr -i k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission/v1beta1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1alpha1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1beta2,k8s.io/kubernetes/pkg/apis/authentication/v1,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization/v1,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/autoscaling/v2beta1,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v1beta1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates/v1beta1,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/kubernetes/pkg/apis/events/v1beta1,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/networking/v1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac/v1,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac/v1beta1,k8s.io/kubernetes/pkg/apis/scheduling/v1alpha1,k8s.io/kubernetes/pkg/apis/settings/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1,k8s.io/kubernetes/pkg/apis/storage/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1beta1,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/appsv1beta1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/appsv1beta2,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/autoscalingv1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/extensionsv1beta1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1 -O zz_generated.conversion

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testapigroup "k8s.io/apimachinery/pkg/apis/testapigroup"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_Carp_To_testapigroup_Carp,
		Convert_testapigroup_Carp_To_v1_Carp,
		Convert_v1_CarpCondition_To_testapigroup_CarpCondition,
		Convert_testapigroup_CarpCondition_To_v1_CarpCondition,
		Convert_v1_CarpList_To_testapigroup_CarpList,
		Convert_testapigroup_CarpList_To_v1_CarpList,
		Convert_v1_CarpSpec_To_testapigroup_CarpSpec,
		Convert_testapigroup_CarpSpec_To_v1_CarpSpec,
		Convert_v1_CarpStatus_To_testapigroup_CarpStatus,
		Convert_testapigroup_CarpStatus_To_v1_CarpStatus,
	)
}

func autoConvert_v1_Carp_To_testapigroup_Carp(in *Carp, out *testapigroup.Carp, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_CarpSpec_To_testapigroup_CarpSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_CarpStatus_To_testapigroup_CarpStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Carp_To_testapigroup_Carp is an autogenerated conversion function.
func Convert_v1_Carp_To_testapigroup_Carp(in *Carp, out *testapigroup.Carp, s conversion.Scope) error {
	return autoConvert_v1_Carp_To_testapigroup_Carp(in, out, s)
}

func autoConvert_testapigroup_Carp_To_v1_Carp(in *testapigroup.Carp, out *Carp, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_testapigroup_CarpSpec_To_v1_CarpSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_testapigroup_CarpStatus_To_v1_CarpStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_testapigroup_Carp_To_v1_Carp is an autogenerated conversion function.
func Convert_testapigroup_Carp_To_v1_Carp(in *testapigroup.Carp, out *Carp, s conversion.Scope) error {
	return autoConvert_testapigroup_Carp_To_v1_Carp(in, out, s)
}

func autoConvert_v1_CarpCondition_To_testapigroup_CarpCondition(in *CarpCondition, out *testapigroup.CarpCondition, s conversion.Scope) error {
	out.Type = testapigroup.CarpConditionType(in.Type)
	out.Status = testapigroup.ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_CarpCondition_To_testapigroup_CarpCondition is an autogenerated conversion function.
func Convert_v1_CarpCondition_To_testapigroup_CarpCondition(in *CarpCondition, out *testapigroup.CarpCondition, s conversion.Scope) error {
	return autoConvert_v1_CarpCondition_To_testapigroup_CarpCondition(in, out, s)
}

func autoConvert_testapigroup_CarpCondition_To_v1_CarpCondition(in *testapigroup.CarpCondition, out *CarpCondition, s conversion.Scope) error {
	out.Type = CarpConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastProbeTime = in.LastProbeTime
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_testapigroup_CarpCondition_To_v1_CarpCondition is an autogenerated conversion function.
func Convert_testapigroup_CarpCondition_To_v1_CarpCondition(in *testapigroup.CarpCondition, out *CarpCondition, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpCondition_To_v1_CarpCondition(in, out, s)
}

func autoConvert_v1_CarpList_To_testapigroup_CarpList(in *CarpList, out *testapigroup.CarpList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]testapigroup.Carp, len(*in))
		for i := range *in {
			if err := Convert_v1_Carp_To_testapigroup_Carp(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1_CarpList_To_testapigroup_CarpList is an autogenerated conversion function.
func Convert_v1_CarpList_To_testapigroup_CarpList(in *CarpList, out *testapigroup.CarpList, s conversion.Scope) error {
	return autoConvert_v1_CarpList_To_testapigroup_CarpList(in, out, s)
}

func autoConvert_testapigroup_CarpList_To_v1_CarpList(in *testapigroup.CarpList, out *CarpList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Carp, len(*in))
		for i := range *in {
			if err := Convert_testapigroup_Carp_To_v1_Carp(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_testapigroup_CarpList_To_v1_CarpList is an autogenerated conversion function.
func Convert_testapigroup_CarpList_To_v1_CarpList(in *testapigroup.CarpList, out *CarpList, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpList_To_v1_CarpList(in, out, s)
}

func autoConvert_v1_CarpSpec_To_testapigroup_CarpSpec(in *CarpSpec, out *testapigroup.CarpSpec, s conversion.Scope) error {
	out.RestartPolicy = testapigroup.RestartPolicy(in.RestartPolicy)
	out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.ServiceAccountName = in.ServiceAccountName
	// INFO: in.DeprecatedServiceAccount opted out of conversion generation
	out.NodeName = in.NodeName
	// INFO: in.HostNetwork opted out of conversion generation
	// INFO: in.HostPID opted out of conversion generation
	// INFO: in.HostIPC opted out of conversion generation
	out.Hostname = in.Hostname
	out.Subdomain = in.Subdomain
	out.SchedulerName = in.SchedulerName
	return nil
}

// Convert_v1_CarpSpec_To_testapigroup_CarpSpec is an autogenerated conversion function.
func Convert_v1_CarpSpec_To_testapigroup_CarpSpec(in *CarpSpec, out *testapigroup.CarpSpec, s conversion.Scope) error {
	return autoConvert_v1_CarpSpec_To_testapigroup_CarpSpec(in, out, s)
}

func autoConvert_testapigroup_CarpSpec_To_v1_CarpSpec(in *testapigroup.CarpSpec, out *CarpSpec, s conversion.Scope) error {
	out.RestartPolicy = RestartPolicy(in.RestartPolicy)
	out.TerminationGracePeriodSeconds = (*int64)(unsafe.Pointer(in.TerminationGracePeriodSeconds))
	out.ActiveDeadlineSeconds = (*int64)(unsafe.Pointer(in.ActiveDeadlineSeconds))
	out.NodeSelector = *(*map[string]string)(unsafe.Pointer(&in.NodeSelector))
	out.ServiceAccountName = in.ServiceAccountName
	out.NodeName = in.NodeName
	out.Hostname = in.Hostname
	out.Subdomain = in.Subdomain
	out.SchedulerName = in.SchedulerName
	return nil
}

// Convert_testapigroup_CarpSpec_To_v1_CarpSpec is an autogenerated conversion function.
func Convert_testapigroup_CarpSpec_To_v1_CarpSpec(in *testapigroup.CarpSpec, out *CarpSpec, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpSpec_To_v1_CarpSpec(in, out, s)
}

func autoConvert_v1_CarpStatus_To_testapigroup_CarpStatus(in *CarpStatus, out *testapigroup.CarpStatus, s conversion.Scope) error {
	out.Phase = testapigroup.CarpPhase(in.Phase)
	out.Conditions = *(*[]testapigroup.CarpCondition)(unsafe.Pointer(&in.Conditions))
	out.Message = in.Message
	out.Reason = in.Reason
	out.HostIP = in.HostIP
	out.CarpIP = in.CarpIP
	out.StartTime = (*meta_v1.Time)(unsafe.Pointer(in.StartTime))
	return nil
}

// Convert_v1_CarpStatus_To_testapigroup_CarpStatus is an autogenerated conversion function.
func Convert_v1_CarpStatus_To_testapigroup_CarpStatus(in *CarpStatus, out *testapigroup.CarpStatus, s conversion.Scope) error {
	return autoConvert_v1_CarpStatus_To_testapigroup_CarpStatus(in, out, s)
}

func autoConvert_testapigroup_CarpStatus_To_v1_CarpStatus(in *testapigroup.CarpStatus, out *CarpStatus, s conversion.Scope) error {
	out.Phase = CarpPhase(in.Phase)
	out.Conditions = *(*[]CarpCondition)(unsafe.Pointer(&in.Conditions))
	out.Message = in.Message
	out.Reason = in.Reason
	out.HostIP = in.HostIP
	out.CarpIP = in.CarpIP
	out.StartTime = (*meta_v1.Time)(unsafe.Pointer(in.StartTime))
	return nil
}

// Convert_testapigroup_CarpStatus_To_v1_CarpStatus is an autogenerated conversion function.
func Convert_testapigroup_CarpStatus_To_v1_CarpStatus(in *testapigroup.CarpStatus, out *CarpStatus, s conversion.Scope) error {
	return autoConvert_testapigroup_CarpStatus_To_v1_CarpStatus(in, out, s)
}
