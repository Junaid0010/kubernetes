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

package sysctl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/kubelet/lifecycle"
)

func TestNewAllowlist(t *testing.T) {
	type Test struct {
		sysctls []string
		err     bool
	}
	for _, test := range []Test{
		{sysctls: []string{"*", "kernel.sem"}, err: true},
		{sysctls: []string{"kernel.msg*", "kernel.sem"}},
		{sysctls: []string{"kernel/msg*", "kernel/sem"}},
		{sysctls: []string{" kernel.msg*"}, err: true},
		{sysctls: []string{"kernel.msg* "}, err: true},
		{sysctls: []string{"net.-"}, err: true},
		{sysctls: []string{"net.*.foo"}, err: true},
		{sysctls: []string{"net.*/foo"}, err: true},
		{sysctls: []string{"foo"}, err: true},
	} {
		_, err := NewAllowlist(append(SafeSysctlAllowlist(), test.sysctls...))
		if test.err && err == nil {
			t.Errorf("expected an error creating a allowlist for %v", test.sysctls)
		} else if !test.err && err != nil {
			t.Errorf("got unexpected error creating a allowlist for %v: %v", test.sysctls, err)
		}
	}
}

func TestAllowlist(t *testing.T) {
	type Test struct {
		sysctl           string
		hostNet, hostIPC bool
	}
	valid := []Test{
		{sysctl: "kernel.shm_rmid_forced"},
		{sysctl: "kernel/shm_rmid_forced"},
		{sysctl: "net.ipv4.ip_local_port_range"},
		{sysctl: "kernel.msgmax"},
		{sysctl: "kernel.sem"},
		{sysctl: "kernel/sem"},
	}
	invalid := []Test{
		{sysctl: "kernel.shm_rmid_forced", hostIPC: true},
		{sysctl: "net.ipv4.ip_local_port_range", hostNet: true},
		{sysctl: "foo"},
		{sysctl: "net.a.b.c", hostNet: false},
		{sysctl: "net.ipv4.ip_local_port_range.a.b.c", hostNet: false},
		{sysctl: "kernel.msgmax", hostIPC: true},
		{sysctl: "net.msgmax", hostNet: true},
		{sysctl: "kernel.sem", hostIPC: true},
	}

	w, err := NewAllowlist(append(SafeSysctlAllowlist(), "kernel.msg*", "kernel.sem", "net.msg*"))
	if err != nil {
		t.Fatalf("failed to create allowlist: %v", err)
	}

	for _, test := range valid {
		if err := w.validateSysctl(test.sysctl, test.hostNet, test.hostIPC); err != nil {
			t.Errorf("expected to be allowlisted: %+v, got: %v", test, err)
		}
	}

	for _, test := range invalid {
		if err := w.validateSysctl(test.sysctl, test.hostNet, test.hostIPC); err == nil {
			t.Errorf("expected to be rejected: %+v", test)
		}
	}
}

func TestAdmit(t *testing.T) {
	tests := []struct {
		name           string
		attrs          *lifecycle.PodAdmitAttributes
		expectedResult lifecycle.PodAdmitResult
	}{
		{
			name:  "no Sysctls info",
			attrs: &lifecycle.PodAdmitAttributes{Pod: &metav1.Pod{}, OtherPods: []*metav1.Pod{}},
			expectedResult: lifecycle.PodAdmitResult{
				Admit: true,
			},
		},
		{
			name: "valid sysctls",
			attrs: &lifecycle.PodAdmitAttributes{Pod: &metav1.Pod{
				Spec: metav1.PodSpec{
					SecurityContext: &metav1.PodSecurityContext{
						Sysctls: []metav1.Sysctl{
							{Name: "kernel.shm_rmid_forced"},
						},
					},
					HostIPC:     false,
					HostNetwork: false,
				},
			}, OtherPods: []*metav1.Pod{}},
			expectedResult: lifecycle.PodAdmitResult{
				Admit: true,
			},
		},
		{
			name: "Invalid sysctls",
			attrs: &lifecycle.PodAdmitAttributes{Pod: &metav1.Pod{
				Spec: metav1.PodSpec{
					SecurityContext: &metav1.PodSecurityContext{
						Sysctls: []metav1.Sysctl{
							{Name: "kernel.shm_rmid_forced"},
						},
					},
					HostIPC:     true,
					HostNetwork: false,
				},
			}, OtherPods: []*metav1.Pod{}},
			expectedResult: lifecycle.PodAdmitResult{
				Admit:   false,
				Reason:  ForbiddenReason,
				Message: fmt.Sprint("forbidden sysctl: \"kernel.shm_rmid_forced\" not allowed with host ipc enabled"),
			},
		},
	}
	w, err := NewAllowlist(append(SafeSysctlAllowlist(), "kernel.msg*", "kernel.sem"))
	if err != nil {
		t.Fatalf("failed to create allowlist: %v", err)
	}

	for _, test := range tests {
		result := w.Admit(test.attrs)
		require.Equal(t, test.expectedResult, result)
	}
}
