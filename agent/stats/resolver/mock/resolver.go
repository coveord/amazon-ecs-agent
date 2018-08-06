// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/stats/resolver (interfaces: ContainerMetadataResolver)

// Package mock_resolver is a generated GoMock package.
package mock_resolver

import (
	reflect "reflect"

	container "github.com/aws/amazon-ecs-agent/agent/api/container"
	task "github.com/aws/amazon-ecs-agent/agent/api/task"
	gomock "github.com/golang/mock/gomock"
)

// MockContainerMetadataResolver is a mock of ContainerMetadataResolver interface
type MockContainerMetadataResolver struct {
	ctrl     *gomock.Controller
	recorder *MockContainerMetadataResolverMockRecorder
}

// MockContainerMetadataResolverMockRecorder is the mock recorder for MockContainerMetadataResolver
type MockContainerMetadataResolverMockRecorder struct {
	mock *MockContainerMetadataResolver
}

// NewMockContainerMetadataResolver creates a new mock instance
func NewMockContainerMetadataResolver(ctrl *gomock.Controller) *MockContainerMetadataResolver {
	mock := &MockContainerMetadataResolver{ctrl: ctrl}
	mock.recorder = &MockContainerMetadataResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContainerMetadataResolver) EXPECT() *MockContainerMetadataResolverMockRecorder {
	return m.recorder
}

// ResolveContainer mocks base method
func (m *MockContainerMetadataResolver) ResolveContainer(arg0 string) (*container.DockerContainer, error) {
	ret := m.ctrl.Call(m, "ResolveContainer", arg0)
	ret0, _ := ret[0].(*container.DockerContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveContainer indicates an expected call of ResolveContainer
func (mr *MockContainerMetadataResolverMockRecorder) ResolveContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveContainer", reflect.TypeOf((*MockContainerMetadataResolver)(nil).ResolveContainer), arg0)
}

// ResolveTask mocks base method
func (m *MockContainerMetadataResolver) ResolveTask(arg0 string) (*task.Task, error) {
	ret := m.ctrl.Call(m, "ResolveTask", arg0)
	ret0, _ := ret[0].(*task.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveTask indicates an expected call of ResolveTask
func (mr *MockContainerMetadataResolverMockRecorder) ResolveTask(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveTask", reflect.TypeOf((*MockContainerMetadataResolver)(nil).ResolveTask), arg0)
}
