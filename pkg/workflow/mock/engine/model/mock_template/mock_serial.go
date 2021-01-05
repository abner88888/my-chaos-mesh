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
// Source: ./engine/model/template/serial.go

// Package mock_template is a generated GoMock package.
package mock_template

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	template "github.com/chaos-mesh/chaos-mesh/pkg/workflow/engine/model/template"
)

// MockSerialTemplate is a mock of SerialTemplate interface
type MockSerialTemplate struct {
	ctrl     *gomock.Controller
	recorder *MockSerialTemplateMockRecorder
}

// MockSerialTemplateMockRecorder is the mock recorder for MockSerialTemplate
type MockSerialTemplateMockRecorder struct {
	mock *MockSerialTemplate
}

// NewMockSerialTemplate creates a new mock instance
func NewMockSerialTemplate(ctrl *gomock.Controller) *MockSerialTemplate {
	mock := &MockSerialTemplate{ctrl: ctrl}
	mock.recorder = &MockSerialTemplateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSerialTemplate) EXPECT() *MockSerialTemplateMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockSerialTemplate) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockSerialTemplateMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockSerialTemplate)(nil).GetName))
}

// GetTemplateType mocks base method
func (m *MockSerialTemplate) GetTemplateType() template.TemplateType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplateType")
	ret0, _ := ret[0].(template.TemplateType)
	return ret0
}

// GetTemplateType indicates an expected call of GetTemplateType
func (mr *MockSerialTemplateMockRecorder) GetTemplateType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplateType", reflect.TypeOf((*MockSerialTemplate)(nil).GetTemplateType))
}

// GetSerialChildrenList mocks base method
func (m *MockSerialTemplate) GetSerialChildrenList() []template.Template {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSerialChildrenList")
	ret0, _ := ret[0].([]template.Template)
	return ret0
}

// GetSerialChildrenList indicates an expected call of GetSerialChildrenList
func (mr *MockSerialTemplateMockRecorder) GetSerialChildrenList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSerialChildrenList", reflect.TypeOf((*MockSerialTemplate)(nil).GetSerialChildrenList))
}
