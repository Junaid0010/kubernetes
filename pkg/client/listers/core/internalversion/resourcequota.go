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

// ResourceQuotaLister helps list ResourceQuotas.
type ResourceQuotaLister interface {
	// List lists all ResourceQuotas in the indexer.
	List(selector labels.Selector) (ret []*core.ResourceQuota, err error)
	// ResourceQuotas returns an object that can list and get ResourceQuotas.
	ResourceQuotas(namespace string) ResourceQuotaNamespaceLister
	ResourceQuotaListerExpansion
}

// resourceQuotaLister implements the ResourceQuotaLister interface.
type resourceQuotaLister struct {
	indexer cache.Indexer
}

// NewResourceQuotaLister returns a new ResourceQuotaLister.
func NewResourceQuotaLister(indexer cache.Indexer) ResourceQuotaLister {
	return &resourceQuotaLister{indexer: indexer}
}

// List lists all ResourceQuotas in the indexer.
func (s *resourceQuotaLister) List(selector labels.Selector) (ret []*core.ResourceQuota, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*core.ResourceQuota))
	})
	return ret, err
}

// ResourceQuotas returns an object that can list and get ResourceQuotas.
func (s *resourceQuotaLister) ResourceQuotas(namespace string) ResourceQuotaNamespaceLister {
	return resourceQuotaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceQuotaNamespaceLister helps list and get ResourceQuotas.
type ResourceQuotaNamespaceLister interface {
	// List lists all ResourceQuotas in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*core.ResourceQuota, err error)
	// Get retrieves the ResourceQuota from the indexer for a given namespace and name.
	Get(name string) (*core.ResourceQuota, error)
	ResourceQuotaNamespaceListerExpansion
}

// resourceQuotaNamespaceLister implements the ResourceQuotaNamespaceLister
// interface.
type resourceQuotaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceQuotas in the indexer for a given namespace.
func (s resourceQuotaNamespaceLister) List(selector labels.Selector) (ret []*core.ResourceQuota, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*core.ResourceQuota))
	})
	return ret, err
}

// Get retrieves the ResourceQuota from the indexer for a given namespace and name.
func (s resourceQuotaNamespaceLister) Get(name string) (*core.ResourceQuota, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(core.Resource("resourcequota"), name)
	}
	return obj.(*core.ResourceQuota), nil
}
