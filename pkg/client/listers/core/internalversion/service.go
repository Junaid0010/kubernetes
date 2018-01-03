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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/lister-gen --input-dirs k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/admission,k8s.io/kubernetes/pkg/apis/admissionregistration,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/networking,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/scheduling,k8s.io/kubernetes/pkg/apis/settings,k8s.io/kubernetes/pkg/apis/storage

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	core "k8s.io/kubernetes/pkg/apis/core"
)

// ServiceLister helps list Services.
type ServiceLister interface {
	// List lists all Services in the indexer.
	List(selector labels.Selector) (ret []*core.Service, err error)
	// Services returns an object that can list and get Services.
	Services(namespace string) ServiceNamespaceLister
	ServiceListerExpansion
}

// serviceLister implements the ServiceLister interface.
type serviceLister struct {
	indexer cache.Indexer
}

// NewServiceLister returns a new ServiceLister.
func NewServiceLister(indexer cache.Indexer) ServiceLister {
	return &serviceLister{indexer: indexer}
}

// List lists all Services in the indexer.
func (s *serviceLister) List(selector labels.Selector) (ret []*core.Service, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*core.Service))
	})
	return ret, err
}

// Services returns an object that can list and get Services.
func (s *serviceLister) Services(namespace string) ServiceNamespaceLister {
	return serviceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceNamespaceLister helps list and get Services.
type ServiceNamespaceLister interface {
	// List lists all Services in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*core.Service, err error)
	// Get retrieves the Service from the indexer for a given namespace and name.
	Get(name string) (*core.Service, error)
	ServiceNamespaceListerExpansion
}

// serviceNamespaceLister implements the ServiceNamespaceLister
// interface.
type serviceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Services in the indexer for a given namespace.
func (s serviceNamespaceLister) List(selector labels.Selector) (ret []*core.Service, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*core.Service))
	})
	return ret, err
}

// Get retrieves the Service from the indexer for a given namespace and name.
func (s serviceNamespaceLister) Get(name string) (*core.Service, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(core.Resource("service"), name)
	}
	return obj.(*core.Service), nil
}
