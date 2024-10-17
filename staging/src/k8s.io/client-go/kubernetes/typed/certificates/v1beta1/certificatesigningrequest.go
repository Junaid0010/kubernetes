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

	v1beta1 "k8s.io/api/certificates/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	certificatesv1beta1 "k8s.io/client-go/applyconfigurations/certificates/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CertificateSigningRequestsGetter has a method to return a CertificateSigningRequestInterface.
// A group's client should implement this interface.
type CertificateSigningRequestsGetter interface {
	CertificateSigningRequests() CertificateSigningRequestInterface
}

// CertificateSigningRequestInterface has methods to work with CertificateSigningRequest resources.
type CertificateSigningRequestInterface interface {
	Create(ctx context.Context, certificateSigningRequest *v1beta1.CertificateSigningRequest, opts v1.CreateOptions) (*v1beta1.CertificateSigningRequest, error)
	Update(ctx context.Context, certificateSigningRequest *v1beta1.CertificateSigningRequest, opts v1.UpdateOptions) (*v1beta1.CertificateSigningRequest, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, certificateSigningRequest *v1beta1.CertificateSigningRequest, opts v1.UpdateOptions) (*v1beta1.CertificateSigningRequest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.CertificateSigningRequest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CertificateSigningRequestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CertificateSigningRequest, err error)
	Apply(ctx context.Context, certificateSigningRequest *certificatesv1beta1.CertificateSigningRequestApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CertificateSigningRequest, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, certificateSigningRequest *certificatesv1beta1.CertificateSigningRequestApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.CertificateSigningRequest, err error)
	CertificateSigningRequestExpansion
}

// certificateSigningRequests implements CertificateSigningRequestInterface
type certificateSigningRequests struct {
	*gentype.ClientWithListAndApply[*v1beta1.CertificateSigningRequest, *v1beta1.CertificateSigningRequestList, *certificatesv1beta1.CertificateSigningRequestApplyConfiguration]
}

// newCertificateSigningRequests returns a CertificateSigningRequests
func newCertificateSigningRequests(c *CertificatesV1beta1Client) *certificateSigningRequests {
	return &certificateSigningRequests{
		gentype.NewClientWithListAndApply[*v1beta1.CertificateSigningRequest, *v1beta1.CertificateSigningRequestList, *certificatesv1beta1.CertificateSigningRequestApplyConfiguration](
			"certificatesigningrequests",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.CertificateSigningRequest { return &v1beta1.CertificateSigningRequest{} },
			func() *v1beta1.CertificateSigningRequestList { return &v1beta1.CertificateSigningRequestList{} },
			gentype.PrefersProtobuf[*v1beta1.CertificateSigningRequest](),
		),
	}
}
