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
	v1 "k8s.io/code-generator/examples/crd/apis/example2/v1"
	example2v1 "k8s.io/code-generator/examples/crd/applyconfiguration/example2/v1"
	scheme "k8s.io/code-generator/examples/crd/clientset/versioned/scheme"
)

// TestTypesGetter has a method to return a TestTypeInterface.
// A group's client should implement this interface.
type TestTypesGetter interface {
	TestTypes(namespace string) TestTypeInterface
}

// TestTypeInterface has methods to work with TestType resources.
type TestTypeInterface interface {
	Create(ctx context.Context, testType *v1.TestType, opts metav1.CreateOptions) (*v1.TestType, error)
	Update(ctx context.Context, testType *v1.TestType, opts metav1.UpdateOptions) (*v1.TestType, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, testType *v1.TestType, opts metav1.UpdateOptions) (*v1.TestType, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TestType, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TestTypeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TestType, err error)
	Apply(ctx context.Context, testType *example2v1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.TestType, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, testType *example2v1.TestTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.TestType, err error)
	TestTypeExpansion
}

// testTypes implements TestTypeInterface
type testTypes struct {
	*gentype.ClientWithListAndApply[*v1.TestType, *v1.TestTypeList, *example2v1.TestTypeApplyConfiguration]
}

// newTestTypes returns a TestTypes
func newTestTypes(c *SecondExampleV1Client, namespace string) *testTypes {
	return &testTypes{
		gentype.NewClientWithListAndApply[*v1.TestType, *v1.TestTypeList, *example2v1.TestTypeApplyConfiguration](
			"testtypes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.TestType { return &v1.TestType{} },
			func() *v1.TestTypeList { return &v1.TestTypeList{} },
			false,
		),
	}
}
