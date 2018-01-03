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

// This file was automatically generated by: _output/bin/conversion-gen --extra-peer-dirs k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/api/core/v1 --v 1 --logtostderr -i k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission/v1beta1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1alpha1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1beta2,k8s.io/kubernetes/pkg/apis/authentication/v1,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization/v1,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/autoscaling/v2beta1,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v1beta1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates/v1beta1,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/kubernetes/pkg/apis/events/v1beta1,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/networking/v1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac/v1,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac/v1beta1,k8s.io/kubernetes/pkg/apis/scheduling/v1alpha1,k8s.io/kubernetes/pkg/apis/settings/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1,k8s.io/kubernetes/pkg/apis/storage/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1beta1,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/appsv1beta1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/appsv1beta2,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/autoscalingv1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/extensionsv1beta1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1 -O zz_generated.conversion

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v2beta1

import (
	v2beta1 "k8s.io/api/autoscaling/v2beta1"
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	core "k8s.io/kubernetes/pkg/apis/core"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference,
		Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference,
		Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler,
		Convert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler,
		Convert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition,
		Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition,
		Convert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList,
		Convert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList,
		Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec,
		Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec,
		Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus,
		Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus,
		Convert_v2beta1_MetricSpec_To_autoscaling_MetricSpec,
		Convert_autoscaling_MetricSpec_To_v2beta1_MetricSpec,
		Convert_v2beta1_MetricStatus_To_autoscaling_MetricStatus,
		Convert_autoscaling_MetricStatus_To_v2beta1_MetricStatus,
		Convert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource,
		Convert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource,
		Convert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus,
		Convert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus,
		Convert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource,
		Convert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource,
		Convert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus,
		Convert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus,
		Convert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource,
		Convert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource,
		Convert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus,
		Convert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus,
	)
}

func autoConvert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v2beta1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

// Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference is an autogenerated conversion function.
func Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in *v2beta1.CrossVersionObjectReference, out *autoscaling.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(in, out, s)
}

func autoConvert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v2beta1.CrossVersionObjectReference, s conversion.Scope) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	return nil
}

// Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference is an autogenerated conversion function.
func Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(in *autoscaling.CrossVersionObjectReference, out *v2beta1.CrossVersionObjectReference, s conversion.Scope) error {
	return autoConvert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(in, out, s)
}

func autoConvert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *v2beta1.HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in *v2beta1.HorizontalPodAutoscaler, out *autoscaling.HorizontalPodAutoscaler, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscaler_To_autoscaling_HorizontalPodAutoscaler(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *v2beta1.HorizontalPodAutoscaler, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler(in *autoscaling.HorizontalPodAutoscaler, out *v2beta1.HorizontalPodAutoscaler, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscaler_To_v2beta1_HorizontalPodAutoscaler(in, out, s)
}

func autoConvert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v2beta1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	out.Type = autoscaling.HorizontalPodAutoscalerConditionType(in.Type)
	out.Status = autoscaling.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in *v2beta1.HorizontalPodAutoscalerCondition, out *autoscaling.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerCondition_To_autoscaling_HorizontalPodAutoscalerCondition(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v2beta1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	out.Type = v2beta1.HorizontalPodAutoscalerConditionType(in.Type)
	out.Status = v1.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition(in *autoscaling.HorizontalPodAutoscalerCondition, out *v2beta1.HorizontalPodAutoscalerCondition, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerCondition_To_v2beta1_HorizontalPodAutoscalerCondition(in, out, s)
}

func autoConvert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v2beta1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]autoscaling.HorizontalPodAutoscaler)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in *v2beta1.HorizontalPodAutoscalerList, out *autoscaling.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerList_To_autoscaling_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v2beta1.HorizontalPodAutoscalerList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v2beta1.HorizontalPodAutoscaler)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList(in *autoscaling.HorizontalPodAutoscalerList, out *v2beta1.HorizontalPodAutoscalerList, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerList_To_v2beta1_HorizontalPodAutoscalerList(in, out, s)
}

func autoConvert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *v2beta1.HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	out.Metrics = *(*[]autoscaling.MetricSpec)(unsafe.Pointer(&in.Metrics))
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in *v2beta1.HorizontalPodAutoscalerSpec, out *autoscaling.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerSpec_To_autoscaling_HorizontalPodAutoscalerSpec(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *v2beta1.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(&in.ScaleTargetRef, &out.ScaleTargetRef, s); err != nil {
		return err
	}
	out.MinReplicas = (*int32)(unsafe.Pointer(in.MinReplicas))
	out.MaxReplicas = in.MaxReplicas
	out.Metrics = *(*[]v2beta1.MetricSpec)(unsafe.Pointer(&in.Metrics))
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec(in *autoscaling.HorizontalPodAutoscalerSpec, out *v2beta1.HorizontalPodAutoscalerSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerSpec_To_v2beta1_HorizontalPodAutoscalerSpec(in, out, s)
}

func autoConvert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *v2beta1.HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	out.CurrentMetrics = *(*[]autoscaling.MetricStatus)(unsafe.Pointer(&in.CurrentMetrics))
	out.Conditions = *(*[]autoscaling.HorizontalPodAutoscalerCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus is an autogenerated conversion function.
func Convert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in *v2beta1.HorizontalPodAutoscalerStatus, out *autoscaling.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	return autoConvert_v2beta1_HorizontalPodAutoscalerStatus_To_autoscaling_HorizontalPodAutoscalerStatus(in, out, s)
}

func autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *v2beta1.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	out.ObservedGeneration = (*int64)(unsafe.Pointer(in.ObservedGeneration))
	out.LastScaleTime = (*meta_v1.Time)(unsafe.Pointer(in.LastScaleTime))
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	out.CurrentMetrics = *(*[]v2beta1.MetricStatus)(unsafe.Pointer(&in.CurrentMetrics))
	out.Conditions = *(*[]v2beta1.HorizontalPodAutoscalerCondition)(unsafe.Pointer(&in.Conditions))
	return nil
}

// Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus is an autogenerated conversion function.
func Convert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(in *autoscaling.HorizontalPodAutoscalerStatus, out *v2beta1.HorizontalPodAutoscalerStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_HorizontalPodAutoscalerStatus_To_v2beta1_HorizontalPodAutoscalerStatus(in, out, s)
}

func autoConvert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(in *v2beta1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	out.Object = (*autoscaling.ObjectMetricSource)(unsafe.Pointer(in.Object))
	out.Pods = (*autoscaling.PodsMetricSource)(unsafe.Pointer(in.Pods))
	out.Resource = (*autoscaling.ResourceMetricSource)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_v2beta1_MetricSpec_To_autoscaling_MetricSpec is an autogenerated conversion function.
func Convert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(in *v2beta1.MetricSpec, out *autoscaling.MetricSpec, s conversion.Scope) error {
	return autoConvert_v2beta1_MetricSpec_To_autoscaling_MetricSpec(in, out, s)
}

func autoConvert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(in *autoscaling.MetricSpec, out *v2beta1.MetricSpec, s conversion.Scope) error {
	out.Type = v2beta1.MetricSourceType(in.Type)
	out.Object = (*v2beta1.ObjectMetricSource)(unsafe.Pointer(in.Object))
	out.Pods = (*v2beta1.PodsMetricSource)(unsafe.Pointer(in.Pods))
	out.Resource = (*v2beta1.ResourceMetricSource)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_autoscaling_MetricSpec_To_v2beta1_MetricSpec is an autogenerated conversion function.
func Convert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(in *autoscaling.MetricSpec, out *v2beta1.MetricSpec, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricSpec_To_v2beta1_MetricSpec(in, out, s)
}

func autoConvert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(in *v2beta1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	out.Type = autoscaling.MetricSourceType(in.Type)
	out.Object = (*autoscaling.ObjectMetricStatus)(unsafe.Pointer(in.Object))
	out.Pods = (*autoscaling.PodsMetricStatus)(unsafe.Pointer(in.Pods))
	out.Resource = (*autoscaling.ResourceMetricStatus)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_v2beta1_MetricStatus_To_autoscaling_MetricStatus is an autogenerated conversion function.
func Convert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(in *v2beta1.MetricStatus, out *autoscaling.MetricStatus, s conversion.Scope) error {
	return autoConvert_v2beta1_MetricStatus_To_autoscaling_MetricStatus(in, out, s)
}

func autoConvert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(in *autoscaling.MetricStatus, out *v2beta1.MetricStatus, s conversion.Scope) error {
	out.Type = v2beta1.MetricSourceType(in.Type)
	out.Object = (*v2beta1.ObjectMetricStatus)(unsafe.Pointer(in.Object))
	out.Pods = (*v2beta1.PodsMetricStatus)(unsafe.Pointer(in.Pods))
	out.Resource = (*v2beta1.ResourceMetricStatus)(unsafe.Pointer(in.Resource))
	return nil
}

// Convert_autoscaling_MetricStatus_To_v2beta1_MetricStatus is an autogenerated conversion function.
func Convert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(in *autoscaling.MetricStatus, out *v2beta1.MetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_MetricStatus_To_v2beta1_MetricStatus(in, out, s)
}

func autoConvert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *v2beta1.ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	if err := Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.TargetValue = in.TargetValue
	return nil
}

// Convert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource is an autogenerated conversion function.
func Convert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in *v2beta1.ObjectMetricSource, out *autoscaling.ObjectMetricSource, s conversion.Scope) error {
	return autoConvert_v2beta1_ObjectMetricSource_To_autoscaling_ObjectMetricSource(in, out, s)
}

func autoConvert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *v2beta1.ObjectMetricSource, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.TargetValue = in.TargetValue
	return nil
}

// Convert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource is an autogenerated conversion function.
func Convert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource(in *autoscaling.ObjectMetricSource, out *v2beta1.ObjectMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_ObjectMetricSource_To_v2beta1_ObjectMetricSource(in, out, s)
}

