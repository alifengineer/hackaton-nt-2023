// Code generated by MockGen. DO NOT EDIT.
// Source: storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	models "github.com/dilmurodov/xakaton_nt/pkg/models"
	storage "github.com/dilmurodov/xakaton_nt/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockStorageI is a mock of StorageI interface.
type MockStorageI struct {
	ctrl     *gomock.Controller
	recorder *MockStorageIMockRecorder
}

// MockStorageIMockRecorder is the mock recorder for MockStorageI.
type MockStorageIMockRecorder struct {
	mock *MockStorageI
}

// NewMockStorageI creates a new mock instance.
func NewMockStorageI(ctrl *gomock.Controller) *MockStorageI {
	mock := &MockStorageI{ctrl: ctrl}
	mock.recorder = &MockStorageIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageI) EXPECT() *MockStorageIMockRecorder {
	return m.recorder
}

// CloseDB mocks base method.
func (m *MockStorageI) CloseDB() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseDB")
}

// CloseDB indicates an expected call of CloseDB.
func (mr *MockStorageIMockRecorder) CloseDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseDB", reflect.TypeOf((*MockStorageI)(nil).CloseDB))
}

// User mocks base method.
func (m *MockStorageI) User() storage.UserRepoI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(storage.UserRepoI)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockStorageIMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockStorageI)(nil).User))
}

// MockUserRepoI is a mock of UserRepoI interface.
type MockUserRepoI struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoIMockRecorder
}

// MockUserRepoIMockRecorder is the mock recorder for MockUserRepoI.
type MockUserRepoIMockRecorder struct {
	mock *MockUserRepoI
}

// NewMockUserRepoI creates a new mock instance.
func NewMockUserRepoI(ctrl *gomock.Controller) *MockUserRepoI {
	mock := &MockUserRepoI{ctrl: ctrl}
	mock.recorder = &MockUserRepoIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepoI) EXPECT() *MockUserRepoIMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepoI) CreateUser(arg0 context.Context, arg1 *models.CreateUserRequest) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoIMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepoI)(nil).CreateUser), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockUserRepoI) GetUserByID(arg0 context.Context, arg1 *models.GetUserByIDRequest) (*models.GetUserByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(*models.GetUserByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepoIMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepoI)(nil).GetUserByID), arg0, arg1)
}

// GetUserPasswordByPhone mocks base method.
func (m *MockUserRepoI) GetUserPasswordByPhone(ctx context.Context, phone string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPasswordByPhone", ctx, phone)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPasswordByPhone indicates an expected call of GetUserPasswordByPhone.
func (mr *MockUserRepoIMockRecorder) GetUserPasswordByPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPasswordByPhone", reflect.TypeOf((*MockUserRepoI)(nil).GetUserPasswordByPhone), ctx, phone)
}