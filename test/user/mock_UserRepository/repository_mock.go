// Code generated by MockGen. DO NOT EDIT.
// Source: DataBaseRequests.go

// Package mock_UserRepository is mock_UserRepository generated GoMock package.
package mock_UserRepository

import (
	UserModel "CleanArch/internal/User/UserModel"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserDB is mock_UserRepository mock of UserDB interface
type MockUserDB struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBMockRecorder
}

// MockUserDBMockRecorder is the mock recorder for MockUserDB
type MockUserDBMockRecorder struct {
	mock *MockUserDB
}

// NewMockUserDB creates mock_UserRepository new mock instance
func NewMockUserDB(ctrl *gomock.Controller) *MockUserDB {
	mock := &MockUserDB{ctrl: ctrl}
	mock.recorder = &MockUserDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDB) EXPECT() *MockUserDBMockRecorder {
	return m.recorder
}

// IsEmailExists mocks base method
func (m *MockUserDB) IsEmailExists(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmailExists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsEmailExists indicates an expected call of IsEmailExists
func (mr *MockUserDBMockRecorder) IsEmailExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmailExists", reflect.TypeOf((*MockUserDB)(nil).IsEmailExists), arg0)
}

// AddSession mocks base method
func (m *MockUserDB) AddSession(arg0 string, arg1 uint64, arg2 *UserModel.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSession indicates an expected call of AddSession
func (mr *MockUserDBMockRecorder) AddSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSession", reflect.TypeOf((*MockUserDB)(nil).AddSession), arg0, arg1, arg2)
}

// AddUser mocks base method
func (m *MockUserDB) AddUser(arg0 *UserModel.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockUserDBMockRecorder) AddUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserDB)(nil).AddUser), arg0)
}

// GenerateSID mocks base method
func (m *MockUserDB) GenerateSID() ([]rune, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSID")
	ret0, _ := ret[0].([]rune)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateSID indicates an expected call of GenerateSID
func (mr *MockUserDBMockRecorder) GenerateSID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSID", reflect.TypeOf((*MockUserDB)(nil).GenerateSID))
}

// GenerateUID mocks base method
func (m *MockUserDB) GenerateUID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateUID indicates an expected call of GenerateUID
func (mr *MockUserDBMockRecorder) GenerateUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUID", reflect.TypeOf((*MockUserDB)(nil).GenerateUID))
}

// GetUserByEmail mocks base method
func (m *MockUserDB) GetUserByEmail(arg0 string) (*UserModel.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*UserModel.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockUserDBMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserDB)(nil).GetUserByEmail), arg0)
}

// GetUserByUID mocks base method
func (m *MockUserDB) GetUserByUID(arg0 uint64) (*UserModel.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUID", arg0)
	ret0, _ := ret[0].(*UserModel.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUID indicates an expected call of GetUserByUID
func (mr *MockUserDBMockRecorder) GetUserByUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUID", reflect.TypeOf((*MockUserDB)(nil).GetUserByUID), arg0)
}

// IsOkSession mocks base method
func (m *MockUserDB) IsOkSession(arg0 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOkSession", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOkSession indicates an expected call of IsOkSession
func (mr *MockUserDBMockRecorder) IsOkSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOkSession", reflect.TypeOf((*MockUserDB)(nil).IsOkSession), arg0)
}

// UpdateProfile mocks base method
func (m *MockUserDB) UpdateProfile(arg0 UserModel.User, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockUserDBMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserDB)(nil).UpdateProfile), arg0, arg1)
}

// RemoveSession mocks base method
func (m *MockUserDB) RemoveSession(arg0 string) (error, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSession", arg0)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// RemoveSession indicates an expected call of RemoveSession
func (mr *MockUserDBMockRecorder) RemoveSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSession", reflect.TypeOf((*MockUserDB)(nil).RemoveSession), arg0)
}

// GetSessionByUID mocks base method
func (m *MockUserDB) GetSessionByUID(arg0 uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByUID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByUID indicates an expected call of GetSessionByUID
func (mr *MockUserDBMockRecorder) GetSessionByUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByUID", reflect.TypeOf((*MockUserDB)(nil).GetSessionByUID), arg0)
}
