// +build !ignore_autogenerated

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

// This file was automatically generated by: _output/bin/deepcopy-gen --v 1 --logtostderr -i k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm,k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/cmd/kubeadm/app/phases/etcd/spec,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission,k8s.io/kubernetes/pkg/apis/admissionregistration,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/networking,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/scheduling,k8s.io/kubernetes/pkg/apis/settings,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/controller/garbagecollector/metaonly,k8s.io/kubernetes/pkg/kubectl/cmd/testing,k8s.io/kubernetes/pkg/kubectl/testing,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/pkg/registry/rbac/reconciliation,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/plugin/pkg/scheduler/api,k8s.io/kubernetes/plugin/pkg/scheduler/api/v1,k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/admissionregistration/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/admissionregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1beta2,k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1,k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/authorization/v1,k8s.io/kubernetes/vendor/k8s.io/api/authorization/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/autoscaling/v1,k8s.io/kubernetes/vendor/k8s.io/api/autoscaling/v2beta1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v2alpha1,k8s.io/kubernetes/vendor/k8s.io/api/certificates/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/core/v1,k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/extensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/imagepolicy/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/networking/v1,k8s.io/kubernetes/vendor/k8s.io/api/policy/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/settings/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/labels,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/testing,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/testing,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/test,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/watch,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/endpoints/openapi/testing,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/endpoints/testing,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/registry/rest,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/storage/testing,k8s.io/kubernetes/vendor/k8s.io/client-go/rest,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme,k8s.io/kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api,k8s.io/kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/crd/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/crd/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1 --bounding-dirs k8s.io/kubernetes,k8s.io/api -O zz_generated.deepcopy

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroup) DeepCopyInto(out *APIGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]GroupVersionForDiscovery, len(*in))
		copy(*out, *in)
	}
	out.PreferredVersion = in.PreferredVersion
	if in.ServerAddressByClientCIDRs != nil {
		in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
		*out = make([]ServerAddressByClientCIDR, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroup.
func (in *APIGroup) DeepCopy() *APIGroup {
	if in == nil {
		return nil
	}
	out := new(APIGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroupList) DeepCopyInto(out *APIGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]APIGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroupList.
func (in *APIGroupList) DeepCopy() *APIGroupList {
	if in == nil {
		return nil
	}
	out := new(APIGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResource) DeepCopyInto(out *APIResource) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make(Verbs, len(*in))
		copy(*out, *in)
	}
	if in.ShortNames != nil {
		in, out := &in.ShortNames, &out.ShortNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResource.
func (in *APIResource) DeepCopy() *APIResource {
	if in == nil {
		return nil
	}
	out := new(APIResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIResourceList) DeepCopyInto(out *APIResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.APIResources != nil {
		in, out := &in.APIResources, &out.APIResources
		*out = make([]APIResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIResourceList.
func (in *APIResourceList) DeepCopy() *APIResourceList {
	if in == nil {
		return nil
	}
	out := new(APIResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIVersions) DeepCopyInto(out *APIVersions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServerAddressByClientCIDRs != nil {
		in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
		*out = make([]ServerAddressByClientCIDR, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIVersions.
func (in *APIVersions) DeepCopy() *APIVersions {
	if in == nil {
		return nil
	}
	out := new(APIVersions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIVersions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteOptions) DeepCopyInto(out *DeleteOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.GracePeriodSeconds != nil {
		in, out := &in.GracePeriodSeconds, &out.GracePeriodSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.Preconditions != nil {
		in, out := &in.Preconditions, &out.Preconditions
		if *in == nil {
			*out = nil
		} else {
			*out = new(Preconditions)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.OrphanDependents != nil {
		in, out := &in.OrphanDependents, &out.OrphanDependents
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.PropagationPolicy != nil {
		in, out := &in.PropagationPolicy, &out.PropagationPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(DeletionPropagation)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteOptions.
func (in *DeleteOptions) DeepCopy() *DeleteOptions {
	if in == nil {
		return nil
	}
	out := new(DeleteOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeleteOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Duration) DeepCopyInto(out *Duration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportOptions) DeepCopyInto(out *ExportOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportOptions.
func (in *ExportOptions) DeepCopy() *ExportOptions {
	if in == nil {
		return nil
	}
	out := new(ExportOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExportOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GetOptions) DeepCopyInto(out *GetOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetOptions.
func (in *GetOptions) DeepCopy() *GetOptions {
	if in == nil {
		return nil
	}
	out := new(GetOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GetOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupKind) DeepCopyInto(out *GroupKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupKind.
func (in *GroupKind) DeepCopy() *GroupKind {
	if in == nil {
		return nil
	}
	out := new(GroupKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersion) DeepCopyInto(out *GroupVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersion.
func (in *GroupVersion) DeepCopy() *GroupVersion {
	if in == nil {
		return nil
	}
	out := new(GroupVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionForDiscovery) DeepCopyInto(out *GroupVersionForDiscovery) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionForDiscovery.
func (in *GroupVersionForDiscovery) DeepCopy() *GroupVersionForDiscovery {
	if in == nil {
		return nil
	}
	out := new(GroupVersionForDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionKind) DeepCopyInto(out *GroupVersionKind) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionKind.
func (in *GroupVersionKind) DeepCopy() *GroupVersionKind {
	if in == nil {
		return nil
	}
	out := new(GroupVersionKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupVersionResource) DeepCopyInto(out *GroupVersionResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResource.
func (in *GroupVersionResource) DeepCopy() *GroupVersionResource {
	if in == nil {
		return nil
	}
	out := new(GroupVersionResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Initializer) DeepCopyInto(out *Initializer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Initializer.
func (in *Initializer) DeepCopy() *Initializer {
	if in == nil {
		return nil
	}
	out := new(Initializer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Initializers) DeepCopyInto(out *Initializers) {
	*out = *in
	if in.Pending != nil {
		in, out := &in.Pending, &out.Pending
		*out = make([]Initializer, len(*in))
		copy(*out, *in)
	}
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		if *in == nil {
			*out = nil
		} else {
			*out = new(Status)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Initializers.
func (in *Initializers) DeepCopy() *Initializers {
	if in == nil {
		return nil
	}
	out := new(Initializers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalEvent) DeepCopyInto(out *InternalEvent) {
	*out = *in
	if in.Object == nil {
		out.Object = nil
	} else {
		out.Object = in.Object.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalEvent.
func (in *InternalEvent) DeepCopy() *InternalEvent {
	if in == nil {
		return nil
	}
	out := new(InternalEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelSelector) DeepCopyInto(out *LabelSelector) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]LabelSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelSelector.
func (in *LabelSelector) DeepCopy() *LabelSelector {
	if in == nil {
		return nil
	}
	out := new(LabelSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelSelectorRequirement) DeepCopyInto(out *LabelSelectorRequirement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelSelectorRequirement.
func (in *LabelSelectorRequirement) DeepCopy() *LabelSelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(LabelSelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *List) DeepCopyInto(out *List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]runtime.RawExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new List.
func (in *List) DeepCopy() *List {
	if in == nil {
		return nil
	}
	out := new(List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListMeta) DeepCopyInto(out *ListMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListMeta.
func (in *ListMeta) DeepCopy() *ListMeta {
	if in == nil {
		return nil
	}
	out := new(ListMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListOptions) DeepCopyInto(out *ListOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListOptions.
func (in *ListOptions) DeepCopy() *ListOptions {
	if in == nil {
		return nil
	}
	out := new(ListOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicroTime.
func (in *MicroTime) DeepCopy() *MicroTime {
	if in == nil {
		return nil
	}
	out := new(MicroTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	in.CreationTimestamp.DeepCopyInto(&out.CreationTimestamp)
	if in.DeletionTimestamp != nil {
		in, out := &in.DeletionTimestamp, &out.DeletionTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = new(Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.DeletionGracePeriodSeconds != nil {
		in, out := &in.DeletionGracePeriodSeconds, &out.DeletionGracePeriodSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]OwnerReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Initializers != nil {
		in, out := &in.Initializers, &out.Initializers
		if *in == nil {
			*out = nil
		} else {
			*out = new(Initializers)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerReference) DeepCopyInto(out *OwnerReference) {
	*out = *in
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.BlockOwnerDeletion != nil {
		in, out := &in.BlockOwnerDeletion, &out.BlockOwnerDeletion
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerReference.
func (in *OwnerReference) DeepCopy() *OwnerReference {
	if in == nil {
		return nil
	}
	out := new(OwnerReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Patch) DeepCopyInto(out *Patch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Patch.
func (in *Patch) DeepCopy() *Patch {
	if in == nil {
		return nil
	}
	out := new(Patch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preconditions) DeepCopyInto(out *Preconditions) {
	*out = *in
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		if *in == nil {
			*out = nil
		} else {
			*out = new(types.UID)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preconditions.
func (in *Preconditions) DeepCopy() *Preconditions {
	if in == nil {
		return nil
	}
	out := new(Preconditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootPaths) DeepCopyInto(out *RootPaths) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootPaths.
func (in *RootPaths) DeepCopy() *RootPaths {
	if in == nil {
		return nil
	}
	out := new(RootPaths)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerAddressByClientCIDR) DeepCopyInto(out *ServerAddressByClientCIDR) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerAddressByClientCIDR.
func (in *ServerAddressByClientCIDR) DeepCopy() *ServerAddressByClientCIDR {
	if in == nil {
		return nil
	}
	out := new(ServerAddressByClientCIDR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		if *in == nil {
			*out = nil
		} else {
			*out = new(StatusDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Status) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCause) DeepCopyInto(out *StatusCause) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCause.
func (in *StatusCause) DeepCopy() *StatusCause {
	if in == nil {
		return nil
	}
	out := new(StatusCause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusDetails) DeepCopyInto(out *StatusDetails) {
	*out = *in
	if in.Causes != nil {
		in, out := &in.Causes, &out.Causes
		*out = make([]StatusCause, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusDetails.
func (in *StatusDetails) DeepCopy() *StatusDetails {
	if in == nil {
		return nil
	}
	out := new(StatusDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Time.
func (in *Time) DeepCopy() *Time {
	if in == nil {
		return nil
	}
	out := new(Time)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timestamp) DeepCopyInto(out *Timestamp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timestamp.
func (in *Timestamp) DeepCopy() *Timestamp {
	if in == nil {
		return nil
	}
	out := new(Timestamp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatchEvent) DeepCopyInto(out *WatchEvent) {
	*out = *in
	in.Object.DeepCopyInto(&out.Object)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatchEvent.
func (in *WatchEvent) DeepCopy() *WatchEvent {
	if in == nil {
		return nil
	}
	out := new(WatchEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatchEvent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}
