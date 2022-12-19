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

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
)

func TestUserNewFilmReviewHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]string{
		"body":   "dwdwqdqdw",
		"filmId": "1",
		"name":   "asasd",
		"type":   "positive",
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/review/new", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	user := models.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	service.EXPECT().NewFilmReview(r.Context(), &user, &models.Review{
		Name: "asasd",
		Type: "positive",
		Body: "dwdwqdqdw",
	}, &constparams.NewFilmReviewParams{
		FilmID: 1,
	}).Return(nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmReviewHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusCreated, w.Code)
}

func TestUserNewFilmReviewHandler_Action_NotOK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]string{
		"body":   "dwdwqdqdw",
		"filmId": "1",
		"name":   "asasd",
		"type":   "positive",
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/review/new", bytes.NewReader(body))

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrGetUserRequest.Error(),
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmReviewHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusInternalServerError, w.Code)

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

func TestUserNewFilmReviewHandler_Action_EmpBody(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/review/new", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrEmptyBody.Error(),
	}

	r.Header.Set("Content-Type", "application/json")

	user := models.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmReviewHandler(service)
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

func TestUserNewFilmReviewHandler_Action_ServiceError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]string{
		"body":   "dwdwqdqdw",
		"filmId": "1",
		"name":   "asasd",
		"type":   "positive",
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/review/new", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	user := models.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedErr := errors.ErrWorkDatabase

	service.EXPECT().NewFilmReview(r.Context(), &user, &models.Review{
		Name: "asasd",
		Type: "positive",
		Body: "dwdwqdqdw",
	}, &constparams.NewFilmReviewParams{
		FilmID: 1,
	}).Return(expectedErr)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmReviewHandler(service)
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
		ErrMassage: expectedErr.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
