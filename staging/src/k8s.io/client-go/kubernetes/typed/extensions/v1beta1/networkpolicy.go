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

	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	extensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// NetworkPoliciesGetter has a method to return a NetworkPolicyInterface.
// A group's client should implement this interface.
type NetworkPoliciesGetter interface {
	NetworkPolicies(namespace string) NetworkPolicyInterface
}

// NetworkPolicyInterface has methods to work with NetworkPolicy resources.
type NetworkPolicyInterface interface {
	Create(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.CreateOptions) (*v1beta1.NetworkPolicy, error)
	Update(ctx context.Context, networkPolicy *v1beta1.NetworkPolicy, opts v1.UpdateOptions) (*v1beta1.NetworkPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.NetworkPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.NetworkPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkPolicy, err error)
	Apply(ctx context.Context, networkPolicy *extensionsv1beta1.NetworkPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.NetworkPolicy, err error)
	NetworkPolicyExpansion
}

// networkPolicies implements NetworkPolicyInterface
type networkPolicies struct {
	*gentype.ClientWithListAndApply[*v1beta1.NetworkPolicy, *v1beta1.NetworkPolicyList, *extensionsv1beta1.NetworkPolicyApplyConfiguration]
}

// newNetworkPolicies returns a NetworkPolicies
func newNetworkPolicies(c *ExtensionsV1beta1Client, namespace string) *networkPolicies {
	return &networkPolicies{
		gentype.NewClientWithListAndApply[*v1beta1.NetworkPolicy, *v1beta1.NetworkPolicyList, *extensionsv1beta1.NetworkPolicyApplyConfiguration](
			"networkpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.NetworkPolicy { return &v1beta1.NetworkPolicy{} },
			func() *v1beta1.NetworkPolicyList { return &v1beta1.NetworkPolicyList{} },
			true,
		),
	}
}
