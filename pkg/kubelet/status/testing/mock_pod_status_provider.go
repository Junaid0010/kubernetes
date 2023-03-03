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

// Code generated by MockGen. DO NOT EDIT.
// Source: status_manager.go

// Package testing is a generated GoMock package.
package testing

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	container "k8s.io/kubernetes/pkg/kubelet/container"
)

// MockPodStatusProvider is a mock of PodStatusProvider interface.
type MockPodStatusProvider struct {
	ctrl     *gomock.Controller
	recorder *MockPodStatusProviderMockRecorder
}

// MockPodStatusProviderMockRecorder is the mock recorder for MockPodStatusProvider.
type MockPodStatusProviderMockRecorder struct {
	mock *MockPodStatusProvider
}

// NewMockPodStatusProvider creates a new mock instance.
func NewMockPodStatusProvider(ctrl *gomock.Controller) *MockPodStatusProvider {
	mock := &MockPodStatusProvider{ctrl: ctrl}
	mock.recorder = &MockPodStatusProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodStatusProvider) EXPECT() *MockPodStatusProviderMockRecorder {
	return m.recorder
}

// GetPodStatus mocks base method.
func (m *MockPodStatusProvider) GetPodStatus(uid types.UID) (v1.PodStatus, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodStatus", uid)
	ret0, _ := ret[0].(v1.PodStatus)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPodStatus indicates an expected call of GetPodStatus.
func (mr *MockPodStatusProviderMockRecorder) GetPodStatus(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodStatus", reflect.TypeOf((*MockPodStatusProvider)(nil).GetPodStatus), uid)
}

// MockPodDeletionSafetyProvider is a mock of PodDeletionSafetyProvider interface.
type MockPodDeletionSafetyProvider struct {
	ctrl     *gomock.Controller
	recorder *MockPodDeletionSafetyProviderMockRecorder
}

// MockPodDeletionSafetyProviderMockRecorder is the mock recorder for MockPodDeletionSafetyProvider.
type MockPodDeletionSafetyProviderMockRecorder struct {
	mock *MockPodDeletionSafetyProvider
}

// NewMockPodDeletionSafetyProvider creates a new mock instance.
func NewMockPodDeletionSafetyProvider(ctrl *gomock.Controller) *MockPodDeletionSafetyProvider {
	mock := &MockPodDeletionSafetyProvider{ctrl: ctrl}
	mock.recorder = &MockPodDeletionSafetyProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodDeletionSafetyProvider) EXPECT() *MockPodDeletionSafetyProviderMockRecorder {
	return m.recorder
}

// PodCouldHaveRunningContainers mocks base method.
func (m *MockPodDeletionSafetyProvider) PodCouldHaveRunningContainers(pod *v1.Pod) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodCouldHaveRunningContainers", pod)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PodCouldHaveRunningContainers indicates an expected call of PodCouldHaveRunningContainers.
func (mr *MockPodDeletionSafetyProviderMockRecorder) PodCouldHaveRunningContainers(pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodCouldHaveRunningContainers", reflect.TypeOf((*MockPodDeletionSafetyProvider)(nil).PodCouldHaveRunningContainers), pod)
}

// PodMightNeedToUnprepareResources mocks base method.
func (m *MockPodDeletionSafetyProvider) PodMightNeedToUnprepareResources(UID types.UID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodMightNeedToUnprepareResources", UID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PodMightNeedToUnprepareResources indicates an expected call of PodMightNeedToUnprepareResources.
func (mr *MockPodDeletionSafetyProviderMockRecorder) PodMightNeedToUnprepareResources(UID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodMightNeedToUnprepareResources", reflect.TypeOf((*MockPodDeletionSafetyProvider)(nil).PodMightNeedToUnprepareResources), UID)
}

// MockPodStartupLatencyStateHelper is a mock of PodStartupLatencyStateHelper interface.
type MockPodStartupLatencyStateHelper struct {
	ctrl     *gomock.Controller
	recorder *MockPodStartupLatencyStateHelperMockRecorder
}

// MockPodStartupLatencyStateHelperMockRecorder is the mock recorder for MockPodStartupLatencyStateHelper.
type MockPodStartupLatencyStateHelperMockRecorder struct {
	mock *MockPodStartupLatencyStateHelper
}

// NewMockPodStartupLatencyStateHelper creates a new mock instance.
func NewMockPodStartupLatencyStateHelper(ctrl *gomock.Controller) *MockPodStartupLatencyStateHelper {
	mock := &MockPodStartupLatencyStateHelper{ctrl: ctrl}
	mock.recorder = &MockPodStartupLatencyStateHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodStartupLatencyStateHelper) EXPECT() *MockPodStartupLatencyStateHelperMockRecorder {
	return m.recorder
}

// DeletePodStartupState mocks base method.
func (m *MockPodStartupLatencyStateHelper) DeletePodStartupState(podUID types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePodStartupState", podUID)
}

// DeletePodStartupState indicates an expected call of DeletePodStartupState.
func (mr *MockPodStartupLatencyStateHelperMockRecorder) DeletePodStartupState(podUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePodStartupState", reflect.TypeOf((*MockPodStartupLatencyStateHelper)(nil).DeletePodStartupState), podUID)
}

// RecordStatusUpdated mocks base method.
func (m *MockPodStartupLatencyStateHelper) RecordStatusUpdated(pod *v1.Pod) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RecordStatusUpdated", pod)
}

// RecordStatusUpdated indicates an expected call of RecordStatusUpdated.
func (mr *MockPodStartupLatencyStateHelperMockRecorder) RecordStatusUpdated(pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordStatusUpdated", reflect.TypeOf((*MockPodStartupLatencyStateHelper)(nil).RecordStatusUpdated), pod)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// GetContainerResourceAllocation mocks base method.
func (m *MockManager) GetContainerResourceAllocation(podUID, containerName string) (v1.ResourceList, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerResourceAllocation", podUID, containerName)
	ret0, _ := ret[0].(v1.ResourceList)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetContainerResourceAllocation indicates an expected call of GetContainerResourceAllocation.
func (mr *MockManagerMockRecorder) GetContainerResourceAllocation(podUID, containerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerResourceAllocation", reflect.TypeOf((*MockManager)(nil).GetContainerResourceAllocation), podUID, containerName)
}

// GetPodResizeStatus mocks base method.
func (m *MockManager) GetPodResizeStatus(podUID string) (v1.PodResizeStatus, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodResizeStatus", podUID)
	ret0, _ := ret[0].(v1.PodResizeStatus)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPodResizeStatus indicates an expected call of GetPodResizeStatus.
func (mr *MockManagerMockRecorder) GetPodResizeStatus(podUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodResizeStatus", reflect.TypeOf((*MockManager)(nil).GetPodResizeStatus), podUID)
}

// GetPodStatus mocks base method.
func (m *MockManager) GetPodStatus(uid types.UID) (v1.PodStatus, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodStatus", uid)
	ret0, _ := ret[0].(v1.PodStatus)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPodStatus indicates an expected call of GetPodStatus.
func (mr *MockManagerMockRecorder) GetPodStatus(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodStatus", reflect.TypeOf((*MockManager)(nil).GetPodStatus), uid)
}

// RemoveOrphanedStatuses mocks base method.
func (m *MockManager) RemoveOrphanedStatuses(podUIDs map[types.UID]bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveOrphanedStatuses", podUIDs)
}

// RemoveOrphanedStatuses indicates an expected call of RemoveOrphanedStatuses.
func (mr *MockManagerMockRecorder) RemoveOrphanedStatuses(podUIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOrphanedStatuses", reflect.TypeOf((*MockManager)(nil).RemoveOrphanedStatuses), podUIDs)
}

// SetContainerReadiness mocks base method.
func (m *MockManager) SetContainerReadiness(podUID types.UID, containerID container.ContainerID, ready bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetContainerReadiness", podUID, containerID, ready)
}

// SetContainerReadiness indicates an expected call of SetContainerReadiness.
func (mr *MockManagerMockRecorder) SetContainerReadiness(podUID, containerID, ready interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContainerReadiness", reflect.TypeOf((*MockManager)(nil).SetContainerReadiness), podUID, containerID, ready)
}

// SetContainerStartup mocks base method.
func (m *MockManager) SetContainerStartup(podUID types.UID, containerID container.ContainerID, started bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetContainerStartup", podUID, containerID, started)
}

// SetContainerStartup indicates an expected call of SetContainerStartup.
func (mr *MockManagerMockRecorder) SetContainerStartup(podUID, containerID, started interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContainerStartup", reflect.TypeOf((*MockManager)(nil).SetContainerStartup), podUID, containerID, started)
}

// SetPodAllocation mocks base method.
func (m *MockManager) SetPodAllocation(pod *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPodAllocation", pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPodAllocation indicates an expected call of SetPodAllocation.
func (mr *MockManagerMockRecorder) SetPodAllocation(pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPodAllocation", reflect.TypeOf((*MockManager)(nil).SetPodAllocation), pod)
}

// SetPodResizeStatus mocks base method.
func (m *MockManager) SetPodResizeStatus(podUID types.UID, resize v1.PodResizeStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPodResizeStatus", podUID, resize)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPodResizeStatus indicates an expected call of SetPodResizeStatus.
func (mr *MockManagerMockRecorder) SetPodResizeStatus(podUID, resize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPodResizeStatus", reflect.TypeOf((*MockManager)(nil).SetPodResizeStatus), podUID, resize)
}

// SetPodStatus mocks base method.
func (m *MockManager) SetPodStatus(pod *v1.Pod, status v1.PodStatus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPodStatus", pod, status)
}

// SetPodStatus indicates an expected call of SetPodStatus.
func (mr *MockManagerMockRecorder) SetPodStatus(pod, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPodStatus", reflect.TypeOf((*MockManager)(nil).SetPodStatus), pod, status)
}

// Start mocks base method.
func (m *MockManager) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start))
}

// TerminatePod mocks base method.
func (m *MockManager) TerminatePod(pod *v1.Pod) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TerminatePod", pod)
}

// TerminatePod indicates an expected call of TerminatePod.
func (mr *MockManagerMockRecorder) TerminatePod(pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminatePod", reflect.TypeOf((*MockManager)(nil).TerminatePod), pod)
}
