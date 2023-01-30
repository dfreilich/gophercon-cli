// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dfreilich/gophercon-cli/pkg/complex (interfaces: Actor)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	complex "github.com/dfreilich/gophercon-cli/pkg/complex"
	gomock "github.com/golang/mock/gomock"
)

// MockActor is a mock of Actor interface.
type MockActor struct {
	ctrl     *gomock.Controller
	recorder *MockActorMockRecorder
}

// MockActorMockRecorder is the mock recorder for MockActor.
type MockActorMockRecorder struct {
	mock *MockActor
}

// NewMockActor creates a new mock instance.
func NewMockActor(ctrl *gomock.Controller) *MockActor {
	mock := &MockActor{ctrl: ctrl}
	mock.recorder = &MockActorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActor) EXPECT() *MockActorMockRecorder {
	return m.recorder
}

// DoSomething mocks base method.
func (m *MockActor) DoSomething(arg0 *complex.DoSomethingArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSomething", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoSomething indicates an expected call of DoSomething.
func (mr *MockActorMockRecorder) DoSomething(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSomething", reflect.TypeOf((*MockActor)(nil).DoSomething), arg0)
}
