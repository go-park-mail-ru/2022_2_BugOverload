package handlers

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	mockFilmService "go-park-mail-ru/2022_2_BugOverload/internal/film/service/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

func TestReviewsHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=1&offset=0", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	res := []modelsGlobal.Review{{
		Name:       "Много насилия и только",
		Type:       "negative",
		Body:       "Игра престолов дюже наполнена насилием и смертью",
		CountLikes: 12,
		CreateTime: "2019.12.31",
		Author: modelsGlobal.User{
			ID:       12,
			Nickname: "steepbyy",
			Profile: modelsGlobal.Profile{
				Avatar:       "12",
				CountReviews: 44,
			},
		},
	}}

	filmService.EXPECT().GetReviewsByFilmID(r.Context(), &pkg.GetReviewsFilmParams{
		Count:  1,
		Offset: 0,
		FilmID: 1,
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	reviewsHandler := NewReviewsHandler(filmService)
	reviewsHandler.Configure(router, nil)

	reviewsHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewReviewsResponse(&res)

	var actualBody []*models.ReviewResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestReviewsHandler_Action_NotOKService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=1&offset=0", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	filmService.EXPECT().GetReviewsByFilmID(r.Context(), &pkg.GetReviewsFilmParams{
		Count:  1,
		Offset: 0,
		FilmID: 1,
	}).Return([]modelsGlobal.Review{}, errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	reviewsHandler := NewReviewsHandler(filmService)
	reviewsHandler.Configure(router, nil)

	reviewsHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrUnsupportedMediaType(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=asd&offset=0", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "app")

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrUnsupportedMediaType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	reviewsHandler := NewReviewsHandler(filmService)
	reviewsHandler.Configure(router, nil)

	reviewsHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusUnsupportedMediaType, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrConvertQuery_Params_CountReviews(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=asd&offset=0", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	reviewsHandler := NewReviewsHandler(filmService)
	reviewsHandler.Configure(router, nil)

	reviewsHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrConvertQuery_Offset(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=1&offset=asd", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	reviewsHandler := NewReviewsHandler(filmService)
	reviewsHandler.Configure(router, nil)

	reviewsHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrBadQueryParams_CountReviews(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=-1&offset=0", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrBadQueryParams.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	reviewsHandler := NewReviewsHandler(filmService)
	reviewsHandler.Configure(router, nil)

	reviewsHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
	r.Header.Set("Content-Type", "")
}

func TestPersonHandler_Action_ErrBind_ErrBadQueryParams_CountImages(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=1&offset=-1", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrBadQueryParams.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	reviewsHandler := NewReviewsHandler(filmService)
	reviewsHandler.Configure(router, nil)

	reviewsHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}
