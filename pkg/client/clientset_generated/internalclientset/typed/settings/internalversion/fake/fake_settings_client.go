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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/client-gen --input-base=k8s.io/kubernetes/pkg/apis --input=api/,admissionregistration/,apps/,authentication/,authorization/,autoscaling/,batch/,certificates/,extensions/,events/,networking/,policy/,rbac/,scheduling/,settings/,storage/

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	internalversion "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/settings/internalversion"
)

type FakeSettings struct {
	*testing.Fake
}

func (c *FakeSettings) PodPresets(namespace string) internalversion.PodPresetInterface {
	return &FakePodPresets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSettings) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
