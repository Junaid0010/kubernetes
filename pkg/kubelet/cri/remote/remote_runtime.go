/*
Copyright 2016 The Kubernetes Authors.

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

package remote

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"

	"k8s.io/component-base/logs/logreduction"
	runtimeapiV1 "k8s.io/cri-api/pkg/apis/runtime/v1"
	runtimeapiV1alpha2 "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	internalapi "k8s.io/kubernetes/pkg/kubelet/apis/cri"
	"k8s.io/kubernetes/pkg/kubelet/cri/remote/util"
	"k8s.io/kubernetes/pkg/probe/exec"
	utilexec "k8s.io/utils/exec"
)

// remoteRuntimeService is a gRPC implementation of internalapi.RuntimeService.
type remoteRuntimeService struct {
	timeout               time.Duration
	runtimeClientV1alpha2 runtimeapiV1alpha2.RuntimeServiceClient
	runtimeClientV1       runtimeapiV1.RuntimeServiceClient
	// Cache last per-container error message to reduce log spam
	logReduction *logreduction.LogReduction
}

const (
	// How frequently to report identical errors
	identicalErrorDelay = 1 * time.Minute
)

// NewRemoteRuntimeService creates a new internalapi.RuntimeService.
func NewRemoteRuntimeService(endpoint string, connectionTimeout time.Duration, apiVersion internalapi.APIVersion) (internalapi.RuntimeService, error) {
	klog.V(3).InfoS("Connecting to runtime service", "endpoint", endpoint)
	addr, dialer, err := util.GetAddressAndDialer(endpoint)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithContextDialer(dialer), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize)))
	if err != nil {
		klog.ErrorS(err, "Connect remote runtime failed", "address", addr)
		return nil, err
	}

	service := &remoteRuntimeService{
		timeout:      connectionTimeout,
		logReduction: logreduction.NewLogReduction(identicalErrorDelay),
	}

	if err := service.determineAPIVersion(conn, apiVersion); err != nil {
		return nil, err
	}

	return service, nil
}

// determineAPIVersion tries to connect to the remote runtime by using the
// provided apiVersion string. If the string is empty, it tries to use v1 over
// v1alpha2.
//
// A GRPC redial will always use the initially selected (or automatically
// determined) CRI API version. If the redial was due to the container runtime
// being upgraded, then the container runtime must also support the initially
// selected version or the redial is expected to fail, which requires a restart
// of kubelet.
func (r *remoteRuntimeService) determineAPIVersion(conn *grpc.ClientConn, apiVersion internalapi.APIVersion) error {
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	switch apiVersion {
	case internalapi.APIVersionV1:
		klog.V(2).InfoS("Using selected CRI v1 API")
		r.runtimeClientV1 = runtimeapiV1.NewRuntimeServiceClient(conn)

	case internalapi.APIVersionV1alpha2:
		klog.V(2).InfoS("Using selected CRI v1alpha2 API")
		r.runtimeClientV1alpha2 = runtimeapiV1alpha2.NewRuntimeServiceClient(conn)

	case "": // Automatically determine the API version while preferring V1
		klog.V(4).InfoS("Finding the CRI API version")
		r.runtimeClientV1 = runtimeapiV1.NewRuntimeServiceClient(conn)

		if _, err := r.runtimeClientV1.Version(ctx, &runtimeapiV1.VersionRequest{}); err == nil {
			klog.V(2).InfoS("Using CRI v1 API")

		} else if status.Code(err) == codes.Unimplemented {
			klog.V(2).InfoS("Falling back to CRI v1alpha2 API (deprecated)")
			r.runtimeClientV1alpha2 = runtimeapiV1alpha2.NewRuntimeServiceClient(conn)

		} else {
			return fmt.Errorf("unable to determine runtime API version: %w", err)
		}

	default:
		return fmt.Errorf(
			"invavlid CRI version, must be empty, %q or %q",
			internalapi.APIVersionV1, internalapi.APIVersionV1alpha2,
		)
	}

	return nil
}

// useV1API returns true if the v1 CRI API should be used instead of v1alpha2.
func (r *remoteRuntimeService) useV1API() bool {
	return r.runtimeClientV1alpha2 == nil
}

func (r *remoteRuntimeService) APIVersion() internalapi.APIVersion {
	if r.useV1API() {
		return internalapi.APIVersionV1
	}
	return internalapi.APIVersionV1alpha2
}

// Version returns the runtime name, runtime version and runtime API version.
func (r *remoteRuntimeService) Version(apiVersion string) (*internalapi.VersionResponse, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] Version", "apiVersion", apiVersion, "timeout", r.timeout)

	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.versionV1(ctx, apiVersion)
	}

	return r.versionV1alpha2(ctx, apiVersion)
}

func (r *remoteRuntimeService) versionV1(ctx context.Context, apiVersion string) (*internalapi.VersionResponse, error) {
	typedVersion, err := r.runtimeClientV1.Version(ctx, &runtimeapiV1.VersionRequest{
		Version: apiVersion,
	})
	if err != nil {
		klog.ErrorS(err, "Version from runtime service failed")
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] Version Response", "apiVersion", typedVersion)

	if typedVersion.Version == "" || typedVersion.RuntimeName == "" || typedVersion.RuntimeApiVersion == "" || typedVersion.RuntimeVersion == "" {
		return nil, fmt.Errorf("not all fields are set in VersionResponse (%q)", *typedVersion)
	}

	return internalapi.FromV1VersionResponse(typedVersion), err
}

func (r *remoteRuntimeService) versionV1alpha2(ctx context.Context, apiVersion string) (*internalapi.VersionResponse, error) {
	typedVersion, err := r.runtimeClientV1alpha2.Version(ctx, &runtimeapiV1alpha2.VersionRequest{
		Version: apiVersion,
	})
	if err != nil {
		klog.ErrorS(err, "Version from runtime service failed")
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] Version Response", "apiVersion", typedVersion)

	if typedVersion.Version == "" || typedVersion.RuntimeName == "" || typedVersion.RuntimeApiVersion == "" || typedVersion.RuntimeVersion == "" {
		return nil, fmt.Errorf("not all fields are set in VersionResponse (%q)", *typedVersion)
	}

	return internalapi.FromV1alpha2VersionResponse(typedVersion), err
}

// RunPodSandbox creates and starts a pod-level sandbox. Runtimes should ensure
// the sandbox is in ready state.
func (r *remoteRuntimeService) RunPodSandbox(config *internalapi.PodSandboxConfig, runtimeHandler string) (string, error) {
	// Use 2 times longer timeout for sandbox operation (4 mins by default)
	// TODO: Make the pod sandbox timeout configurable.
	timeout := r.timeout * 2

	klog.V(10).InfoS("[RemoteRuntimeService] RunPodSandbox", "config", config, "runtimeHandler", runtimeHandler, "timeout", timeout)

	ctx, cancel := getContextWithTimeout(timeout)
	defer cancel()

	var podSandboxID string
	if r.useV1API() {
		resp, err := r.runtimeClientV1.RunPodSandbox(ctx, &runtimeapiV1.RunPodSandboxRequest{
			Config:         internalapi.V1PodSandboxConfig(config),
			RuntimeHandler: runtimeHandler,
		})

		if err != nil {
			klog.ErrorS(err, "RunPodSandbox from runtime service failed")
			return "", err
		}
		podSandboxID = resp.PodSandboxId
	} else {
		resp, err := r.runtimeClientV1alpha2.RunPodSandbox(ctx, &runtimeapiV1alpha2.RunPodSandboxRequest{
			Config:         internalapi.V1alpha2PodSandboxConfig(config),
			RuntimeHandler: runtimeHandler,
		})

		if err != nil {
			klog.ErrorS(err, "RunPodSandbox from runtime service failed")
			return "", err
		}
		podSandboxID = resp.PodSandboxId
	}

	if podSandboxID == "" {
		errorMessage := fmt.Sprintf("PodSandboxId is not set for sandbox %q", config.Metadata)
		err := errors.New(errorMessage)
		klog.ErrorS(err, "RunPodSandbox failed")
		return "", err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] RunPodSandbox Response", "podSandboxID", podSandboxID)

	return podSandboxID, nil
}

// StopPodSandbox stops the sandbox. If there are any running containers in the
// sandbox, they should be forced to termination.
func (r *remoteRuntimeService) StopPodSandbox(podSandBoxID string) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] StopPodSandbox", "podSandboxID", podSandBoxID, "timeout", r.timeout)

	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		_, err = r.runtimeClientV1.StopPodSandbox(ctx, &runtimeapiV1.StopPodSandboxRequest{
			PodSandboxId: podSandBoxID,
		})
	} else {
		_, err = r.runtimeClientV1alpha2.StopPodSandbox(ctx, &runtimeapiV1alpha2.StopPodSandboxRequest{
			PodSandboxId: podSandBoxID,
		})
	}
	if err != nil {
		klog.ErrorS(err, "StopPodSandbox from runtime service failed", "podSandboxID", podSandBoxID)
		return err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] StopPodSandbox Response", "podSandboxID", podSandBoxID)

	return nil
}

// RemovePodSandbox removes the sandbox. If there are any containers in the
// sandbox, they should be forcibly removed.
func (r *remoteRuntimeService) RemovePodSandbox(podSandBoxID string) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] RemovePodSandbox", "podSandboxID", podSandBoxID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		_, err = r.runtimeClientV1.RemovePodSandbox(ctx, &runtimeapiV1.RemovePodSandboxRequest{
			PodSandboxId: podSandBoxID,
		})
	} else {
		_, err = r.runtimeClientV1alpha2.RemovePodSandbox(ctx, &runtimeapiV1alpha2.RemovePodSandboxRequest{
			PodSandboxId: podSandBoxID,
		})
	}
	if err != nil {
		klog.ErrorS(err, "RemovePodSandbox from runtime service failed", "podSandboxID", podSandBoxID)
		return err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] RemovePodSandbox Response", "podSandboxID", podSandBoxID)

	return nil
}

// PodSandboxStatus returns the status of the PodSandbox.
func (r *remoteRuntimeService) PodSandboxStatus(podSandBoxID string) (*internalapi.PodSandboxStatus, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] PodSandboxStatus", "podSandboxID", podSandBoxID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.podSandboxStatusV1(ctx, podSandBoxID)
	}

	return r.podSandboxStatusV1alpha2(ctx, podSandBoxID)
}

func (r *remoteRuntimeService) podSandboxStatusV1alpha2(ctx context.Context, podSandBoxID string) (*internalapi.PodSandboxStatus, error) {
	resp, err := r.runtimeClientV1alpha2.PodSandboxStatus(ctx, &runtimeapiV1alpha2.PodSandboxStatusRequest{
		PodSandboxId: podSandBoxID,
	})
	if err != nil {
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] PodSandboxStatus Response", "podSandboxID", podSandBoxID, "status", resp.Status)

	status := internalapi.FromV1alpha2PodSandboxStatus(resp.Status)
	if resp.Status != nil {
		if err := verifySandboxStatus(status); err != nil {
			return nil, err
		}
	}

	return status, nil
}

func (r *remoteRuntimeService) podSandboxStatusV1(ctx context.Context, podSandBoxID string) (*internalapi.PodSandboxStatus, error) {
	resp, err := r.runtimeClientV1.PodSandboxStatus(ctx, &runtimeapiV1.PodSandboxStatusRequest{
		PodSandboxId: podSandBoxID,
	})
	if err != nil {
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] PodSandboxStatus Response", "podSandboxID", podSandBoxID, "status", resp.Status)

	status := internalapi.FromV1PodSandboxStatus(resp.Status)
	if resp.Status != nil {
		if err := verifySandboxStatus(status); err != nil {
			return nil, err
		}
	}

	return status, nil
}

// ListPodSandbox returns a list of PodSandboxes.
func (r *remoteRuntimeService) ListPodSandbox(filter *internalapi.PodSandboxFilter) ([]*internalapi.PodSandbox, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ListPodSandbox", "filter", filter, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.listPodSandboxV1(ctx, filter)
	}

	return r.listPodSandboxV1alpha2(ctx, filter)
}

func (r *remoteRuntimeService) listPodSandboxV1alpha2(ctx context.Context, filter *internalapi.PodSandboxFilter) ([]*internalapi.PodSandbox, error) {
	resp, err := r.runtimeClientV1alpha2.ListPodSandbox(ctx, &runtimeapiV1alpha2.ListPodSandboxRequest{
		Filter: internalapi.V1alpha2PodSandboxFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListPodSandbox with filter from runtime service failed", "filter", filter)
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] ListPodSandbox Response", "filter", filter, "items", resp.Items)

	return internalapi.FromV1alpha2PodSandboxes(resp.Items), nil
}

func (r *remoteRuntimeService) listPodSandboxV1(ctx context.Context, filter *internalapi.PodSandboxFilter) ([]*internalapi.PodSandbox, error) {
	resp, err := r.runtimeClientV1.ListPodSandbox(ctx, &runtimeapiV1.ListPodSandboxRequest{
		Filter: internalapi.V1PodSandboxFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListPodSandbox with filter from runtime service failed", "filter", filter)
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] ListPodSandbox Response", "filter", filter, "items", resp.Items)

	return internalapi.FromV1PodSandboxes(resp.Items), nil
}

// CreateContainer creates a new container in the specified PodSandbox.
func (r *remoteRuntimeService) CreateContainer(podSandBoxID string, config *internalapi.ContainerConfig, sandboxConfig *internalapi.PodSandboxConfig) (string, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] CreateContainer", "podSandboxID", podSandBoxID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.createContainerV1(ctx, podSandBoxID, config, sandboxConfig)
	}

	return r.createContainerV1alpha2(ctx, podSandBoxID, config, sandboxConfig)
}

func (r *remoteRuntimeService) createContainerV1alpha2(ctx context.Context, podSandBoxID string, config *internalapi.ContainerConfig, sandboxConfig *internalapi.PodSandboxConfig) (string, error) {
	resp, err := r.runtimeClientV1alpha2.CreateContainer(ctx, &runtimeapiV1alpha2.CreateContainerRequest{
		PodSandboxId:  podSandBoxID,
		Config:        internalapi.V1alpha2ContainerConfig(config),
		SandboxConfig: internalapi.V1alpha2PodSandboxConfig(sandboxConfig),
	})
	if err != nil {
		klog.ErrorS(err, "CreateContainer in sandbox from runtime service failed", "podSandboxID", podSandBoxID)
		return "", err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] CreateContainer", "podSandboxID", podSandBoxID, "containerID", resp.ContainerId)
	if resp.ContainerId == "" {
		errorMessage := fmt.Sprintf("ContainerId is not set for container %q", config.Metadata)
		err := errors.New(errorMessage)
		klog.ErrorS(err, "CreateContainer failed")
		return "", err
	}

	return resp.ContainerId, nil
}

func (r *remoteRuntimeService) createContainerV1(ctx context.Context, podSandBoxID string, config *internalapi.ContainerConfig, sandboxConfig *internalapi.PodSandboxConfig) (string, error) {
	resp, err := r.runtimeClientV1.CreateContainer(ctx, &runtimeapiV1.CreateContainerRequest{
		PodSandboxId:  podSandBoxID,
		Config:        internalapi.V1ContainerConfig(config),
		SandboxConfig: internalapi.V1PodSandboxConfig(sandboxConfig),
	})
	if err != nil {
		klog.ErrorS(err, "CreateContainer in sandbox from runtime service failed", "podSandboxID", podSandBoxID)
		return "", err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] CreateContainer", "podSandboxID", podSandBoxID, "containerID", resp.ContainerId)
	if resp.ContainerId == "" {
		errorMessage := fmt.Sprintf("ContainerId is not set for container %q", config.Metadata)
		err := errors.New(errorMessage)
		klog.ErrorS(err, "CreateContainer failed")
		return "", err
	}

	return resp.ContainerId, nil
}

// StartContainer starts the container.
func (r *remoteRuntimeService) StartContainer(containerID string) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] StartContainer", "containerID", containerID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		_, err = r.runtimeClientV1.StartContainer(ctx, &runtimeapiV1.StartContainerRequest{
			ContainerId: containerID,
		})
	} else {
		_, err = r.runtimeClientV1alpha2.StartContainer(ctx, &runtimeapiV1alpha2.StartContainerRequest{
			ContainerId: containerID,
		})
	}

	if err != nil {
		klog.ErrorS(err, "StartContainer from runtime service failed", "containerID", containerID)
		return err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] StartContainer Response", "containerID", containerID)

	return nil
}

// StopContainer stops a running container with a grace period (i.e., timeout).
func (r *remoteRuntimeService) StopContainer(containerID string, timeout int64) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] StopContainer", "containerID", containerID, "timeout", timeout)
	// Use timeout + default timeout (2 minutes) as timeout to leave extra time
	// for SIGKILL container and request latency.
	t := r.timeout + time.Duration(timeout)*time.Second
	ctx, cancel := getContextWithTimeout(t)
	defer cancel()

	r.logReduction.ClearID(containerID)

	if r.useV1API() {
		_, err = r.runtimeClientV1.StopContainer(ctx, &runtimeapiV1.StopContainerRequest{
			ContainerId: containerID,
			Timeout:     timeout,
		})
	} else {
		_, err = r.runtimeClientV1alpha2.StopContainer(ctx, &runtimeapiV1alpha2.StopContainerRequest{
			ContainerId: containerID,
			Timeout:     timeout,
		})
	}
	if err != nil {
		klog.ErrorS(err, "StopContainer from runtime service failed", "containerID", containerID)
		return err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] StopContainer Response", "containerID", containerID)

	return nil
}

// RemoveContainer removes the container. If the container is running, the container
// should be forced to removal.
func (r *remoteRuntimeService) RemoveContainer(containerID string) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] RemoveContainer", "containerID", containerID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	r.logReduction.ClearID(containerID)
	if r.useV1API() {
		_, err = r.runtimeClientV1.RemoveContainer(ctx, &runtimeapiV1.RemoveContainerRequest{
			ContainerId: containerID,
		})
	} else {
		_, err = r.runtimeClientV1alpha2.RemoveContainer(ctx, &runtimeapiV1alpha2.RemoveContainerRequest{
			ContainerId: containerID,
		})
	}
	if err != nil {
		klog.ErrorS(err, "RemoveContainer from runtime service failed", "containerID", containerID)
		return err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] RemoveContainer Response", "containerID", containerID)

	return nil
}

// ListContainers lists containers by filters.
func (r *remoteRuntimeService) ListContainers(filter *internalapi.ContainerFilter) ([]*internalapi.Container, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ListContainers", "filter", filter, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.listContainersV1(ctx, filter)
	}

	return r.listContainersV1alpha2(ctx, filter)
}

func (r *remoteRuntimeService) listContainersV1alpha2(ctx context.Context, filter *internalapi.ContainerFilter) ([]*internalapi.Container, error) {
	resp, err := r.runtimeClientV1alpha2.ListContainers(ctx, &runtimeapiV1alpha2.ListContainersRequest{
		Filter: internalapi.V1alpha2ContainerFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListContainers with filter from runtime service failed", "filter", filter)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] ListContainers Response", "filter", filter, "containers", resp.Containers)

	return internalapi.FromV1alpha2Containers(resp.Containers), nil
}

func (r *remoteRuntimeService) listContainersV1(ctx context.Context, filter *internalapi.ContainerFilter) ([]*internalapi.Container, error) {
	resp, err := r.runtimeClientV1.ListContainers(ctx, &runtimeapiV1.ListContainersRequest{
		Filter: internalapi.V1ContainerFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListContainers with filter from runtime service failed", "filter", filter)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] ListContainers Response", "filter", filter, "containers", resp.Containers)

	return internalapi.FromV1Containers(resp.Containers), nil
}

// ContainerStatus returns the container status.
func (r *remoteRuntimeService) ContainerStatus(containerID string) (*internalapi.ContainerStatus, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ContainerStatus", "containerID", containerID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.containerStatusV1(ctx, containerID)
	}

	return r.containerStatusV1alpha2(ctx, containerID)
}

func (r *remoteRuntimeService) containerStatusV1alpha2(ctx context.Context, containerID string) (*internalapi.ContainerStatus, error) {
	resp, err := r.runtimeClientV1alpha2.ContainerStatus(ctx, &runtimeapiV1alpha2.ContainerStatusRequest{
		ContainerId: containerID,
	})
	if err != nil {
		// Don't spam the log with endless messages about the same failure.
		if r.logReduction.ShouldMessageBePrinted(err.Error(), containerID) {
			klog.ErrorS(err, "ContainerStatus from runtime service failed", "containerID", containerID)
		}
		return nil, err
	}
	r.logReduction.ClearID(containerID)
	klog.V(10).InfoS("[RemoteRuntimeService] ContainerStatus Response", "containerID", containerID, "status", resp.Status)

	status := internalapi.FromV1alpha2ContainerStatus(resp.Status)
	if resp.Status != nil {
		if err := verifyContainerStatus(status); err != nil {
			klog.ErrorS(err, "verify ContainerStatus failed", "containerID", containerID)
			return nil, err
		}
	}

	return status, nil
}

func (r *remoteRuntimeService) containerStatusV1(ctx context.Context, containerID string) (*internalapi.ContainerStatus, error) {
	resp, err := r.runtimeClientV1.ContainerStatus(ctx, &runtimeapiV1.ContainerStatusRequest{
		ContainerId: containerID,
	})
	if err != nil {
		// Don't spam the log with endless messages about the same failure.
		if r.logReduction.ShouldMessageBePrinted(err.Error(), containerID) {
			klog.ErrorS(err, "ContainerStatus from runtime service failed", "containerID", containerID)
		}
		return nil, err
	}
	r.logReduction.ClearID(containerID)
	klog.V(10).InfoS("[RemoteRuntimeService] ContainerStatus Response", "containerID", containerID, "status", resp.Status)

	status := internalapi.FromV1ContainerStatus(resp.Status)
	if resp.Status != nil {
		if err := verifyContainerStatus(status); err != nil {
			klog.ErrorS(err, "verify ContainerStatus failed", "containerID", containerID)
			return nil, err
		}
	}

	return status, nil
}

// UpdateContainerResources updates a containers resource config
func (r *remoteRuntimeService) UpdateContainerResources(containerID string, resources *internalapi.LinuxContainerResources) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] UpdateContainerResources", "containerID", containerID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		_, err = r.runtimeClientV1.UpdateContainerResources(ctx, &runtimeapiV1.UpdateContainerResourcesRequest{
			ContainerId: containerID,
			Linux:       internalapi.V1ContainerResources(resources),
		})
	} else {
		_, err = r.runtimeClientV1alpha2.UpdateContainerResources(ctx, &runtimeapiV1alpha2.UpdateContainerResourcesRequest{
			ContainerId: containerID,
			Linux:       internalapi.V1alpha2ContainerResources(resources),
		})
	}
	if err != nil {
		klog.ErrorS(err, "UpdateContainerResources from runtime service failed", "containerID", containerID)
		return err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] UpdateContainerResources Response", "containerID", containerID)

	return nil
}

// ExecSync executes a command in the container, and returns the stdout output.
// If command exits with a non-zero exit code, an error is returned.
func (r *remoteRuntimeService) ExecSync(containerID string, cmd []string, timeout time.Duration) (stdout []byte, stderr []byte, err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ExecSync", "containerID", containerID, "timeout", timeout)
	// Do not set timeout when timeout is 0.
	var ctx context.Context
	var cancel context.CancelFunc
	if timeout != 0 {
		// Use timeout + default timeout (2 minutes) as timeout to leave some time for
		// the runtime to do cleanup.
		ctx, cancel = getContextWithTimeout(r.timeout + timeout)
	} else {
		ctx, cancel = getContextWithCancel()
	}
	defer cancel()

	if r.useV1API() {
		return r.execSyncV1(ctx, containerID, cmd, timeout)
	}

	return r.execSyncV1alpha2(ctx, containerID, cmd, timeout)
}

func (r *remoteRuntimeService) execSyncV1alpha2(ctx context.Context, containerID string, cmd []string, timeout time.Duration) (stdout []byte, stderr []byte, err error) {
	timeoutSeconds := int64(timeout.Seconds())
	req := &runtimeapiV1alpha2.ExecSyncRequest{
		ContainerId: containerID,
		Cmd:         cmd,
		Timeout:     timeoutSeconds,
	}
	resp, err := r.runtimeClientV1alpha2.ExecSync(ctx, req)
	if err != nil {
		klog.ErrorS(err, "ExecSync cmd from runtime service failed", "containerID", containerID, "cmd", cmd)

		// interpret DeadlineExceeded gRPC errors as timedout probes
		if status.Code(err) == codes.DeadlineExceeded {
			err = exec.NewTimeoutError(fmt.Errorf("command %q timed out", strings.Join(cmd, " ")), timeout)
		}

		return nil, nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] ExecSync Response", "containerID", containerID, "exitCode", resp.ExitCode)
	err = nil
	if resp.ExitCode != 0 {
		err = utilexec.CodeExitError{
			Err:  fmt.Errorf("command '%s' exited with %d: %s", strings.Join(cmd, " "), resp.ExitCode, resp.Stderr),
			Code: int(resp.ExitCode),
		}
	}

	return resp.Stdout, resp.Stderr, err
}

func (r *remoteRuntimeService) execSyncV1(ctx context.Context, containerID string, cmd []string, timeout time.Duration) (stdout []byte, stderr []byte, err error) {
	timeoutSeconds := int64(timeout.Seconds())
	req := &runtimeapiV1.ExecSyncRequest{
		ContainerId: containerID,
		Cmd:         cmd,
		Timeout:     timeoutSeconds,
	}
	resp, err := r.runtimeClientV1.ExecSync(ctx, req)
	if err != nil {
		klog.ErrorS(err, "ExecSync cmd from runtime service failed", "containerID", containerID, "cmd", cmd)

		// interpret DeadlineExceeded gRPC errors as timedout probes
		if status.Code(err) == codes.DeadlineExceeded {
			err = exec.NewTimeoutError(fmt.Errorf("command %q timed out", strings.Join(cmd, " ")), timeout)
		}

		return nil, nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] ExecSync Response", "containerID", containerID, "exitCode", resp.ExitCode)
	err = nil
	if resp.ExitCode != 0 {
		err = utilexec.CodeExitError{
			Err:  fmt.Errorf("command '%s' exited with %d: %s", strings.Join(cmd, " "), resp.ExitCode, resp.Stderr),
			Code: int(resp.ExitCode),
		}
	}

	return resp.Stdout, resp.Stderr, err
}

// Exec prepares a streaming endpoint to execute a command in the container, and returns the address.
func (r *remoteRuntimeService) Exec(req *internalapi.ExecRequest) (*internalapi.ExecResponse, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] Exec", "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.execV1(ctx, req)
	}

	return r.execV1alpha2(ctx, req)
}

func (r *remoteRuntimeService) execV1alpha2(ctx context.Context, req *internalapi.ExecRequest) (*internalapi.ExecResponse, error) {
	resp, err := r.runtimeClientV1alpha2.Exec(ctx, internalapi.V1alpha2ExecRequest(req))
	if err != nil {
		klog.ErrorS(err, "Exec cmd from runtime service failed", "containerID", req.ContainerId, "cmd", req.Cmd)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] Exec Response")

	if resp.Url == "" {
		errorMessage := "URL is not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "Exec failed")
		return nil, err
	}

	return internalapi.FromV1alpha2ExecResponse(resp), nil
}

func (r *remoteRuntimeService) execV1(ctx context.Context, req *internalapi.ExecRequest) (*internalapi.ExecResponse, error) {
	resp, err := r.runtimeClientV1.Exec(ctx, internalapi.V1ExecRequest(req))
	if err != nil {
		klog.ErrorS(err, "Exec cmd from runtime service failed", "containerID", req.ContainerId, "cmd", req.Cmd)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] Exec Response")

	if resp.Url == "" {
		errorMessage := "URL is not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "Exec failed")
		return nil, err
	}

	return internalapi.FromV1ExecResponse(resp), nil
}

// Attach prepares a streaming endpoint to attach to a running container, and returns the address.
func (r *remoteRuntimeService) Attach(req *internalapi.AttachRequest) (*internalapi.AttachResponse, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] Attach", "containerID", req.ContainerId, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.attachV1(ctx, req)
	}

	return r.attachV1alpha2(ctx, req)
}

func (r *remoteRuntimeService) attachV1alpha2(ctx context.Context, req *internalapi.AttachRequest) (*internalapi.AttachResponse, error) {
	resp, err := r.runtimeClientV1alpha2.Attach(ctx, internalapi.V1alpha2AttachRequest(req))
	if err != nil {
		klog.ErrorS(err, "Attach container from runtime service failed", "containerID", req.ContainerId)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] Attach Response", "containerID", req.ContainerId)

	if resp.Url == "" {
		errorMessage := "URL is not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "Attach failed")
		return nil, err
	}
	return internalapi.FromV1alpha2AttachResponse(resp), nil
}

func (r *remoteRuntimeService) attachV1(ctx context.Context, req *internalapi.AttachRequest) (*internalapi.AttachResponse, error) {
	resp, err := r.runtimeClientV1.Attach(ctx, internalapi.V1AttachRequest(req))
	if err != nil {
		klog.ErrorS(err, "Attach container from runtime service failed", "containerID", req.ContainerId)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] Attach Response", "containerID", req.ContainerId)

	if resp.Url == "" {
		errorMessage := "URL is not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "Attach failed")
		return nil, err
	}
	return internalapi.FromV1AttachResponse(resp), nil
}

// PortForward prepares a streaming endpoint to forward ports from a PodSandbox, and returns the address.
func (r *remoteRuntimeService) PortForward(req *internalapi.PortForwardRequest) (*internalapi.PortForwardResponse, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] PortForward", "podSandboxID", req.PodSandboxId, "port", req.Port, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.portForwardV1(ctx, req)
	}

	return r.portForwardV1alpha2(ctx, req)
}

func (r *remoteRuntimeService) portForwardV1alpha2(ctx context.Context, req *internalapi.PortForwardRequest) (*internalapi.PortForwardResponse, error) {
	resp, err := r.runtimeClientV1alpha2.PortForward(ctx, internalapi.V1alpha2PortForwardRequest(req))
	if err != nil {
		klog.ErrorS(err, "PortForward from runtime service failed", "podSandboxID", req.PodSandboxId)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] PortForward Response", "podSandboxID", req.PodSandboxId)

	if resp.Url == "" {
		errorMessage := "URL is not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "PortForward failed")
		return nil, err
	}

	return internalapi.FromV1alpha2PortForwardResponse(resp), nil
}

func (r *remoteRuntimeService) portForwardV1(ctx context.Context, req *internalapi.PortForwardRequest) (*internalapi.PortForwardResponse, error) {
	resp, err := r.runtimeClientV1.PortForward(ctx, internalapi.V1PortForwardRequest(req))
	if err != nil {
		klog.ErrorS(err, "PortForward from runtime service failed", "podSandboxID", req.PodSandboxId)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] PortForward Response", "podSandboxID", req.PodSandboxId)

	if resp.Url == "" {
		errorMessage := "URL is not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "PortForward failed")
		return nil, err
	}

	return internalapi.FromV1PortForwardResponse(resp), nil
}

// UpdateRuntimeConfig updates the config of a runtime service. The only
// update payload currently supported is the pod CIDR assigned to a node,
// and the runtime service just proxies it down to the network plugin.
func (r *remoteRuntimeService) UpdateRuntimeConfig(runtimeConfig *internalapi.RuntimeConfig) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] UpdateRuntimeConfig", "runtimeConfig", runtimeConfig, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	// Response doesn't contain anything of interest. This translates to an
	// Event notification to the network plugin, which can't fail, so we're
	// really looking to surface destination unreachable.
	if r.useV1API() {
		_, err = r.runtimeClientV1.UpdateRuntimeConfig(ctx, &runtimeapiV1.UpdateRuntimeConfigRequest{
			RuntimeConfig: internalapi.V1RuntimeConfig(runtimeConfig),
		})
	} else {
		_, err = r.runtimeClientV1alpha2.UpdateRuntimeConfig(ctx, &runtimeapiV1alpha2.UpdateRuntimeConfigRequest{
			RuntimeConfig: internalapi.V1alpha2RuntimeConfig(runtimeConfig),
		})
	}

	if err != nil {
		return err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] UpdateRuntimeConfig Response", "runtimeConfig", runtimeConfig)

	return nil
}

// Status returns the status of the runtime.
func (r *remoteRuntimeService) Status() (*internalapi.RuntimeStatus, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] Status", "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.statusV1(ctx)
	}

	return r.statusV1alpha2(ctx)
}

func (r *remoteRuntimeService) statusV1alpha2(ctx context.Context) (*internalapi.RuntimeStatus, error) {
	resp, err := r.runtimeClientV1alpha2.Status(ctx, &runtimeapiV1alpha2.StatusRequest{})
	if err != nil {
		klog.ErrorS(err, "Status from runtime service failed")
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] Status Response", "status", resp.Status)

	if resp.Status == nil || len(resp.Status.Conditions) < 2 {
		errorMessage := "RuntimeReady or NetworkReady condition are not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "Status failed")
		return nil, err
	}

	return internalapi.FromV1alpha2RuntimeStatus(resp.Status), nil
}

func (r *remoteRuntimeService) statusV1(ctx context.Context) (*internalapi.RuntimeStatus, error) {
	resp, err := r.runtimeClientV1.Status(ctx, &runtimeapiV1.StatusRequest{})
	if err != nil {
		klog.ErrorS(err, "Status from runtime service failed")
		return nil, err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] Status Response", "status", resp.Status)

	if resp.Status == nil || len(resp.Status.Conditions) < 2 {
		errorMessage := "RuntimeReady or NetworkReady condition are not set"
		err := errors.New(errorMessage)
		klog.ErrorS(err, "Status failed")
		return nil, err
	}

	return internalapi.FromV1RuntimeStatus(resp.Status), nil
}

// ContainerStats returns the stats of the container.
func (r *remoteRuntimeService) ContainerStats(containerID string) (*internalapi.ContainerStats, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ContainerStats", "containerID", containerID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.containerStatsV1(ctx, containerID)
	}

	return r.containerStatsV1alpha2(ctx, containerID)
}

func (r *remoteRuntimeService) containerStatsV1alpha2(ctx context.Context, containerID string) (*internalapi.ContainerStats, error) {
	resp, err := r.runtimeClientV1alpha2.ContainerStats(ctx, &runtimeapiV1alpha2.ContainerStatsRequest{
		ContainerId: containerID,
	})
	if err != nil {
		if r.logReduction.ShouldMessageBePrinted(err.Error(), containerID) {
			klog.ErrorS(err, "ContainerStats from runtime service failed", "containerID", containerID)
		}
		return nil, err
	}
	r.logReduction.ClearID(containerID)
	klog.V(10).InfoS("[RemoteRuntimeService] ContainerStats Response", "containerID", containerID, "stats", resp.GetStats())

	return internalapi.FromV1alpha2ContainerStats(resp.GetStats()), nil
}

func (r *remoteRuntimeService) containerStatsV1(ctx context.Context, containerID string) (*internalapi.ContainerStats, error) {
	resp, err := r.runtimeClientV1.ContainerStats(ctx, &runtimeapiV1.ContainerStatsRequest{
		ContainerId: containerID,
	})
	if err != nil {
		if r.logReduction.ShouldMessageBePrinted(err.Error(), containerID) {
			klog.ErrorS(err, "ContainerStats from runtime service failed", "containerID", containerID)
		}
		return nil, err
	}
	r.logReduction.ClearID(containerID)
	klog.V(10).InfoS("[RemoteRuntimeService] ContainerStats Response", "containerID", containerID, "stats", resp.GetStats())

	return internalapi.FromV1ContainerStats(resp.GetStats()), nil
}

// ListContainerStats returns the list of ContainerStats given the filter.
func (r *remoteRuntimeService) ListContainerStats(filter *internalapi.ContainerStatsFilter) ([]*internalapi.ContainerStats, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ListContainerStats", "filter", filter)
	// Do not set timeout, because writable layer stats collection takes time.
	// TODO(random-liu): Should we assume runtime should cache the result, and set timeout here?
	ctx, cancel := getContextWithCancel()
	defer cancel()

	if r.useV1API() {
		return r.listContainerStatsV1(ctx, filter)
	}

	return r.listContainerStatsV1alpha2(ctx, filter)
}

func (r *remoteRuntimeService) listContainerStatsV1alpha2(ctx context.Context, filter *internalapi.ContainerStatsFilter) ([]*internalapi.ContainerStats, error) {
	resp, err := r.runtimeClientV1alpha2.ListContainerStats(ctx, &runtimeapiV1alpha2.ListContainerStatsRequest{
		Filter: internalapi.V1alpha2ContainerStatsFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListContainerStats with filter from runtime service failed", "filter", filter)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] ListContainerStats Response", "filter", filter, "stats", resp.GetStats())

	return internalapi.FromV1alpha2ContainerStatsList(resp.GetStats()), nil
}

func (r *remoteRuntimeService) listContainerStatsV1(ctx context.Context, filter *internalapi.ContainerStatsFilter) ([]*internalapi.ContainerStats, error) {
	resp, err := r.runtimeClientV1.ListContainerStats(ctx, &runtimeapiV1.ListContainerStatsRequest{
		Filter: internalapi.V1ContainerStatsFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListContainerStats with filter from runtime service failed", "filter", filter)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] ListContainerStats Response", "filter", filter, "stats", resp.GetStats())

	return internalapi.FromV1ContainerStatsList(resp.GetStats()), nil
}

// PodSandboxStats returns the stats of the pod.
func (r *remoteRuntimeService) PodSandboxStats(podSandboxID string) (*internalapi.PodSandboxStats, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] PodSandboxStats", "podSandboxID", podSandboxID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.podSandboxStatsV1(ctx, podSandboxID)
	}

	return r.podSandboxStatsV1alpha2(ctx, podSandboxID)
}

func (r *remoteRuntimeService) podSandboxStatsV1alpha2(ctx context.Context, podSandboxID string) (*internalapi.PodSandboxStats, error) {
	resp, err := r.runtimeClientV1alpha2.PodSandboxStats(ctx, &runtimeapiV1alpha2.PodSandboxStatsRequest{
		PodSandboxId: podSandboxID,
	})
	if err != nil {
		if r.logReduction.ShouldMessageBePrinted(err.Error(), podSandboxID) {
			klog.ErrorS(err, "PodSandbox from runtime service failed", "podSandboxID", podSandboxID)
		}
		return nil, err
	}
	r.logReduction.ClearID(podSandboxID)
	klog.V(10).InfoS("[RemoteRuntimeService] PodSandbox Response", "podSandboxID", podSandboxID, "stats", resp.GetStats())

	return internalapi.FromV1alpha2PodSandboxStats(resp.GetStats()), nil
}

func (r *remoteRuntimeService) podSandboxStatsV1(ctx context.Context, podSandboxID string) (*internalapi.PodSandboxStats, error) {
	resp, err := r.runtimeClientV1.PodSandboxStats(ctx, &runtimeapiV1.PodSandboxStatsRequest{
		PodSandboxId: podSandboxID,
	})
	if err != nil {
		if r.logReduction.ShouldMessageBePrinted(err.Error(), podSandboxID) {
			klog.ErrorS(err, "PodSandbox from runtime service failed", "podSandboxID", podSandboxID)
		}
		return nil, err
	}
	r.logReduction.ClearID(podSandboxID)
	klog.V(10).InfoS("[RemoteRuntimeService] PodSandbox Response", "podSandboxID", podSandboxID, "stats", resp.GetStats())

	return internalapi.FromV1PodSandboxStats(resp.GetStats()), nil
}

// ListPodSandboxStats returns the list of pod sandbox stats given the filter
func (r *remoteRuntimeService) ListPodSandboxStats(filter *internalapi.PodSandboxStatsFilter) ([]*internalapi.PodSandboxStats, error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ListPodSandboxStats", "filter", filter)
	// Set timeout, because runtimes are able to cache disk stats results
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		return r.listPodSandboxStatsV1(ctx, filter)
	}

	return r.listPodSandboxStatsV1alpha2(ctx, filter)
}

func (r *remoteRuntimeService) listPodSandboxStatsV1alpha2(ctx context.Context, filter *internalapi.PodSandboxStatsFilter) ([]*internalapi.PodSandboxStats, error) {
	resp, err := r.runtimeClientV1alpha2.ListPodSandboxStats(ctx, &runtimeapiV1alpha2.ListPodSandboxStatsRequest{
		Filter: internalapi.V1alpha2PodSandboxStatsFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListPodSandboxStats with filter from runtime service failed", "filter", filter)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] ListPodSandboxStats Response", "filter", filter, "stats", resp.GetStats())

	return internalapi.FromV1alpha2PodSandboxStatsList(resp.GetStats()), nil
}

func (r *remoteRuntimeService) listPodSandboxStatsV1(ctx context.Context, filter *internalapi.PodSandboxStatsFilter) ([]*internalapi.PodSandboxStats, error) {
	resp, err := r.runtimeClientV1.ListPodSandboxStats(ctx, &runtimeapiV1.ListPodSandboxStatsRequest{
		Filter: internalapi.V1PodSandboxStatsFilter(filter),
	})
	if err != nil {
		klog.ErrorS(err, "ListPodSandboxStats with filter from runtime service failed", "filter", filter)
		return nil, err
	}
	klog.V(10).InfoS("[RemoteRuntimeService] ListPodSandboxStats Response", "filter", filter, "stats", resp.GetStats())

	return internalapi.FromV1PodSandboxStatsList(resp.GetStats()), nil
}

// ReopenContainerLog reopens the container log file.
func (r *remoteRuntimeService) ReopenContainerLog(containerID string) (err error) {
	klog.V(10).InfoS("[RemoteRuntimeService] ReopenContainerLog", "containerID", containerID, "timeout", r.timeout)
	ctx, cancel := getContextWithTimeout(r.timeout)
	defer cancel()

	if r.useV1API() {
		_, err = r.runtimeClientV1.ReopenContainerLog(ctx, &runtimeapiV1.ReopenContainerLogRequest{ContainerId: containerID})
	} else {
		_, err = r.runtimeClientV1alpha2.ReopenContainerLog(ctx, &runtimeapiV1alpha2.ReopenContainerLogRequest{ContainerId: containerID})
	}
	if err != nil {
		klog.ErrorS(err, "ReopenContainerLog from runtime service failed", "containerID", containerID)
		return err
	}

	klog.V(10).InfoS("[RemoteRuntimeService] ReopenContainerLog Response", "containerID", containerID)
	return nil
}
