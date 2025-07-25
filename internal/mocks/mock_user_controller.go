// Code generated by MockGen. DO NOT EDIT.
// Source: internal/controller/user/controller.go
//
// Generated by this command:
//
//	mockgen --source=internal/controller/user/controller.go --destination=internal/mocks/mock_user_controller.go --package=mocks --mock_names=service=MockUserService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	user "github.com/nachoconques0/diagnosis_svc/internal/entity/user"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of service interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetByEmail mocks base method.
func (m *MockUserService) GetByEmail(ctx context.Context, email string) (*user.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*user.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserServiceMockRecorder) GetByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUserService)(nil).GetByEmail), ctx, email)
}
