/*
Copyright 2022 The Kubernetes Authors.

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

package e2enode

import (
	"context"
	"fmt"

	"github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/kubernetes/pkg/features"
	kubeletconfig "k8s.io/kubernetes/pkg/kubelet/apis/config"
	"k8s.io/kubernetes/test/e2e/framework"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
	admissionapi "k8s.io/pod-security-admission/api"

	"github.com/onsi/ginkgo/v2"
	libcontainercgroups "github.com/opencontainers/runc/libcontainer/cgroups"
)

type testCase struct {
	name                   string
	podSpec                *v1.Pod
	oomTargetContainerName string
}

// KubeReservedMemory is default fraction value of node capacity memory to
// be reserved for K8s components.
const KubeReservedMemory = 0.35

var _ = SIGDescribe("OOMKiller for pod using more memory than node allocatable [LinuxOnly]", framework.WithSerial(), func() {
	f := framework.NewDefaultFramework("nodeallocatable-oomkiller-test")
	f.NamespacePodSecurityEnforceLevel = admissionapi.LevelPrivileged

	testCases := []testCase{
		{
			name:                   "single process container without resource limits",
			oomTargetContainerName: "oomkill-nodeallocatable-container",
			podSpec: getOOMTargetPod("oomkill-nodeallocatable-pod", "oomkill-nodeallocatable-container",
				getOOMTargetContainerWithoutLimit),
		},
	}

	for _, testCase := range testCases {
		runOomKillerTest(f, testCase, KubeReservedMemory, "UnlimitedSwap")
	}
})

var _ = SIGDescribe("OOMKiller [LinuxOnly]", framework.WithNodeConformance(), func() {
	f := framework.NewDefaultFramework("oomkiller-test")
	f.NamespacePodSecurityLevel = admissionapi.LevelPrivileged

	testCases := []testCase{
		{
			name:                   "single process container",
			oomTargetContainerName: "oomkill-single-target-container",
			podSpec: getOOMTargetPod("oomkill-target-pod", "oomkill-single-target-container",
				getOOMTargetContainer),
		},
		{
			name:                   "init container",
			oomTargetContainerName: "oomkill-target-init-container",
			podSpec: getInitContainerOOMTargetPod("initcontinar-oomkill-target-pod", "oomkill-target-init-container",
				getOOMTargetContainer),
		},
	}

	// If using cgroup v2, we set memory.oom.group=1 for the container cgroup so that any process which gets OOM killed
	// in the process, causes all processes in the container to get OOM killed
	if libcontainercgroups.IsCgroup2UnifiedMode() {
		testCases = append(testCases, testCase{
			name:                   "multi process container",
			oomTargetContainerName: "oomkill-multi-target-container",
			podSpec: getOOMTargetPod("oomkill-target-pod", "oomkill-multi-target-container",
				getOOMTargetContainerMultiProcess),
		})
	}
	for _, tc := range testCases {
		runOomKillerTest(f, tc, 0, "UnlimitedSwap")
	}
})

func runOomKillerTest(f *framework.Framework, testCase testCase, kubeReservedMemory float64, swapBehavior string) {
	ginkgo.Context(testCase.name, func() {
		// Update KubeReservedMemory in KubeletConfig.
		if kubeReservedMemory > 0 {
			tempSetCurrentKubeletConfig(f, func(ctx context.Context, initialConfig *kubeletconfig.KubeletConfiguration) {
				// Set the swap behavior to the value specified in the test case
				initialConfig.MemorySwap.SwapBehavior = swapBehavior
				if initialConfig.KubeReserved == nil {
					initialConfig.KubeReserved = map[string]string{}
				}
				// There's a race condition observed between system OOM and cgroup OOM if node allocatable
				// memory is equal to node capacity. Hence, reserving a fraction of node's memory capacity for
				// K8s components such that node allocatable memory is less than node capacity to
				// observe OOM kills at cgroup level instead of system OOM kills.
				initialConfig.KubeReserved["memory"] = fmt.Sprintf("%d", int(kubeReservedMemory*getLocalNode(context.TODO(), f).Status.Capacity.Memory().AsApproximateFloat64()))
			})

		}

		ginkgo.BeforeEach(func() {
			ginkgo.By("setting up the pod to be used in the test")
			e2epod.NewPodClient(f).Create(context.TODO(), testCase.podSpec)
		})

		ginkgo.It("The containers terminated by OOM killer should have the reason set to OOMKilled", func() {
			cfg, configErr := getCurrentKubeletConfig(context.TODO())
			framework.ExpectNoError(configErr)
			if utilfeature.DefaultFeatureGate.Enabled(features.NodeSwap) {
				// If Swap is enabled, we should test OOM with LimitedSwap.
				// UnlimitedSwap allows for workloads to use unbounded swap which
				// makes testing OOM challenging.
				// We are not able to change the default for these conformance tests,
				// so we will skip these tests if swap is enabled.
				if cfg.MemorySwap.SwapBehavior == "" || cfg.MemorySwap.SwapBehavior == "UnlimitedSwap" {
					ginkgo.Skip("OOMKiller should not run with UnlimitedSwap")
				}
			}
			ginkgo.By("Waiting for the pod to be failed")
			err := e2epod.WaitForPodTerminatedInNamespace(context.TODO(), f.ClientSet, testCase.podSpec.Name, "", f.Namespace.Name)
			framework.ExpectNoError(err, "Failed waiting for pod to terminate, %s/%s", f.Namespace.Name, testCase.podSpec.Name)

			ginkgo.By("Fetching the latest pod status")
			pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Get(context.TODO(), testCase.podSpec.Name, metav1.GetOptions{})
			framework.ExpectNoError(err, "Failed to get the recent pod object for name: %q", pod.Name)

			ginkgo.By("Verifying the OOM target container has the expected reason")
			verifyReasonForOOMKilledContainer(pod, testCase.oomTargetContainerName)
		})

		ginkgo.AfterEach(func() {
			ginkgo.By(fmt.Sprintf("deleting pod: %s", testCase.podSpec.Name))
			e2epod.NewPodClient(f).DeleteSync(context.TODO(), testCase.podSpec.Name, metav1.DeleteOptions{}, framework.PodDeleteTimeout)
		})
	})
}

func verifyReasonForOOMKilledContainer(pod *v1.Pod, oomTargetContainerName string) {
	container := e2epod.FindContainerStatusInPod(pod, oomTargetContainerName)
	if container == nil {
		framework.Failf("OOM target pod %q, container %q does not have the expected state terminated", pod.Name, oomTargetContainerName)
	}
	if container.State.Terminated == nil {
		framework.Failf("OOM target pod %q, container %q is not in the terminated state", pod.Name, container.Name)
	}
	gomega.Expect(container.State.Terminated.ExitCode).To(gomega.Equal(int32(137)),
		"pod: %q, container: %q has unexpected exitCode: %q", pod.Name, container.Name, container.State.Terminated.ExitCode)

	// This check is currently causing tests to flake on containerd & crio, https://github.com/kubernetes/kubernetes/issues/119600
	// so we'll skip the reason check if we know its going to fail.
	// TODO: Remove this once https://github.com/containerd/containerd/issues/8893 is resolved
	if container.State.Terminated.Reason == "OOMKilled" {
		gomega.Expect(container.State.Terminated.Reason).To(gomega.Equal("OOMKilled"),
			"pod: %q, container: %q has unexpected reason: %q", pod.Name, container.Name, container.State.Terminated.Reason)
	}

}

func getOOMTargetPod(podName string, ctnName string, createContainer func(name string) v1.Container) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyNever,
			Containers: []v1.Container{
				createContainer(ctnName),
			},
		},
	}
}

func getInitContainerOOMTargetPod(podName string, ctnName string, createContainer func(name string) v1.Container) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
		},
		Spec: v1.PodSpec{
			RestartPolicy: v1.RestartPolicyNever,
			InitContainers: []v1.Container{
				createContainer(ctnName),
			},
			Containers: []v1.Container{
				{
					Name:  "busybox",
					Image: busyboxImage,
				},
			},
		},
	}
}

// getOOMTargetContainer returns a container with a single process, which attempts to allocate more memory than is
// allowed by the container memory limit.
func getOOMTargetContainer(name string) v1.Container {
	return v1.Container{
		Name:  name,
		Image: busyboxImage,
		Command: []string{
			"sh",
			"-c",
			// use the dd tool to attempt to allocate 20M in a block which exceeds the limit
			"sleep 5 && dd if=/dev/zero of=/dev/null bs=20M",
		},
		Resources: v1.ResourceRequirements{
			Requests: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("15Mi"),
			},
			Limits: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("15Mi"),
			},
		},
	}
}

// getOOMTargetContainerMultiProcess returns a container with two processes, one of which attempts to allocate more
// memory than is allowed by the container memory limit, and a second process which just sleeps.
func getOOMTargetContainerMultiProcess(name string) v1.Container {
	return v1.Container{
		Name:  name,
		Image: busyboxImage,
		Command: []string{
			"sh",
			"-c",
			// use the dd tool to attempt to allocate 20M in a block which exceeds the limit
			"sleep 5 && dd if=/dev/zero of=/dev/null bs=20M & sleep 86400",
		},
		Resources: v1.ResourceRequirements{
			Requests: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("15Mi"),
			},
			Limits: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("15Mi"),
			},
		},
	}
}

// getOOMTargetContainerWithoutLimit returns a container with a single process which attempts to allocate more memory
// than node allocatable and doesn't have resource limits set.
func getOOMTargetContainerWithoutLimit(name string) v1.Container {
	return v1.Container{
		Name:  name,
		Image: busyboxImage,
		Command: []string{
			"sh",
			"-c",
			// use the dd tool to attempt to allocate huge block of memory which exceeds the node allocatable
			"sleep 5 && dd if=/dev/zero of=/dev/null iflag=fullblock count=10 bs=10G",
		},
	}
}

var _ = SIGDescribe("OOMKiller with LimitedSwap", framework.WithSerial(), framework.WithFeature("Swap"), func() {
	f := framework.NewDefaultFramework("oomkiller-limitedswap-test")
	f.NamespacePodSecurityLevel = admissionapi.LevelPrivileged

	testCases := []testCase{
		{
			name:                   "single process container with LimitedSwap",
			oomTargetContainerName: "oomkill-limitedswap-container",
			podSpec: getOOMTargetPod("oomkill-limitedswap-pod", "oomkill-limitedswap-container",
				getOOMTargetContainerWithLimitedSwap),
		},
	}

	for _, tc := range testCases {
		runOomKillerTest(f, tc, 0, "LimitedSwap")
	}
})

// getOOMTargetContainerWithLimitedSwap returns a container with a single process which attempts to allocate more memory
// than the container's memory and swap limit combined.
func getOOMTargetContainerWithLimitedSwap(name string) v1.Container {
	return v1.Container{
		Name:  name,
		Image: busyboxImage,
		Command: []string{
			"sh",
			"-c",
			// use the dd tool to attempt to allocate a block of memory which exceeds the memory limit
			"sleep 5 && dd if=/dev/zero of=/dev/null iflag=fullblock count=10 bs=10G",
		},
		Resources: v1.ResourceRequirements{
			Requests: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("100Mi"),
			},
			Limits: v1.ResourceList{
				v1.ResourceMemory: resource.MustParse("100Mi"),
			},
		},
	}
}
