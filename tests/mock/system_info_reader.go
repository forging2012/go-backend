// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MongoDBNavigator/go-backend/domain/system/repository (interfaces: SystemInfoReader)

// Package mock is a generated GoMock package.
package mock

import (
	model "github.com/MongoDBNavigator/go-backend/domain/system/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSystemInfoReader is a mock of SystemInfoReader interface
type MockSystemInfoReader struct {
	ctrl     *gomock.Controller
	recorder *MockSystemInfoReaderMockRecorder
}

// MockSystemInfoReaderMockRecorder is the mock recorder for MockSystemInfoReader
type MockSystemInfoReaderMockRecorder struct {
	mock *MockSystemInfoReader
}

// NewMockSystemInfoReader creates a new mock instance
func NewMockSystemInfoReader(ctrl *gomock.Controller) *MockSystemInfoReader {
	mock := &MockSystemInfoReader{ctrl: ctrl}
	mock.recorder = &MockSystemInfoReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSystemInfoReader) EXPECT() *MockSystemInfoReaderMockRecorder {
	return m.recorder
}

// Reade mocks base method
func (m *MockSystemInfoReader) Reade() (*model.SystemInfo, error) {
	ret := m.ctrl.Call(m, "Reade")
	ret0, _ := ret[0].(*model.SystemInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reade indicates an expected call of Reade
func (mr *MockSystemInfoReaderMockRecorder) Reade() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reade", reflect.TypeOf((*MockSystemInfoReader)(nil).Reade))
}
