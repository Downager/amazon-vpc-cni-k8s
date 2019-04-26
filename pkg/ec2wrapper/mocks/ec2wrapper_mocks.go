// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/ec2wrapper (interfaces: EC2)

// Package mock_ec2wrapper is a generated GoMock package.
package mock_ec2wrapper

import (
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MockEC2 is a mock of EC2 interface
type MockEC2 struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MockRecorder
}

// MockEC2MockRecorder is the mock recorder for MockEC2
type MockEC2MockRecorder struct {
	mock *MockEC2
}

// NewMockEC2 creates a new mock instance
func NewMockEC2(ctrl *gomock.Controller) *MockEC2 {
	mock := &MockEC2{ctrl: ctrl}
	mock.recorder = &MockEC2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2) EXPECT() *MockEC2MockRecorder {
	return m.recorder
}

// AssignPrivateIpAddresses mocks base method
func (m *MockEC2) AssignPrivateIpAddresses(arg0 *ec2.AssignPrivateIpAddressesInput) (*ec2.AssignPrivateIpAddressesOutput, error) {
	ret := m.ctrl.Call(m, "AssignPrivateIpAddresses", arg0)
	ret0, _ := ret[0].(*ec2.AssignPrivateIpAddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignPrivateIpAddresses indicates an expected call of AssignPrivateIpAddresses
func (mr *MockEC2MockRecorder) AssignPrivateIpAddresses(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignPrivateIpAddresses", reflect.TypeOf((*MockEC2)(nil).AssignPrivateIpAddresses), arg0)
}

// AttachNetworkInterface mocks base method
func (m *MockEC2) AttachNetworkInterface(arg0 *ec2.AttachNetworkInterfaceInput) (*ec2.AttachNetworkInterfaceOutput, error) {
	ret := m.ctrl.Call(m, "AttachNetworkInterface", arg0)
	ret0, _ := ret[0].(*ec2.AttachNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachNetworkInterface indicates an expected call of AttachNetworkInterface
func (mr *MockEC2MockRecorder) AttachNetworkInterface(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachNetworkInterface", reflect.TypeOf((*MockEC2)(nil).AttachNetworkInterface), arg0)
}

// CreateNetworkInterface mocks base method
func (m *MockEC2) CreateNetworkInterface(arg0 *ec2.CreateNetworkInterfaceInput) (*ec2.CreateNetworkInterfaceOutput, error) {
	ret := m.ctrl.Call(m, "CreateNetworkInterface", arg0)
	ret0, _ := ret[0].(*ec2.CreateNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNetworkInterface indicates an expected call of CreateNetworkInterface
func (mr *MockEC2MockRecorder) CreateNetworkInterface(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetworkInterface", reflect.TypeOf((*MockEC2)(nil).CreateNetworkInterface), arg0)
}

// CreateTags mocks base method
func (m *MockEC2) CreateTags(arg0 *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	ret := m.ctrl.Call(m, "CreateTags", arg0)
	ret0, _ := ret[0].(*ec2.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTags indicates an expected call of CreateTags
func (mr *MockEC2MockRecorder) CreateTags(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockEC2)(nil).CreateTags), arg0)
}

// DeleteNetworkInterface mocks base method
func (m *MockEC2) DeleteNetworkInterface(arg0 *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
	ret := m.ctrl.Call(m, "DeleteNetworkInterface", arg0)
	ret0, _ := ret[0].(*ec2.DeleteNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNetworkInterface indicates an expected call of DeleteNetworkInterface
func (mr *MockEC2MockRecorder) DeleteNetworkInterface(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkInterface", reflect.TypeOf((*MockEC2)(nil).DeleteNetworkInterface), arg0)
}

// DescribeInstances mocks base method
func (m *MockEC2) DescribeInstances(arg0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	ret := m.ctrl.Call(m, "DescribeInstances", arg0)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances
func (mr *MockEC2MockRecorder) DescribeInstances(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockEC2)(nil).DescribeInstances), arg0)
}

// DescribeNetworkInterfaces mocks base method
func (m *MockEC2) DescribeNetworkInterfaces(arg0 *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	ret := m.ctrl.Call(m, "DescribeNetworkInterfaces", arg0)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNetworkInterfaces indicates an expected call of DescribeNetworkInterfaces
func (mr *MockEC2MockRecorder) DescribeNetworkInterfaces(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNetworkInterfaces", reflect.TypeOf((*MockEC2)(nil).DescribeNetworkInterfaces), arg0)
}

// DetachNetworkInterface mocks base method
func (m *MockEC2) DetachNetworkInterface(arg0 *ec2.DetachNetworkInterfaceInput) (*ec2.DetachNetworkInterfaceOutput, error) {
	ret := m.ctrl.Call(m, "DetachNetworkInterface", arg0)
	ret0, _ := ret[0].(*ec2.DetachNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachNetworkInterface indicates an expected call of DetachNetworkInterface
func (mr *MockEC2MockRecorder) DetachNetworkInterface(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachNetworkInterface", reflect.TypeOf((*MockEC2)(nil).DetachNetworkInterface), arg0)
}

// ModifyNetworkInterfaceAttribute mocks base method
func (m *MockEC2) ModifyNetworkInterfaceAttribute(arg0 *ec2.ModifyNetworkInterfaceAttributeInput) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	ret := m.ctrl.Call(m, "ModifyNetworkInterfaceAttribute", arg0)
	ret0, _ := ret[0].(*ec2.ModifyNetworkInterfaceAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyNetworkInterfaceAttribute indicates an expected call of ModifyNetworkInterfaceAttribute
func (mr *MockEC2MockRecorder) ModifyNetworkInterfaceAttribute(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyNetworkInterfaceAttribute", reflect.TypeOf((*MockEC2)(nil).ModifyNetworkInterfaceAttribute), arg0)
}
