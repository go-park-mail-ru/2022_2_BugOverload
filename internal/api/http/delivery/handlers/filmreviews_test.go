package handlers

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
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
	mockFilmService "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/film/service/mocks"
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
			ID:           12,
			Nickname:     "steepbyy",
			Avatar:       "12",
			CountReviews: 44,
		},
	}}

	filmService.EXPECT().GetReviewsByFilmID(r.Context(), &constparams.GetReviewsFilmParams{
		CountReviews: 1,
		Offset:       0,
		FilmID:       1,
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

	require.Equal(t, expectedBody, actualBody, "Wrong body")
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

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	filmService.EXPECT().GetReviewsByFilmID(r.Context(), &constparams.GetReviewsFilmParams{
		CountReviews: 1,
		Offset:       0,
		FilmID:       1,
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
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

	expectedBody := wrapper.ErrResponse{
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrConvertQueryParams_CountReviews_Empty(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=&offset=0", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrConvertQuery_Offset_Empty(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=1&offset=", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrBadQueryParams_CountReviews_ErrConvert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=asd&offset=0", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := wrapper.ErrResponse{
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
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

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParams.Error(),
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrBadQueryParams_Offset_ErrConvert(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=23&offset=asd", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := wrapper.ErrResponse{
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestReviewsHandler_Action_ErrBind_ErrBadQueryParams_Offset(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=1&offset=-1", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParams.Error(),
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestPersonHandler_Action_ErrBind_ErrBadQueryParams_CountImages_Empty(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1/reviews?count_reviews=1&offset=", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
