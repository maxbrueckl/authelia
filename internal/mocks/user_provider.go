// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/authentication (interfaces: UserProvider)
//
// Generated by this command:
//
//	mockgen -package mocks -destination user_provider.go -mock_names UserProvider=MockUserProvider github.com/authelia/authelia/v4/internal/authentication UserProvider
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	authentication "github.com/authelia/authelia/v4/internal/authentication"
	gomock "go.uber.org/mock/gomock"
)

// MockUserProvider is a mock of UserProvider interface.
type MockUserProvider struct {
	ctrl     *gomock.Controller
	recorder *MockUserProviderMockRecorder
	isgomock struct{}
}

// MockUserProviderMockRecorder is the mock recorder for MockUserProvider.
type MockUserProviderMockRecorder struct {
	mock *MockUserProvider
}

// NewMockUserProvider creates a new mock instance.
func NewMockUserProvider(ctrl *gomock.Controller) *MockUserProvider {
	mock := &MockUserProvider{ctrl: ctrl}
	mock.recorder = &MockUserProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserProvider) EXPECT() *MockUserProviderMockRecorder {
	return m.recorder
}

// CheckUserPassword mocks base method.
func (m *MockUserProvider) CheckUserPassword(username, password string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserPassword", username, password)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserPassword indicates an expected call of CheckUserPassword.
func (mr *MockUserProviderMockRecorder) CheckUserPassword(username, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserPassword", reflect.TypeOf((*MockUserProvider)(nil).CheckUserPassword), username, password)
}

// GetDetails mocks base method.
func (m *MockUserProvider) GetDetails(username string) (*authentication.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetails", username)
	ret0, _ := ret[0].(*authentication.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetails indicates an expected call of GetDetails.
func (mr *MockUserProviderMockRecorder) GetDetails(username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetails", reflect.TypeOf((*MockUserProvider)(nil).GetDetails), username)
}

// Shutdown mocks base method.
func (m *MockUserProvider) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockUserProviderMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockUserProvider)(nil).Shutdown))
}

// StartupCheck mocks base method.
func (m *MockUserProvider) StartupCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartupCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartupCheck indicates an expected call of StartupCheck.
func (mr *MockUserProviderMockRecorder) StartupCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartupCheck", reflect.TypeOf((*MockUserProvider)(nil).StartupCheck))
}

// UpdatePassword mocks base method.
func (m *MockUserProvider) UpdatePassword(username, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", username, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserProviderMockRecorder) UpdatePassword(username, newPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserProvider)(nil).UpdatePassword), username, newPassword)
}
