package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
)

func TestUserGetUserCollectionsHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/collections?sort_param=create_time&count_collections=15&delimiter=now", nil)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	var userCollections []modelsGlobal.Collection

	userCollections = append(userCollections, modelsGlobal.Collection{
		ID:         12,
		Name:       "Избранное",
		Poster:     "42",
		CountLikes: 1023,
		CountFilms: 10,
		UpdateTime: "2020.12.12 15:15:15",
		CreateTime: "2012.06.05 01:25:00",
	})

	// Settings mock
	rateService.EXPECT().GetUserCollections(r.Context(), &user, &constparams.GetUserCollectionsParams{
		SortParam:        "create_time",
		CountCollections: 15,
		Delimiter:        "now",
	}).Return(userCollections, nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetUserCollectionsHandler(rateService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code)

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewShortFilmCollectionResponse(userCollections)

	var actualBody []models.ShortFilmCollectionResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetUserCollectionsHandler_Action_NotOK_AuthService(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/collections?sort_param=create_time&count_collections=15&delimiter=now", nil)

	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	var userCollections []modelsGlobal.Collection

	userCollections = append(userCollections, modelsGlobal.Collection{
		ID:         12,
		Name:       "Избранное",
		Poster:     "42",
		CountLikes: 1023,
		CountFilms: 10,
		UpdateTime: "2020.12.12 15:15:15",
		CreateTime: "2012.06.05 01:25:00",
	})

	// Settings mock
	rateService.EXPECT().GetUserCollections(r.Context(), &user, &constparams.GetUserCollectionsParams{
		SortParam:        "create_time",
		CountCollections: 15,
		Delimiter:        "now",
	}).Return(userCollections, errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetUserCollectionsHandler(rateService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode")
	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetUserCollectionsHandler_Action_CountCol_Empty(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/collections?sort_param=create_time&count_collections=&delimiter=now", nil)

	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetUserCollectionsHandler(rateService)
	handler.Configure(router, nil)

	handler.Action(w, r)

	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")
	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetUserCollectionsHandler_Action_Fatal_Error(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/collections?sort_param=create_time&count_collections=&delimiter=now", nil)

	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrGetUserRequest.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetUserCollectionsHandler(rateService)
	handler.Configure(router, nil)

	handler.Action(w, r)

	require.Equal(t, http.StatusInternalServerError, w.Code, "Wrong StatusCode")
	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetUserCollectionsHandler_Action_CountCol_Wrong(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/collections?sort_param=create_time&count_collections=adsa&delimiter=now", nil)

	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetUserCollectionsHandler(rateService)
	handler.Configure(router, nil)

	handler.Action(w, r)

	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")
	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetUserCollectionsHandler_Action_SortPar_Empty(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/collections?sort_param=&count_collections=14&delimiter=now", nil)

	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetUserCollectionsHandler(rateService)
	handler.Configure(router, nil)

	handler.Action(w, r)

	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")
	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetUserCollectionsHandler_Action_Delimiter_Empty(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/collections?sort_param=create_time&count_collections=14&delimiter=", nil)

	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetUserCollectionsHandler(rateService)
	handler.Configure(router, nil)

	handler.Action(w, r)

	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")
	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
