// Code generated by MockGen. DO NOT EDIT.
// Source: ./Letter.go

// Package mock_LetterUseCase is a generated GoMock package.
package mock_LetterUseCase

import (
	LetterModel "Mailer/MainApplication/internal/Letter/LetterModel"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLetterUseCase is a mock of LetterUseCase interface
type MockLetterUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockLetterUseCaseMockRecorder
}

// MockLetterUseCaseMockRecorder is the mock recorder for MockLetterUseCase
type MockLetterUseCaseMockRecorder struct {
	mock *MockLetterUseCase
}

// NewMockLetterUseCase creates a new mock instance
func NewMockLetterUseCase(ctrl *gomock.Controller) *MockLetterUseCase {
	mock := &MockLetterUseCase{ctrl: ctrl}
	mock.recorder = &MockLetterUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLetterUseCase) EXPECT() *MockLetterUseCaseMockRecorder {
	return m.recorder
}

// SaveLetter mocks base method
func (m *MockLetterUseCase) SaveLetter(letter *LetterModel.Letter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLetter", letter)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLetter indicates an expected call of SaveLetter
func (mr *MockLetterUseCaseMockRecorder) SaveLetter(letter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLetter", reflect.TypeOf((*MockLetterUseCase)(nil).SaveLetter), letter)
}

// GetReceivedLetters mocks base method
func (m *MockLetterUseCase) GetReceivedLetters(email string, limit, offset uint64) (error, []LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceivedLetters", email, limit, offset)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]LetterModel.Letter)
	return ret0, ret1
}

// GetReceivedLetters indicates an expected call of GetReceivedLetters
func (mr *MockLetterUseCaseMockRecorder) GetReceivedLetters(email, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceivedLetters", reflect.TypeOf((*MockLetterUseCase)(nil).GetReceivedLetters), email, limit, offset)
}

// GetSendedLetters mocks base method
func (m *MockLetterUseCase) GetSendedLetters(email string) (error, []LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSendedLetters", email)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]LetterModel.Letter)
	return ret0, ret1
}

// GetSendedLetters indicates an expected call of GetSendedLetters
func (mr *MockLetterUseCaseMockRecorder) GetSendedLetters(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSendedLetters", reflect.TypeOf((*MockLetterUseCase)(nil).GetSendedLetters), email)
}

// GetReceivedLettersDir mocks base method
func (m *MockLetterUseCase) GetReceivedLettersDir(dir uint64) (error, []LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceivedLettersDir", dir)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]LetterModel.Letter)
	return ret0, ret1
}

// GetReceivedLettersDir indicates an expected call of GetReceivedLettersDir
func (mr *MockLetterUseCaseMockRecorder) GetReceivedLettersDir(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceivedLettersDir", reflect.TypeOf((*MockLetterUseCase)(nil).GetReceivedLettersDir), dir)
}

// GetSendedLettersDir mocks base method
func (m *MockLetterUseCase) GetSendedLettersDir(dir uint64) (error, []LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSendedLettersDir", dir)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]LetterModel.Letter)
	return ret0, ret1
}

// GetSendedLettersDir indicates an expected call of GetSendedLettersDir
func (mr *MockLetterUseCaseMockRecorder) GetSendedLettersDir(dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSendedLettersDir", reflect.TypeOf((*MockLetterUseCase)(nil).GetSendedLettersDir), dir)
}

// WatchLetter mocks base method
func (m *MockLetterUseCase) WatchLetter(lid uint64) (error, LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchLetter", lid)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(LetterModel.Letter)
	return ret0, ret1
}

// WatchLetter indicates an expected call of WatchLetter
func (mr *MockLetterUseCaseMockRecorder) WatchLetter(lid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLetter", reflect.TypeOf((*MockLetterUseCase)(nil).WatchLetter), lid)
}

// DeleteLetter mocks base method
func (m *MockLetterUseCase) DeleteLetter(lid uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLetter", lid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLetter indicates an expected call of DeleteLetter
func (mr *MockLetterUseCaseMockRecorder) DeleteLetter(lid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLetter", reflect.TypeOf((*MockLetterUseCase)(nil).DeleteLetter), lid)
}

// FindSim mocks base method
func (m *MockLetterUseCase) FindSim(sim string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSim", sim)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindSim indicates an expected call of FindSim
func (mr *MockLetterUseCaseMockRecorder) FindSim(sim interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSim", reflect.TypeOf((*MockLetterUseCase)(nil).FindSim), sim)
}

// GetLetterBy mocks base method
func (m *MockLetterUseCase) GetLetterBy(what, val string) (error, []LetterModel.Letter) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLetterBy", what, val)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]LetterModel.Letter)
	return ret0, ret1
}

// GetLetterBy indicates an expected call of GetLetterBy
func (mr *MockLetterUseCaseMockRecorder) GetLetterBy(what, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLetterBy", reflect.TypeOf((*MockLetterUseCase)(nil).GetLetterBy), what, val)
}
