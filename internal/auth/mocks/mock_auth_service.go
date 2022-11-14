// Code generated by MockGen. DO NOT EDIT.
// Source: go-park-mail-ru/2022_2_BugOverload/internal/auth/service (interfaces: AuthService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "go-park-mail-ru/2022_2_BugOverload/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockAuthService) Auth(arg0 context.Context, arg1 *models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth.
func (mr *MockAuthServiceMockRecorder) Auth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockAuthService)(nil).Auth), arg0, arg1)
}

// GetAccess mocks base method.
func (m *MockAuthService) GetAccess(arg0 context.Context, arg1 *models.User, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAccess indicates an expected call of GetAccess.
func (mr *MockAuthServiceMockRecorder) GetAccess(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccess", reflect.TypeOf((*MockAuthService)(nil).GetAccess), arg0, arg1, arg2)
}

// Login mocks base method.
func (m *MockAuthService) Login(arg0 context.Context, arg1 *models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), arg0, arg1)
}

// Signup mocks base method.
func (m *MockAuthService) Signup(arg0 context.Context, arg1 *models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", arg0, arg1)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthServiceMockRecorder) Signup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuthService)(nil).Signup), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockAuthService) UpdatePassword(arg0 context.Context, arg1 *models.User, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockAuthServiceMockRecorder) UpdatePassword(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockAuthService)(nil).UpdatePassword), arg0, arg1, arg2)
}
