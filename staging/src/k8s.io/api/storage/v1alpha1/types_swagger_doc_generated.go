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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_CSIStorageCapacity = map[string]string{
	"":                  "CSIStorageCapacity stores the result of one CSI GetCapacity call. For a given StorageClass, this describes the available capacity in a particular topology segment.  This can be used when considering where to instantiate new PersistentVolumes.\n\nFor example this can express things like:\n- StorageClass \"standard\" has \"1234 GiB\" available in \"topology.kubernetes.io/zone=us-east1\"\n- StorageClass \"localssd\" has \"10 GiB\" available in \"kubernetes.io/hostname=knode-abc123\"\n\nThe following three cases all imply that no capacity is available for a certain combination:\n- no object exists with suitable topology and storage class name\n- such an object exists, but the capacity is unset\n- such an object exists, but the capacity is zero\n\nThe producer of these objects can decide which approach is more suitable.\n\nThey are consumed by the kube-scheduler when a CSI driver opts into capacity-aware scheduling with CSIDriverSpec.StorageCapacity. The scheduler compares the MaximumVolumeSize against the requested size of pending volumes to filter out unsuitable nodes. If MaximumVolumeSize is unset, it falls back to a comparison against the less precise Capacity. If that is also unset, the scheduler assumes that capacity is insufficient and tries some other node.",
	"metadata":          "Standard object's metadata. The name has no particular meaning. It must be be a DNS subdomain (dots allowed, 253 characters). To ensure that there are no conflicts with other CSI drivers on the cluster, the recommendation is to use csisc-<uuid>, a generated name, or a reverse-domain name which ends with the unique CSI driver name.\n\nObjects are namespaced.\n\nMore info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"nodeTopology":      "nodeTopology defines which nodes have access to the storage for which capacity was reported. If not set, the storage is not accessible from any node in the cluster. If empty, the storage is accessible from all nodes. This field is immutable.",
	"storageClassName":  "storageClassName represents the name of the StorageClass that the reported capacity applies to. It must meet the same requirements as the name of a StorageClass object (non-empty, DNS subdomain). If that object no longer exists, the CSIStorageCapacity object is obsolete and should be removed by its creator. This field is immutable.",
	"capacity":          "capacity is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.\n\nThe semantic is currently (CSI spec 1.2) defined as: The available capacity, in bytes, of the storage that can be used to provision volumes. If not set, that information is currently unavailable.",
	"maximumVolumeSize": "maximumVolumeSize is the value reported by the CSI driver in its GetCapacityResponse for a GetCapacityRequest with topology and parameters that match the previous fields.\n\nThis is defined since CSI spec 1.4.0 as the largest size that may be used in a CreateVolumeRequest.capacity_range.required_bytes field to create a volume with the same parameters as those in GetCapacityRequest. The corresponding value in the Kubernetes API is ResourceRequirements.Requests in a volume claim.",
}

func (CSIStorageCapacity) SwaggerDoc() map[string]string {
	return map_CSIStorageCapacity
}

var map_CSIStorageCapacityList = map[string]string{
	"":         "CSIStorageCapacityList is a collection of CSIStorageCapacity objects.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of CSIStorageCapacity objects.",
}

func (CSIStorageCapacityList) SwaggerDoc() map[string]string {
	return map_CSIStorageCapacityList
}

var map_VolumeAttachment = map[string]string{
	"":         "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.\n\nVolumeAttachment objects are non-namespaced.",
	"metadata": "Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "spec represents specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.",
	"status":   "status represents status of the VolumeAttachment request. Populated by the entity completing the attach or detach operation, i.e. the external-attacher.",
}

func (VolumeAttachment) SwaggerDoc() map[string]string {
	return map_VolumeAttachment
}

var map_VolumeAttachmentList = map[string]string{
	"":         "VolumeAttachmentList is a collection of VolumeAttachment objects.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"items":    "items is the list of VolumeAttachments",
}

func (VolumeAttachmentList) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentList
}

var map_VolumeAttachmentSource = map[string]string{
	"":                     "VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.",
	"persistentVolumeName": "persistentVolumeName represents the name of the persistent volume to attach.",
}

func (VolumeAttachmentSource) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentSource
}

var map_VolumeAttachmentSpec = map[string]string{
	"":         "VolumeAttachmentSpec is the specification of a VolumeAttachment request.",
	"attacher": "attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().",
	"source":   "source represents the volume that should be attached.",
	"nodeName": "nodeName represents the node that the volume should be attached to.",
}

func (VolumeAttachmentSpec) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentSpec
}

var map_VolumeAttachmentStatus = map[string]string{
	"":                   "VolumeAttachmentStatus is the status of a VolumeAttachment request.",
	"attached":           "attached indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.",
	"attachmentMetadata": "attachmentMetadata is populated with any information returned by the attach operation, upon successful attach, that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.",
	"attachError":        "attachError represents the last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.",
	"detachError":        "detachError represents the last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.",
}

func (VolumeAttachmentStatus) SwaggerDoc() map[string]string {
	return map_VolumeAttachmentStatus
}

var map_VolumeError = map[string]string{
	"":        "VolumeError captures an error encountered during a volume operation.",
	"time":    "time represents the time the error was encountered.",
	"message": "message represents the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.",
}

func (VolumeError) SwaggerDoc() map[string]string {
	return map_VolumeError
}

// AUTO-GENERATED FUNCTIONS END HERE
