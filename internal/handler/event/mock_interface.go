// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/event/interface.go
//
// Generated by this command:
//
//	mockgen -source internal/handler/event/interface.go -destination internal/handler/event/mock_interface.go -package event
//

// Package event is a generated GoMock package.
package event

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEventHandler) Create(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockEventHandlerMockRecorder) Create(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEventHandler)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockEventHandler) Delete(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEventHandlerMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventHandler)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockEventHandler) GetAll(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEventHandlerMockRecorder) GetAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEventHandler)(nil).GetAll), arg0)
}

// GetOne mocks base method.
func (m *MockEventHandler) GetOne(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetOne indicates an expected call of GetOne.
func (mr *MockEventHandlerMockRecorder) GetOne(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockEventHandler)(nil).GetOne), arg0)
}

// Update mocks base method.
func (m *MockEventHandler) Update(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEventHandlerMockRecorder) Update(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventHandler)(nil).Update), arg0)
}