func autoConvert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *v2beta1.ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	if err := Convert_v2beta1_CrossVersionObjectReference_To_autoscaling_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.CurrentValue = in.CurrentValue
	return nil
}

// Convert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus is an autogenerated conversion function.
func Convert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in *v2beta1.ObjectMetricStatus, out *autoscaling.ObjectMetricStatus, s conversion.Scope) error {
	return autoConvert_v2beta1_ObjectMetricStatus_To_autoscaling_ObjectMetricStatus(in, out, s)
}

func autoConvert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *v2beta1.ObjectMetricStatus, s conversion.Scope) error {
	if err := Convert_autoscaling_CrossVersionObjectReference_To_v2beta1_CrossVersionObjectReference(&in.Target, &out.Target, s); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.CurrentValue = in.CurrentValue
	return nil
}

// Convert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus is an autogenerated conversion function.
func Convert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus(in *autoscaling.ObjectMetricStatus, out *v2beta1.ObjectMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ObjectMetricStatus_To_v2beta1_ObjectMetricStatus(in, out, s)
}

func autoConvert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *v2beta1.PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.TargetAverageValue = in.TargetAverageValue
	return nil
}

// Convert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource is an autogenerated conversion function.
func Convert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource(in *v2beta1.PodsMetricSource, out *autoscaling.PodsMetricSource, s conversion.Scope) error {
	return autoConvert_v2beta1_PodsMetricSource_To_autoscaling_PodsMetricSource(in, out, s)
}

func autoConvert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *v2beta1.PodsMetricSource, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.TargetAverageValue = in.TargetAverageValue
	return nil
}

// Convert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource is an autogenerated conversion function.
func Convert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource(in *autoscaling.PodsMetricSource, out *v2beta1.PodsMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_PodsMetricSource_To_v2beta1_PodsMetricSource(in, out, s)
}

func autoConvert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *v2beta1.PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus is an autogenerated conversion function.
func Convert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in *v2beta1.PodsMetricStatus, out *autoscaling.PodsMetricStatus, s conversion.Scope) error {
	return autoConvert_v2beta1_PodsMetricStatus_To_autoscaling_PodsMetricStatus(in, out, s)
}

func autoConvert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *v2beta1.PodsMetricStatus, s conversion.Scope) error {
	out.MetricName = in.MetricName
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus is an autogenerated conversion function.
func Convert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus(in *autoscaling.PodsMetricStatus, out *v2beta1.PodsMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_PodsMetricStatus_To_v2beta1_PodsMetricStatus(in, out, s)
}

func autoConvert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *v2beta1.ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	out.Name = core.ResourceName(in.Name)
	out.TargetAverageUtilization = (*int32)(unsafe.Pointer(in.TargetAverageUtilization))
	out.TargetAverageValue = (*resource.Quantity)(unsafe.Pointer(in.TargetAverageValue))
	return nil
}

// Convert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource is an autogenerated conversion function.
func Convert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in *v2beta1.ResourceMetricSource, out *autoscaling.ResourceMetricSource, s conversion.Scope) error {
	return autoConvert_v2beta1_ResourceMetricSource_To_autoscaling_ResourceMetricSource(in, out, s)
}

func autoConvert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *v2beta1.ResourceMetricSource, s conversion.Scope) error {
	out.Name = v1.ResourceName(in.Name)
	out.TargetAverageUtilization = (*int32)(unsafe.Pointer(in.TargetAverageUtilization))
	out.TargetAverageValue = (*resource.Quantity)(unsafe.Pointer(in.TargetAverageValue))
	return nil
}

// Convert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource is an autogenerated conversion function.
func Convert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource(in *autoscaling.ResourceMetricSource, out *v2beta1.ResourceMetricSource, s conversion.Scope) error {
	return autoConvert_autoscaling_ResourceMetricSource_To_v2beta1_ResourceMetricSource(in, out, s)
}

func autoConvert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *v2beta1.ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	out.Name = core.ResourceName(in.Name)
	out.CurrentAverageUtilization = (*int32)(unsafe.Pointer(in.CurrentAverageUtilization))
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus is an autogenerated conversion function.
func Convert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in *v2beta1.ResourceMetricStatus, out *autoscaling.ResourceMetricStatus, s conversion.Scope) error {
	return autoConvert_v2beta1_ResourceMetricStatus_To_autoscaling_ResourceMetricStatus(in, out, s)
}

func autoConvert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *v2beta1.ResourceMetricStatus, s conversion.Scope) error {
	out.Name = v1.ResourceName(in.Name)
	out.CurrentAverageUtilization = (*int32)(unsafe.Pointer(in.CurrentAverageUtilization))
	out.CurrentAverageValue = in.CurrentAverageValue
	return nil
}

// Convert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus is an autogenerated conversion function.
func Convert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus(in *autoscaling.ResourceMetricStatus, out *v2beta1.ResourceMetricStatus, s conversion.Scope) error {
	return autoConvert_autoscaling_ResourceMetricStatus_To_v2beta1_ResourceMetricStatus(in, out, s)
}
