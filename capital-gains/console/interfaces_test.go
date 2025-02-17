// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source=interfaces.go -destination=interfaces_test.go -package=console
//

// Package console is a generated GoMock package.
package console

import (
	reflect "reflect"

	logic "github.nubank.com/capital-gains/logic"
	gomock "go.uber.org/mock/gomock"
)

// Mockprocessor is a mock of processor interface.
type Mockprocessor struct {
	ctrl     *gomock.Controller
	recorder *MockprocessorMockRecorder
}

// MockprocessorMockRecorder is the mock recorder for Mockprocessor.
type MockprocessorMockRecorder struct {
	mock *Mockprocessor
}

// NewMockprocessor creates a new mock instance.
func NewMockprocessor(ctrl *gomock.Controller) *Mockprocessor {
	mock := &Mockprocessor{ctrl: ctrl}
	mock.recorder = &MockprocessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockprocessor) EXPECT() *MockprocessorMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *Mockprocessor) Process(operations logic.Operations) ([]logic.TaxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", operations)
	ret0, _ := ret[0].([]logic.TaxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Process indicates an expected call of Process.
func (mr *MockprocessorMockRecorder) Process(operations any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*Mockprocessor)(nil).Process), operations)
}
