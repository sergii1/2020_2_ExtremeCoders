package tests

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockItemRepository is a mock of ItemRepository interface
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockItemRepository) GetAll() ([]*items.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*items.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockItemRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockItemRepository)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockItemRepository) GetByID(arg0 int64) (*items.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*items.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockItemRepositoryMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockItemRepository)(nil).GetByID), arg0)
}

// Add mocks base method
func (m *MockItemRepository) Add(arg0 *items.Item) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockItemRepositoryMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockItemRepository)(nil).Add), arg0)
}

// Update mocks base method
func (m *MockItemRepository) Update(arg0 *items.Item) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockItemRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemRepository)(nil).Update), arg0)
}

// Delete mocks base method
func (m *MockItemRepository) Delete(arg0 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockItemRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemRepository)(nil).Delete), arg0)
}