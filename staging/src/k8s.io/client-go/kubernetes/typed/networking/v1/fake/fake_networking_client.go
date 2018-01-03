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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/client-gen --output-base /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/vendor --clientset-path=k8s.io/client-go --clientset-name=kubernetes --input-base=k8s.io/kubernetes/vendor/k8s.io/api --input=core/v1,admissionregistration/v1alpha1,admissionregistration/v1beta1,apps/v1beta1,apps/v1beta2,apps/v1,authentication/v1,authentication/v1beta1,authorization/v1,authorization/v1beta1,autoscaling/v1,autoscaling/v2beta1,batch/v1,batch/v1beta1,batch/v2alpha1,certificates/v1beta1,extensions/v1beta1,events/v1beta1,networking/v1,policy/v1beta1,rbac/v1,rbac/v1beta1,rbac/v1alpha1,scheduling/v1alpha1,settings/v1alpha1,storage/v1beta1,storage/v1,storage/v1alpha1

package fake

import (
	v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkingV1 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1) NetworkPolicies(namespace string) v1.NetworkPolicyInterface {
	return &FakeNetworkPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
