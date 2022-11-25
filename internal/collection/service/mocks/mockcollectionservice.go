// Code generated by MockGen. DO NOT EDIT.
// Source: collectionservice.go

// Package mockCollectionService is a generated GoMock package.
package mockCollectionService

import (
	context "context"
	models "go-park-mail-ru/2022_2_BugOverload/internal/models"
	pkg "go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCollectionService is a mock of CollectionService interface.
type MockCollectionService struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionServiceMockRecorder
}

// MockCollectionServiceMockRecorder is the mock recorder for MockCollectionService.
type MockCollectionServiceMockRecorder struct {
	mock *MockCollectionService
}

// NewMockCollectionService creates a new mock instance.
func NewMockCollectionService(ctrl *gomock.Controller) *MockCollectionService {
	mock := &MockCollectionService{ctrl: ctrl}
	mock.recorder = &MockCollectionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollectionService) EXPECT() *MockCollectionServiceMockRecorder {
	return m.recorder
}

// GetCollectionByGenre mocks base method.
func (m *MockCollectionService) GetCollectionByGenre(ctx context.Context, params *pkg.GetStdCollectionParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionByGenre", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionByGenre indicates an expected call of GetCollectionByGenre.
func (mr *MockCollectionServiceMockRecorder) GetCollectionByGenre(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionByGenre", reflect.TypeOf((*MockCollectionService)(nil).GetCollectionByGenre), ctx, params)
}

// GetCollectionByTag mocks base method.
func (m *MockCollectionService) GetCollectionByTag(ctx context.Context, params *pkg.GetStdCollectionParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionByTag", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionByTag indicates an expected call of GetCollectionByTag.
func (mr *MockCollectionServiceMockRecorder) GetCollectionByTag(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionByTag", reflect.TypeOf((*MockCollectionService)(nil).GetCollectionByTag), ctx, params)
}

// GetPremieresCollection mocks base method.
func (m *MockCollectionService) GetPremieresCollection(ctx context.Context, params *pkg.GetStdCollectionParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPremieresCollection", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPremieresCollection indicates an expected call of GetPremieresCollection.
func (mr *MockCollectionServiceMockRecorder) GetPremieresCollection(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPremieresCollection", reflect.TypeOf((*MockCollectionService)(nil).GetPremieresCollection), ctx, params)
}

// GetStdCollection mocks base method.
func (m *MockCollectionService) GetStdCollection(ctx context.Context, params *pkg.GetStdCollectionParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStdCollection", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStdCollection indicates an expected call of GetStdCollection.
func (mr *MockCollectionServiceMockRecorder) GetStdCollection(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStdCollection", reflect.TypeOf((*MockCollectionService)(nil).GetStdCollection), ctx, params)
}

// GetUserCollections mocks base method.
func (m *MockCollectionService) GetUserCollections(ctx context.Context, user *models.User, params *pkg.GetUserCollectionsParams) ([]models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCollections", ctx, user, params)
	ret0, _ := ret[0].([]models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCollections indicates an expected call of GetUserCollections.
func (mr *MockCollectionServiceMockRecorder) GetUserCollections(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCollections", reflect.TypeOf((*MockCollectionService)(nil).GetUserCollections), ctx, user, params)
}
