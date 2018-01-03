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

package v1beta1

import (
	rest "k8s.io/client-go/rest"
)

// ScalesGetter has a method to return a ScaleInterface.
// A group's client should implement this interface.
type ScalesGetter interface {
	Scales(namespace string) ScaleInterface
}

// ScaleInterface has methods to work with Scale resources.
type ScaleInterface interface {
	ScaleExpansion
}

// scales implements ScaleInterface
type scales struct {
	client rest.Interface
	ns     string
}

// newScales returns a Scales
func newScales(c *AppsV1beta1Client, namespace string) *scales {
	return &scales{
		client: c.RESTClient(),
		ns:     namespace,
	}
}
