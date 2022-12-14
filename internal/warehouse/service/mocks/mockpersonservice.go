// Code generated by MockGen. DO NOT EDIT.
// Source: personservice.go

// Package mockWarehouseService is a generated GoMock package.
package mockWarehouseService

import (
	context "context"
	models "go-park-mail-ru/2022_2_BugOverload/internal/models"
	constparams "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPersonService is a mock of PersonService interface.
type MockPersonService struct {
	ctrl     *gomock.Controller
	recorder *MockPersonServiceMockRecorder
}

// MockPersonServiceMockRecorder is the mock recorder for MockPersonService.
type MockPersonServiceMockRecorder struct {
	mock *MockPersonService
}

// NewMockPersonService creates a new mock instance.
func NewMockPersonService(ctrl *gomock.Controller) *MockPersonService {
	mock := &MockPersonService{ctrl: ctrl}
	mock.recorder = &MockPersonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonService) EXPECT() *MockPersonServiceMockRecorder {
	return m.recorder
}

// GetPersonByID mocks base method.
func (m *MockPersonService) GetPersonByID(ctx context.Context, person *models.Person, params *constparams.GetPersonParams) (models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonByID", ctx, person, params)
	ret0, _ := ret[0].(models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPersonByID indicates an expected call of GetPersonByID.
func (mr *MockPersonServiceMockRecorder) GetPersonByID(ctx, person, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonByID", reflect.TypeOf((*MockPersonService)(nil).GetPersonByID), ctx, person, params)
}
