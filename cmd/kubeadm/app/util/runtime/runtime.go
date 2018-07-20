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

package util

import (
	"fmt"
	"path/filepath"
	goruntime "runtime"
	"strings"

	"k8s.io/apimachinery/pkg/util/errors"
	kubeadmapiv1alpha3 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha3"
	utilsexec "k8s.io/utils/exec"
)

const (
	criUnixPrefix = "unix://"
)

// ContainerRuntime is an interface for working with container runtimes
type ContainerRuntime interface {
	IsDocker() bool
	IsRunning() error
	ListKubeContainers() ([]string, error)
	RemoveContainers(containers []string) error
	PullImage(image string) error
	ImageExists(image string) (bool, error)
}

// CRIRuntime is a struct that interfaces with the CRI
type CRIRuntime struct {
	exec      utilsexec.Interface
	criSocket string
}

// NewContainerRuntime sets up and returns a ContainerRuntime struct
func NewContainerRuntime(execer utilsexec.Interface, criSocket string) (ContainerRuntime, error) {
	toolName := "crictl"
	// !!! temporary work around crictl warning:
	// Using "/var/run/crio/crio.sock" as endpoint is deprecated,
	// please consider using full url format "unix:///var/run/crio/crio.sock"
	if filepath.IsAbs(criSocket) && goruntime.GOOS != "windows" {
		criSocket = criUnixPrefix + criSocket
	}

	runtime := &CRIRuntime{execer, criSocket}
	if _, err := execer.LookPath(toolName); err != nil {
		return nil, fmt.Errorf("%s is required for container runtime: %v", toolName, err)
	}

	return runtime, nil
}

// IsDocker returns true if the runtime is docker
func (runtime *CRIRuntime) IsDocker() bool {
	// Use the socket path to determine whether the runtime is Docker.
	// TODO: Use the outoupt of `crictl version` to get the name of the
	// runtime.
	trimmedPath := strings.TrimPrefix(runtime.criSocket, criUnixPrefix)
	return trimmedPath == kubeadmapiv1alpha3.DefaultCRISocket
}

// IsRunning checks if runtime is running
func (runtime *CRIRuntime) IsRunning() error {
	if err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "info").Run(); err != nil {
		return fmt.Errorf("container runtime is not running: %v", err)
	}
	return nil
}

// ListKubeContainers lists running k8s CRI pods
func (runtime *CRIRuntime) ListKubeContainers() ([]string, error) {
	output, err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "pods", "-q").CombinedOutput()
	if err != nil {
		return nil, err
	}
	pods := []string{}
	for _, pod := range strings.Fields(string(output)) {
		if strings.HasPrefix(pod, "k8s_") {
			pods = append(pods, pod)
		}
	}
	return pods, nil
}

// RemoveContainers removes running k8s pods
func (runtime *CRIRuntime) RemoveContainers(containers []string) error {
	errs := []error{}
	for _, container := range containers {
		err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "stopp", container).Run()
		if err != nil {
			// don't stop on errors, try to remove as many containers as possible
			errs = append(errs, fmt.Errorf("failed to stop running pod %s: %v", container, err))
		} else {
			err = runtime.exec.Command("crictl", "-r", runtime.criSocket, "rmp", container).Run()
			if err != nil {
				errs = append(errs, fmt.Errorf("failed to remove running container %s: %v", container, err))
			}
		}
	}
	return errors.NewAggregate(errs)
}

// PullImage pulls the image
func (runtime *CRIRuntime) PullImage(image string) error {
	return runtime.exec.Command("crictl", "-r", runtime.criSocket, "pull", image).Run()
}

// ImageExists checks to see if the image exists on the system
func (runtime *CRIRuntime) ImageExists(image string) (bool, error) {
	err := runtime.exec.Command("crictl", "-r", runtime.criSocket, "inspecti", image).Run()
	return err == nil, err
}
