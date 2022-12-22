// Code generated by MockGen. DO NOT EDIT.
// Source: collectionpgx.go

// Package mockCollectionRepository is a generated GoMock package.
package mockCollectionRepository

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

// CheckCollectionIsPublic mocks base method.
func (m *MockRepository) CheckCollectionIsPublic(ctx context.Context, params *constparams.CollectionGetFilmsRequestParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCollectionIsPublic", ctx, params)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCollectionIsPublic indicates an expected call of CheckCollectionIsPublic.
func (mr *MockRepositoryMockRecorder) CheckCollectionIsPublic(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCollectionIsPublic", reflect.TypeOf((*MockRepository)(nil).CheckCollectionIsPublic), ctx, params)
}

// CheckUserIsAuthor mocks base method.
func (m *MockRepository) CheckUserIsAuthor(ctx context.Context, user *models.User, params *constparams.CollectionGetFilmsRequestParams) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserIsAuthor", ctx, user, params)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserIsAuthor indicates an expected call of CheckUserIsAuthor.
func (mr *MockRepositoryMockRecorder) CheckUserIsAuthor(ctx, user, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserIsAuthor", reflect.TypeOf((*MockRepository)(nil).CheckUserIsAuthor), ctx, user, params)
}

// GetCollection mocks base method.
func (m *MockRepository) GetCollection(ctx context.Context, params *constparams.CollectionGetFilmsRequestParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockRepositoryMockRecorder) GetCollection(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockRepository)(nil).GetCollection), ctx, params)
}

// GetCollectionAuthor mocks base method.
func (m *MockRepository) GetCollectionAuthor(ctx context.Context, params *constparams.CollectionGetFilmsRequestParams) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionAuthor", ctx, params)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionAuthor indicates an expected call of GetCollectionAuthor.
func (mr *MockRepositoryMockRecorder) GetCollectionAuthor(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionAuthor", reflect.TypeOf((*MockRepository)(nil).GetCollectionAuthor), ctx, params)
}

// GetCollectionByGenre mocks base method.
func (m *MockRepository) GetCollectionByGenre(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionByGenre", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionByGenre indicates an expected call of GetCollectionByGenre.
func (mr *MockRepositoryMockRecorder) GetCollectionByGenre(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionByGenre", reflect.TypeOf((*MockRepository)(nil).GetCollectionByGenre), ctx, params)
}

// GetCollectionByTag mocks base method.
func (m *MockRepository) GetCollectionByTag(ctx context.Context, params *constparams.GetStdCollectionParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionByTag", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionByTag indicates an expected call of GetCollectionByTag.
func (mr *MockRepositoryMockRecorder) GetCollectionByTag(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionByTag", reflect.TypeOf((*MockRepository)(nil).GetCollectionByTag), ctx, params)
}

// GetPremieresCollection mocks base method.
func (m *MockRepository) GetPremieresCollection(ctx context.Context, params *constparams.GetPremiersCollectionParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPremieresCollection", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPremieresCollection indicates an expected call of GetPremieresCollection.
func (mr *MockRepositoryMockRecorder) GetPremieresCollection(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPremieresCollection", reflect.TypeOf((*MockRepository)(nil).GetPremieresCollection), ctx, params)
}

// GetSimilarFilms mocks base method.
func (m *MockRepository) GetSimilarFilms(ctx context.Context, params *constparams.GetSimilarFilmsParams) (models.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSimilarFilms", ctx, params)
	ret0, _ := ret[0].(models.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSimilarFilms indicates an expected call of GetSimilarFilms.
func (mr *MockRepositoryMockRecorder) GetSimilarFilms(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimilarFilms", reflect.TypeOf((*MockRepository)(nil).GetSimilarFilms), ctx, params)
}
