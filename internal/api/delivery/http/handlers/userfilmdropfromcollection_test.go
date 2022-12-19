package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
)

func TestDropFilmFromUserCollectionHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]int{
		"collection_id": 401,
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/remove", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	service.EXPECT().DropFilmFromUserCollection(r.Context(), &user, &constparams.UserCollectionFilmsUpdateParams{
		FilmID:       1,
		CollectionID: 401,
	}).Return(nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewDropFilmFromUserCollectionHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestDropFilmFromUserCollectionHandler_Action_NotOK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]int{
		"collection_id": 401,
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/remove", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	service.EXPECT().DropFilmFromUserCollection(r.Context(), &user, &constparams.UserCollectionFilmsUpdateParams{
		FilmID:       1,
		CollectionID: 401,
	}).Return(errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewDropFilmFromUserCollectionHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusNotFound, w.Code)
	// CheckNewNotification body
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

func TestDropFilmFromUserCollectionHandler_Action_InvBody(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/remove", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrEmptyBody.Error(),
	}

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewDropFilmFromUserCollectionHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusBadRequest, w.Code)
	// CheckNewNotification body
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

func TestDropFilmFromUserCollectionHandler_Action_InvId(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]int{
		"collection_id": 401,
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/remove", bytes.NewReader(body))

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewDropFilmFromUserCollectionHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusBadRequest, w.Code)
	// CheckNewNotification body
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

func TestDropFilmFromUserCollectionHandler_Action_UserNotFound(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]int{
		"collection_id": 401,
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/remove", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewDropFilmFromUserCollectionHandler(service)
	handler.Configure(router, nil)

	// CheckNewNotification result
	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusInternalServerError, w.Code, "Wrong StatusCode")

	// CheckNewNotification body
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	require.Nil(t, errResponse, "io.ReadAll must be success")

	err := response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrGetUserRequest.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
