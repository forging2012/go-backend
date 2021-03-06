// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MongoDBNavigator/go-backend/domain/database/repository (interfaces: IndexWriter)

// Package mock is a generated GoMock package.
package mock

import (
	model "github.com/MongoDBNavigator/go-backend/domain/database/model"
	value "github.com/MongoDBNavigator/go-backend/domain/database/value"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIndexWriter is a mock of IndexWriter interface
type MockIndexWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIndexWriterMockRecorder
}

// MockIndexWriterMockRecorder is the mock recorder for MockIndexWriter
type MockIndexWriterMockRecorder struct {
	mock *MockIndexWriter
}

// NewMockIndexWriter creates a new mock instance
func NewMockIndexWriter(ctrl *gomock.Controller) *MockIndexWriter {
	mock := &MockIndexWriter{ctrl: ctrl}
	mock.recorder = &MockIndexWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexWriter) EXPECT() *MockIndexWriterMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIndexWriter) Create(arg0 value.DBName, arg1 value.CollName, arg2 *model.Index) error {
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIndexWriterMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIndexWriter)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method
func (m *MockIndexWriter) Delete(arg0 value.DBName, arg1 value.CollName, arg2 value.IndexName) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIndexWriterMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIndexWriter)(nil).Delete), arg0, arg1, arg2)
}
