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

func TestUserFilmRateDropHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/rate/drop", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	resRate := modelsGlobal.Film{
		CountRatings: 12,
		Rating:       6,
	}

	// Settings mock
	rateService.EXPECT().FilmRateDrop(r.Context(), &user, &constparams.FilmRateDropParams{
		FilmID: 1,
	}).Return(resRate, nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmRateDropHandler(rateService)
	handler.Configure(router, nil)

	// CheckNewNotification result
	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusOK, w.Code)

	// CheckNewNotification body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewFilmRateDropResponse(&resRate)

	var actualBody *models.FilmRateDropResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserFilmRateDropHandler_Action_NotOK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/rate/drop", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	resRate := modelsGlobal.Film{
		CountRatings: 12,
		Rating:       6,
	}

	// Settings mock
	rateService.EXPECT().FilmRateDrop(r.Context(), &user, &constparams.FilmRateDropParams{
		FilmID: 1,
	}).Return(resRate, errors.ErrNotFoundInDB)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmRateDropHandler(rateService)
	handler.Configure(router, nil)

	// CheckNewNotification result
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

func TestUserFilmRateDropHandler_Action_InvParam(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/rate/drop", nil)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmRateDropHandler(rateService)
	handler.Configure(router, nil)

	// CheckNewNotification result
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

func TestUserFilmRateDropHandler_Action_UserNotFound(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/rate/drop", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmRateDropHandler(rateService)
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
