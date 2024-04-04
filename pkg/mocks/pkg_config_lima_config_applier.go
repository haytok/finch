// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch/pkg/config (interfaces: LimaConfigApplier)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	flog "github.com/runfinch/finch/pkg/flog"
	afero "github.com/spf13/afero"
)

// LimaConfigApplier is a mock of LimaConfigApplier interface.
type LimaConfigApplier struct {
	ctrl     *gomock.Controller
	recorder *LimaConfigApplierMockRecorder
}

// LimaConfigApplierMockRecorder is the mock recorder for LimaConfigApplier.
type LimaConfigApplierMockRecorder struct {
	mock *LimaConfigApplier
}

// NewLimaConfigApplier creates a new mock instance.
func NewLimaConfigApplier(ctrl *gomock.Controller) *LimaConfigApplier {
	mock := &LimaConfigApplier{ctrl: ctrl}
	mock.recorder = &LimaConfigApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *LimaConfigApplier) EXPECT() *LimaConfigApplierMockRecorder {
	return m.recorder
}

// ConfigureDefaultLimaYaml mocks base method.
func (m *LimaConfigApplier) ConfigureDefaultLimaYaml() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureDefaultLimaYaml")
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureDefaultLimaYaml indicates an expected call of ConfigureDefaultLimaYaml.
func (mr *LimaConfigApplierMockRecorder) ConfigureDefaultLimaYaml() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureDefaultLimaYaml", reflect.TypeOf((*LimaConfigApplier)(nil).ConfigureDefaultLimaYaml))
}

// ConfigureOverrideLimaYaml mocks base method.
func (m *LimaConfigApplier) ConfigureOverrideLimaYaml() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureOverrideLimaYaml")
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureOverrideLimaYaml indicates an expected call of ConfigureOverrideLimaYaml.
func (mr *LimaConfigApplierMockRecorder) ConfigureOverrideLimaYaml() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureOverrideLimaYaml", reflect.TypeOf((*LimaConfigApplier)(nil).ConfigureOverrideLimaYaml))
}

// ModifyFinchConfig mocks base method.
func (m *LimaConfigApplier) ModifyFinchConfig(arg0 afero.Fs, arg1 flog.Logger, arg2 int, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyFinchConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyFinchConfig indicates an expected call of ModifyFinchConfig.
func (mr *LimaConfigApplierMockRecorder) ModifyFinchConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyFinchConfig", reflect.TypeOf((*LimaConfigApplier)(nil).ModifyFinchConfig), arg0, arg1, arg2, arg3)
}
