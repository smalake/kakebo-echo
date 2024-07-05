// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/auth/interface.go
//
// Generated by this command:
//
//	mockgen -source internal/repository/auth/interface.go -destination internal/repository/auth/mock_interface.go -package auth
//

// Package auth is a generated GoMock package.
package auth

import (
	model "kakebo-echo/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// FindUser mocks base method.
func (m *MockAuthRepository) FindUser(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockAuthRepositoryMockRecorder) FindUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockAuthRepository)(nil).FindUser), arg0)
}

// LoginCheck mocks base method.
func (m *MockAuthRepository) LoginCheck(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginCheck", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginCheck indicates an expected call of LoginCheck.
func (mr *MockAuthRepositoryMockRecorder) LoginCheck(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginCheck", reflect.TypeOf((*MockAuthRepository)(nil).LoginCheck), arg0)
}

// Logout mocks base method.
func (m *MockAuthRepository) Logout() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthRepositoryMockRecorder) Logout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthRepository)(nil).Logout))
}

// Register mocks base method.
func (m *MockAuthRepository) Register(arg0 *model.RegisterRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockAuthRepositoryMockRecorder) Register(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuthRepository)(nil).Register), arg0)
}
