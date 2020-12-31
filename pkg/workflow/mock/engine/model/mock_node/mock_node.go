// Copyright Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: ./engine/model/node/node.go

// Package mock_node is a generated GoMock package.
package mock_node

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	node "github.com/chaos-mesh/chaos-mesh/pkg/workflow/engine/model/node"
)

// MockNode is a mock of Node interface
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockNode) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockNodeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockNode)(nil).GetName))
}

// GetNodePhase mocks base method
func (m *MockNode) GetNodePhase() node.NodePhase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodePhase")
	ret0, _ := ret[0].(node.NodePhase)
	return ret0
}

// GetNodePhase indicates an expected call of GetNodePhase
func (mr *MockNodeMockRecorder) GetNodePhase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodePhase", reflect.TypeOf((*MockNode)(nil).GetNodePhase))
}

// GetParentNodeName mocks base method
func (m *MockNode) GetParentNodeName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParentNodeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetParentNodeName indicates an expected call of GetParentNodeName
func (mr *MockNodeMockRecorder) GetParentNodeName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentNodeName", reflect.TypeOf((*MockNode)(nil).GetParentNodeName))
}

// GetTemplateName mocks base method
func (m *MockNode) GetTemplateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTemplateName indicates an expected call of GetTemplateName
func (mr *MockNodeMockRecorder) GetTemplateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplateName", reflect.TypeOf((*MockNode)(nil).GetTemplateName))
}

// MockNodeTreeNode is a mock of NodeTreeNode interface
type MockNodeTreeNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeTreeNodeMockRecorder
}

// MockNodeTreeNodeMockRecorder is the mock recorder for MockNodeTreeNode
type MockNodeTreeNodeMockRecorder struct {
	mock *MockNodeTreeNode
}

// NewMockNodeTreeNode creates a new mock instance
func NewMockNodeTreeNode(ctrl *gomock.Controller) *MockNodeTreeNode {
	mock := &MockNodeTreeNode{ctrl: ctrl}
	mock.recorder = &MockNodeTreeNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeTreeNode) EXPECT() *MockNodeTreeNodeMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockNodeTreeNode) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockNodeTreeNodeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockNodeTreeNode)(nil).GetName))
}

// GetTemplateName mocks base method
func (m *MockNodeTreeNode) GetTemplateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTemplateName indicates an expected call of GetTemplateName
func (mr *MockNodeTreeNodeMockRecorder) GetTemplateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplateName", reflect.TypeOf((*MockNodeTreeNode)(nil).GetTemplateName))
}

// GetChildren mocks base method
func (m *MockNodeTreeNode) GetChildren() node.NodeTreeChildren {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildren")
	ret0, _ := ret[0].(node.NodeTreeChildren)
	return ret0
}

// GetChildren indicates an expected call of GetChildren
func (mr *MockNodeTreeNodeMockRecorder) GetChildren() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildren", reflect.TypeOf((*MockNodeTreeNode)(nil).GetChildren))
}

// FetchNodeByName mocks base method
func (m *MockNodeTreeNode) FetchNodeByName(nodeName string) node.NodeTreeNode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchNodeByName", nodeName)
	ret0, _ := ret[0].(node.NodeTreeNode)
	return ret0
}

// FetchNodeByName indicates an expected call of FetchNodeByName
func (mr *MockNodeTreeNodeMockRecorder) FetchNodeByName(nodeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchNodeByName", reflect.TypeOf((*MockNodeTreeNode)(nil).FetchNodeByName), nodeName)
}

// MockNodeTreeChildren is a mock of NodeTreeChildren interface
type MockNodeTreeChildren struct {
	ctrl     *gomock.Controller
	recorder *MockNodeTreeChildrenMockRecorder
}

// MockNodeTreeChildrenMockRecorder is the mock recorder for MockNodeTreeChildren
type MockNodeTreeChildrenMockRecorder struct {
	mock *MockNodeTreeChildren
}

// NewMockNodeTreeChildren creates a new mock instance
func NewMockNodeTreeChildren(ctrl *gomock.Controller) *MockNodeTreeChildren {
	mock := &MockNodeTreeChildren{ctrl: ctrl}
	mock.recorder = &MockNodeTreeChildrenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeTreeChildren) EXPECT() *MockNodeTreeChildrenMockRecorder {
	return m.recorder
}

// Length mocks base method
func (m *MockNodeTreeChildren) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockNodeTreeChildrenMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockNodeTreeChildren)(nil).Length))
}

// ContainsTemplate mocks base method
func (m *MockNodeTreeChildren) ContainsTemplate(templateName string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsTemplate", templateName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainsTemplate indicates an expected call of ContainsTemplate
func (mr *MockNodeTreeChildrenMockRecorder) ContainsTemplate(templateName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsTemplate", reflect.TypeOf((*MockNodeTreeChildren)(nil).ContainsTemplate), templateName)
}

// GetAllChildrenNode mocks base method
func (m *MockNodeTreeChildren) GetAllChildrenNode() []node.NodeTreeNode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChildrenNode")
	ret0, _ := ret[0].([]node.NodeTreeNode)
	return ret0
}

// GetAllChildrenNode indicates an expected call of GetAllChildrenNode
func (mr *MockNodeTreeChildrenMockRecorder) GetAllChildrenNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChildrenNode", reflect.TypeOf((*MockNodeTreeChildren)(nil).GetAllChildrenNode))
}
