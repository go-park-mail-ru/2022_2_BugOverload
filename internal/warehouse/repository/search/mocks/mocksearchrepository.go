// Code generated by MockGen. DO NOT EDIT.
// Source: searchpgx.go

// Package mockSearchRepository is a generated GoMock package.
package mockSearchRepository

import (
	context "context"
	models "go-park-mail-ru/2022_2_BugOverload/internal/models"
	constparams "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// SearchFilms mocks base method.
func (m *MockRepository) SearchFilms(ctx context.Context, params *constparams.SearchParams) ([]models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFilms", ctx, params)
	ret0, _ := ret[0].([]models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFilms indicates an expected call of SearchFilms.
func (mr *MockRepositoryMockRecorder) SearchFilms(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFilms", reflect.TypeOf((*MockRepository)(nil).SearchFilms), ctx, params)
}

// SearchPersons mocks base method.
func (m *MockRepository) SearchPersons(ctx context.Context, params *constparams.SearchParams) ([]models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPersons", ctx, params)
	ret0, _ := ret[0].([]models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPersons indicates an expected call of SearchPersons.
func (mr *MockRepositoryMockRecorder) SearchPersons(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPersons", reflect.TypeOf((*MockRepository)(nil).SearchPersons), ctx, params)
}

// SearchSeries mocks base method.
func (m *MockRepository) SearchSeries(ctx context.Context, params *constparams.SearchParams) ([]models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchSeries", ctx, params)
	ret0, _ := ret[0].([]models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchSeries indicates an expected call of SearchSeries.
func (mr *MockRepositoryMockRecorder) SearchSeries(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchSeries", reflect.TypeOf((*MockRepository)(nil).SearchSeries), ctx, params)
}
