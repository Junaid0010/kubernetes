/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	networkingv1beta1 "k8s.io/client-go/applyconfigurations/networking/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// IngressClassesGetter has a method to return a IngressClassInterface.
// A group's client should implement this interface.
type IngressClassesGetter interface {
	IngressClasses() IngressClassInterface
}

// IngressClassInterface has methods to work with IngressClass resources.
type IngressClassInterface interface {
	Create(ctx context.Context, ingressClass *v1beta1.IngressClass, opts v1.CreateOptions) (*v1beta1.IngressClass, error)
	Update(ctx context.Context, ingressClass *v1beta1.IngressClass, opts v1.UpdateOptions) (*v1beta1.IngressClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.IngressClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.IngressClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IngressClass, err error)
	Apply(ctx context.Context, ingressClass *networkingv1beta1.IngressClassApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.IngressClass, err error)
	IngressClassExpansion
}

// ingressClasses implements IngressClassInterface
type ingressClasses struct {
	*gentype.ClientWithListAndApply[*v1beta1.IngressClass, *v1beta1.IngressClassList, *networkingv1beta1.IngressClassApplyConfiguration]
}

// newIngressClasses returns a IngressClasses
func newIngressClasses(c *NetworkingV1beta1Client) *ingressClasses {
	return &ingressClasses{
		gentype.NewClientWithListAndApply[*v1beta1.IngressClass, *v1beta1.IngressClassList, *networkingv1beta1.IngressClassApplyConfiguration](
			"ingressclasses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.IngressClass { return &v1beta1.IngressClass{} },
			func() *v1beta1.IngressClassList { return &v1beta1.IngressClassList{} },
			gentype.PrefersProtobuf[*v1beta1.IngressClass](),
		),
	}
}
