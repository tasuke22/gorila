// Code generated by MockGen. DO NOT EDIT.
// Source: todo_service.go
//
// Generated by this command:
//
//	mockgen -package services -source todo_service.go -destination mock_todo_service.go
//
// Package services is a generated GoMock package.
package services

import (
	http "net/http"
	reflect "reflect"

	models "github.com/tasuke/go-mux-task/models"
	gomock "go.uber.org/mock/gomock"
)

// MockTodoService is a mock of TodoService interface.
type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

// MockTodoServiceMockRecorder is the mock recorder for MockTodoService.
type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

// NewMockTodoService creates a new mock instance.
func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockTodoService) CreateTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", w, r, userId)
	ret0, _ := ret[0].(models.BaseTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockTodoServiceMockRecorder) CreateTodo(w, r, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockTodoService)(nil).CreateTodo), w, r, userId)
}

// DeleteTodo mocks base method.
func (m *MockTodoService) DeleteTodo(w http.ResponseWriter, r *http.Request, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", w, r, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockTodoServiceMockRecorder) DeleteTodo(w, r, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodoService)(nil).DeleteTodo), w, r, userId)
}

// GetAllTodos mocks base method.
func (m *MockTodoService) GetAllTodos(w http.ResponseWriter, userId int) ([]models.BaseTodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTodos", w, userId)
	ret0, _ := ret[0].([]models.BaseTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTodos indicates an expected call of GetAllTodos.
func (mr *MockTodoServiceMockRecorder) GetAllTodos(w, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodos", reflect.TypeOf((*MockTodoService)(nil).GetAllTodos), w, userId)
}

// GetTodo mocks base method.
func (m *MockTodoService) GetTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodo", w, r, userId)
	ret0, _ := ret[0].(models.BaseTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodo indicates an expected call of GetTodo.
func (mr *MockTodoServiceMockRecorder) GetTodo(w, r, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodo", reflect.TypeOf((*MockTodoService)(nil).GetTodo), w, r, userId)
}

// SendAllTodoResponse mocks base method.
func (m *MockTodoService) SendAllTodoResponse(w http.ResponseWriter, todos *[]models.BaseTodoResponse) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAllTodoResponse", w, todos)
}

// SendAllTodoResponse indicates an expected call of SendAllTodoResponse.
func (mr *MockTodoServiceMockRecorder) SendAllTodoResponse(w, todos any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAllTodoResponse", reflect.TypeOf((*MockTodoService)(nil).SendAllTodoResponse), w, todos)
}

// SendCreateTodoResponse mocks base method.
func (m *MockTodoService) SendCreateTodoResponse(w http.ResponseWriter, todo *models.BaseTodoResponse) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendCreateTodoResponse", w, todo)
}

// SendCreateTodoResponse indicates an expected call of SendCreateTodoResponse.
func (mr *MockTodoServiceMockRecorder) SendCreateTodoResponse(w, todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCreateTodoResponse", reflect.TypeOf((*MockTodoService)(nil).SendCreateTodoResponse), w, todo)
}

// SendDeleteTodoResponse mocks base method.
func (m *MockTodoService) SendDeleteTodoResponse(w http.ResponseWriter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendDeleteTodoResponse", w)
}

// SendDeleteTodoResponse indicates an expected call of SendDeleteTodoResponse.
func (mr *MockTodoServiceMockRecorder) SendDeleteTodoResponse(w any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDeleteTodoResponse", reflect.TypeOf((*MockTodoService)(nil).SendDeleteTodoResponse), w)
}

// SendTodoResponse mocks base method.
func (m *MockTodoService) SendTodoResponse(w http.ResponseWriter, todo *models.BaseTodoResponse) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendTodoResponse", w, todo)
}

// SendTodoResponse indicates an expected call of SendTodoResponse.
func (mr *MockTodoServiceMockRecorder) SendTodoResponse(w, todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTodoResponse", reflect.TypeOf((*MockTodoService)(nil).SendTodoResponse), w, todo)
}

// UpdateTodo mocks base method.
func (m *MockTodoService) UpdateTodo(w http.ResponseWriter, r *http.Request, userId int) (models.BaseTodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", w, r, userId)
	ret0, _ := ret[0].(models.BaseTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockTodoServiceMockRecorder) UpdateTodo(w, r, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodoService)(nil).UpdateTodo), w, r, userId)
}