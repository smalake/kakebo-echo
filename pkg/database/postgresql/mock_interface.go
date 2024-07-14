// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/postgresql/interface.go
//
// Generated by this command:
//
//	mockgen -source pkg/database/postgresql/interface.go -destination pkg/database/postgresql/mock_interface.go -package postgresql
//

// Package postgresql is a generated GoMock package.
package postgresql

import (
	reflect "reflect"

	sqlx "github.com/jmoiron/sqlx"
	gomock "go.uber.org/mock/gomock"
)

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// GetDB mocks base method.
func (m *MockClientInterface) GetDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// GetDB indicates an expected call of GetDB.
func (mr *MockClientInterfaceMockRecorder) GetDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockClientInterface)(nil).GetDB))
}