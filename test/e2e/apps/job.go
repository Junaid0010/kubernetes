/*
Copyright 2017 The Kubernetes Authors.

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

package apps

import (
	"fmt"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	batchinternal "k8s.io/kubernetes/pkg/apis/batch"
	"k8s.io/kubernetes/test/e2e/framework"
	jobutil "k8s.io/kubernetes/test/e2e/framework/job"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = SIGDescribe("Job", func() {
	f := framework.NewDefaultFramework("job")
	parallelism := int32(2)
	completions := int32(4)
	backoffLimit := int32(6) // default value

	/*
		Release : v1.15
		Testname: Jobs, pod tasks successful execution
		Description: Create a job. Execute all the tasks successfully. After pod execution, the Job MUST reach completion.
	*/
	framework.ConformanceIt("should run a job to completion when tasks succeed", func() {
		By("Creating a job")
		job := jobutil.NewTestJob("succeed", "all-succeed", v1.RestartPolicyNever, parallelism, completions, nil, backoffLimit)
		job, err := jobutil.CreateJob(f.ClientSet, f.Namespace.Name, job)
		Expect(err).NotTo(HaveOccurred(), "failed to create job in namespace: %s", f.Namespace.Name)

		By("Ensuring job reaches completions")
		err = jobutil.WaitForJobComplete(f.ClientSet, f.Namespace.Name, job.Name, completions)
		Expect(err).NotTo(HaveOccurred(), "failed to ensure job completion in namespace: %s", f.Namespace.Name)
	})

	// Pods sometimes fail, but eventually succeed.
	It("should run a job to completion when tasks sometimes fail and are locally restarted", func() {
		By("Creating a job")
		// One failure, then a success, local restarts.
		// We can't use the random failure approach used by the
		// non-local test below, because kubelet will throttle
		// frequently failing containers in a given pod, ramping
		// up to 5 minutes between restarts, making test timeouts
		// due to successive failures too likely with a reasonable
		// test timeout.
		job := jobutil.NewTestJob("failOnce", "fail-once-local", v1.RestartPolicyOnFailure, parallelism, completions, nil, backoffLimit)
		job, err := jobutil.CreateJob(f.ClientSet, f.Namespace.Name, job)
		Expect(err).NotTo(HaveOccurred(), "failed to create job in namespace: %s", f.Namespace.Name)

		By("Ensuring job reaches completions")
		err = jobutil.WaitForJobComplete(f.ClientSet, f.Namespace.Name, job.Name, completions)
		Expect(err).NotTo(HaveOccurred(), "failed to ensure job completion in namespace: %s", f.Namespace.Name)
	})

	// Pods sometimes fail, but eventually succeed, after pod restarts
	It("should run a job to completion when tasks sometimes fail and are not locally restarted", func() {
		By("Creating a job")
		// 50% chance of container success, local restarts.
		// Can't use the failOnce approach because that relies
		// on an emptyDir, which is not preserved across new pods.
		// Worst case analysis: 15 failures, each taking 1 minute to
		// run due to some slowness, 1 in 2^15 chance of happening,
		// causing test flake.  Should be very rare.
		// With the introduction of backoff limit and high failure rate this
		// is hitting its timeout, the 3 is a reasonable that should make this
		// test less flaky, for now.
		job := jobutil.NewTestJob("randomlySucceedOrFail", "rand-non-local", v1.RestartPolicyNever, parallelism, 3, nil, 999)
		job, err := jobutil.CreateJob(f.ClientSet, f.Namespace.Name, job)
		Expect(err).NotTo(HaveOccurred(), "failed to create job in namespace: %s", f.Namespace.Name)

		By("Ensuring job reaches completions")
		err = jobutil.WaitForJobComplete(f.ClientSet, f.Namespace.Name, job.Name, *job.Spec.Completions)
		Expect(err).NotTo(HaveOccurred(), "failed to ensure job completion in namespace: %s", f.Namespace.Name)
	})

	It("should exceed active deadline", func() {
		By("Creating a job")
		var activeDeadlineSeconds int64 = 1
		job := jobutil.NewTestJob("notTerminate", "exceed-active-deadline", v1.RestartPolicyNever, parallelism, completions, &activeDeadlineSeconds, backoffLimit)
		job, err := jobutil.CreateJob(f.ClientSet, f.Namespace.Name, job)
		Expect(err).NotTo(HaveOccurred(), "failed to create job in namespace: %s", f.Namespace.Name)
		By("Ensuring job past active deadline")
		err = jobutil.WaitForJobFailure(f.ClientSet, f.Namespace.Name, job.Name, time.Duration(activeDeadlineSeconds+10)*time.Second, "DeadlineExceeded")
		Expect(err).NotTo(HaveOccurred(), "failed to ensure job past active deadline in namespace: %s", f.Namespace.Name)
	})

	/*
		Release : v1.15
		Testname: Jobs, active pods, graceful termination
		Description: Create a job. Ensure the active pods reflect paralellism in the namespace and delete the job. Job MUST be deleted successfully.
	*/
	framework.ConformanceIt("should delete a job", func() {
		By("Creating a job")
		job := jobutil.NewTestJob("notTerminate", "foo", v1.RestartPolicyNever, parallelism, completions, nil, backoffLimit)
		job, err := jobutil.CreateJob(f.ClientSet, f.Namespace.Name, job)
		Expect(err).NotTo(HaveOccurred(), "failed to create job in namespace: %s", f.Namespace.Name)

		By("Ensuring active pods == parallelism")
		err = jobutil.WaitForAllJobPodsRunning(f.ClientSet, f.Namespace.Name, job.Name, parallelism)
		Expect(err).NotTo(HaveOccurred(), "failed to ensure active pods == parallelism in namespace: %s", f.Namespace.Name)

		By("delete a job")
		framework.ExpectNoError(framework.DeleteResourceAndWaitForGC(f.ClientSet, batchinternal.Kind("Job"), f.Namespace.Name, job.Name))

		By("Ensuring job was deleted")
		_, err = jobutil.GetJob(f.ClientSet, f.Namespace.Name, job.Name)
		Expect(err).To(HaveOccurred(), "failed to ensure job %s was deleted in namespace: %s", job.Name, f.Namespace.Name)
		Expect(errors.IsNotFound(err)).To(BeTrue())
	})

	It("should adopt matching orphans and release non-matching pods", func() {
		By("Creating a job")
		job := jobutil.NewTestJob("notTerminate", "adopt-release", v1.RestartPolicyNever, parallelism, completions, nil, backoffLimit)
		// Replace job with the one returned from Create() so it has the UID.
		// Save Kind since it won't be populated in the returned job.
		kind := job.Kind
		job, err := jobutil.CreateJob(f.ClientSet, f.Namespace.Name, job)
		Expect(err).NotTo(HaveOccurred(), "failed to create job in namespace: %s", f.Namespace.Name)
		job.Kind = kind

		By("Ensuring active pods == parallelism")
		err = jobutil.WaitForAllJobPodsRunning(f.ClientSet, f.Namespace.Name, job.Name, parallelism)
		Expect(err).NotTo(HaveOccurred(), "failed to ensure active pods == parallelism in namespace: %s", f.Namespace.Name)

		By("Orphaning one of the Job's Pods")
		pods, err := jobutil.GetJobPods(f.ClientSet, f.Namespace.Name, job.Name)
		Expect(err).NotTo(HaveOccurred(), "failed to get PodList for job %s in namespace: %s", job.Name, f.Namespace.Name)
		Expect(pods.Items).To(HaveLen(int(parallelism)))
		pod := pods.Items[0]
		f.PodClient().Update(pod.Name, func(pod *v1.Pod) {
			pod.OwnerReferences = nil
		})

		By("Checking that the Job readopts the Pod")
		Expect(framework.WaitForPodCondition(f.ClientSet, pod.Namespace, pod.Name, "adopted", jobutil.JobTimeout,
			func(pod *v1.Pod) (bool, error) {
				controllerRef := metav1.GetControllerOf(pod)
				if controllerRef == nil {
					return false, nil
				}
				if controllerRef.Kind != job.Kind || controllerRef.Name != job.Name || controllerRef.UID != job.UID {
					return false, fmt.Errorf("pod has wrong controllerRef: got %v, want %v", controllerRef, job)
				}
				return true, nil
			},
		)).To(Succeed(), "wait for pod %q to be readopted", pod.Name)

		By("Removing the labels from the Job's Pod")
		f.PodClient().Update(pod.Name, func(pod *v1.Pod) {
			pod.Labels = nil
		})

		By("Checking that the Job releases the Pod")
		Expect(framework.WaitForPodCondition(f.ClientSet, pod.Namespace, pod.Name, "released", jobutil.JobTimeout,
			func(pod *v1.Pod) (bool, error) {
				controllerRef := metav1.GetControllerOf(pod)
				if controllerRef != nil {
					return false, nil
				}
				return true, nil
			},
		)).To(Succeed(), "wait for pod %q to be released", pod.Name)
	})

	It("should exceed backoffLimit", func() {
		By("Creating a job")
		backoff := 1
		job := jobutil.NewTestJob("fail", "backofflimit", v1.RestartPolicyNever, 1, 1, nil, int32(backoff))
		job, err := jobutil.CreateJob(f.ClientSet, f.Namespace.Name, job)
		Expect(err).NotTo(HaveOccurred(), "failed to create job in namespace: %s", f.Namespace.Name)
		By("Ensuring job exceed backofflimit")

		err = jobutil.WaitForJobFailure(f.ClientSet, f.Namespace.Name, job.Name, jobutil.JobTimeout, "BackoffLimitExceeded")
		Expect(err).NotTo(HaveOccurred(), "failed to ensure job exceed backofflimit in namespace: %s", f.Namespace.Name)

		By(fmt.Sprintf("Checking that %d pod created and status is failed", backoff+1))
		pods, err := jobutil.GetJobPods(f.ClientSet, f.Namespace.Name, job.Name)
		Expect(err).NotTo(HaveOccurred(), "failed to get PodList for job %s in namespace: %s", job.Name, f.Namespace.Name)
		// Expect(pods.Items).To(HaveLen(backoff + 1))
		// due to NumRequeus not being stable enough, especially with failed status
		// updates we need to allow more than backoff+1
		// TODO revert this back to above when https://github.com/kubernetes/kubernetes/issues/64787 gets fixed
		if len(pods.Items) < backoff+1 {
			framework.Failf("Not enough pod created expected at least %d, got %#v", backoff+1, pods.Items)
		}
		for _, pod := range pods.Items {
			Expect(pod.Status.Phase).To(Equal(v1.PodFailed))
		}
	})
})
