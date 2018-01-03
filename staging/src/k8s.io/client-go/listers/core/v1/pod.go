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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/lister-gen --output-base /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/vendor --output-package k8s.io/client-go/listers --input-dirs k8s.io/api/admission/v1beta1,k8s.io/api/admissionregistration/v1alpha1,k8s.io/api/admissionregistration/v1beta1,k8s.io/api/apps/v1,k8s.io/api/apps/v1beta1,k8s.io/api/apps/v1beta2,k8s.io/api/authentication/v1,k8s.io/api/authentication/v1beta1,k8s.io/api/authorization/v1,k8s.io/api/authorization/v1beta1,k8s.io/api/autoscaling/v1,k8s.io/api/autoscaling/v2beta1,k8s.io/api/batch/v1,k8s.io/api/batch/v1beta1,k8s.io/api/batch/v2alpha1,k8s.io/api/certificates/v1beta1,k8s.io/api/core/v1,k8s.io/api/events/v1beta1,k8s.io/api/extensions/v1beta1,k8s.io/api/imagepolicy/v1alpha1,k8s.io/api/networking/v1,k8s.io/api/policy/v1beta1,k8s.io/api/rbac/v1,k8s.io/api/rbac/v1alpha1,k8s.io/api/rbac/v1beta1,k8s.io/api/scheduling/v1alpha1,k8s.io/api/settings/v1alpha1,k8s.io/api/storage/v1,k8s.io/api/storage/v1alpha1,k8s.io/api/storage/v1beta1

// This file was automatically generated by lister-gen

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodLister helps list Pods.
type PodLister interface {
	// List lists all Pods in the indexer.
	List(selector labels.Selector) (ret []*v1.Pod, err error)
	// Pods returns an object that can list and get Pods.
	Pods(namespace string) PodNamespaceLister
	PodListerExpansion
}

// podLister implements the PodLister interface.
type podLister struct {
	indexer cache.Indexer
}

// NewPodLister returns a new PodLister.
func NewPodLister(indexer cache.Indexer) PodLister {
	return &podLister{indexer: indexer}
}

// List lists all Pods in the indexer.
func (s *podLister) List(selector labels.Selector) (ret []*v1.Pod, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pod))
	})
	return ret, err
}

// Pods returns an object that can list and get Pods.
func (s *podLister) Pods(namespace string) PodNamespaceLister {
	return podNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodNamespaceLister helps list and get Pods.
type PodNamespaceLister interface {
	// List lists all Pods in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Pod, err error)
	// Get retrieves the Pod from the indexer for a given namespace and name.
	Get(name string) (*v1.Pod, error)
	PodNamespaceListerExpansion
}

// podNamespaceLister implements the PodNamespaceLister
// interface.
type podNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Pods in the indexer for a given namespace.
func (s podNamespaceLister) List(selector labels.Selector) (ret []*v1.Pod, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pod))
	})
	return ret, err
}

// Get retrieves the Pod from the indexer for a given namespace and name.
func (s podNamespaceLister) Get(name string) (*v1.Pod, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pod"), name)
	}
	return obj.(*v1.Pod), nil
}
