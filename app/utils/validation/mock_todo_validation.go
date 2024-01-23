// Code generated by MockGen. DO NOT EDIT.
// Source: todo_validation.go
//
// Generated by this command:
//
//	mockgen -package validation -source todo_validation.go -destination mock_todo_validation.go
//
// Package validation is a generated GoMock package.
package validation

import (
	reflect "reflect"

	models "github.com/tasuke/go-mux-task/models"
	gomock "go.uber.org/mock/gomock"
)

// MockITodoValidation is a mock of ITodoValidation interface.
type MockITodoValidation struct {
	ctrl     *gomock.Controller
	recorder *MockITodoValidationMockRecorder
}

// MockITodoValidationMockRecorder is the mock recorder for MockITodoValidation.
type MockITodoValidationMockRecorder struct {
	mock *MockITodoValidation
}

// NewMockITodoValidation creates a new mock instance.
func NewMockITodoValidation(ctrl *gomock.Controller) *MockITodoValidation {
	mock := &MockITodoValidation{ctrl: ctrl}
	mock.recorder = &MockITodoValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITodoValidation) EXPECT() *MockITodoValidationMockRecorder {
	return m.recorder
}

// MutationTodoValidate mocks base method.
func (m *MockITodoValidation) MutationTodoValidate(mutationTodoRequest models.MutationTodoRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutationTodoValidate", mutationTodoRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// MutationTodoValidate indicates an expected call of MutationTodoValidate.
func (mr *MockITodoValidationMockRecorder) MutationTodoValidate(mutationTodoRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutationTodoValidate", reflect.TypeOf((*MockITodoValidation)(nil).MutationTodoValidate), mutationTodoRequest)
}
