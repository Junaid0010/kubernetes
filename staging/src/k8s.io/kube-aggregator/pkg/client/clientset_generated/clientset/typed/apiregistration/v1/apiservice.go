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

package v1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"
	scheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
)

// APIServicesGetter has a method to return a APIServiceInterface.
// A group's client should implement this interface.
type APIServicesGetter interface {
	APIServices() APIServiceInterface
}

// APIServiceInterface has methods to work with APIService resources.
type APIServiceInterface interface {
	Create(ctx context.Context, aPIService *v1.APIService, opts metav1.CreateOptions) (*v1.APIService, error)
	Update(ctx context.Context, aPIService *v1.APIService, opts metav1.UpdateOptions) (*v1.APIService, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, aPIService *v1.APIService, opts metav1.UpdateOptions) (*v1.APIService, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.APIService, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.APIServiceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.APIService, err error)
	APIServiceExpansion
}

// aPIServices implements APIServiceInterface
type aPIServices struct {
	*gentype.ClientWithList[*v1.APIService, *v1.APIServiceList]
}

// newAPIServices returns a APIServices
func newAPIServices(c *ApiregistrationV1Client) *aPIServices {
	return &aPIServices{
		gentype.NewClientWithList[*v1.APIService, *v1.APIServiceList](
			"apiservices",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.APIService { return &v1.APIService{} },
			func() *v1.APIServiceList { return &v1.APIServiceList{} },
			gentype.PrefersProtobuf[*v1.APIService](),
		),
	}
}
