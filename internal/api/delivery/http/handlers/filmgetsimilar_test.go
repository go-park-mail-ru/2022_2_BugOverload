package handlers

import (
	"context"
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	mockWarehouseClient "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client/mocks"
)

func TestGetSimilarFilmsHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/10/similar", nil)

	vars := make(map[string]string)
	vars["id"] = "10"
	r = mux.SetURLVars(r, vars)

	params := constparams.GetSimilarFilmsParams{
		FilmID: 10,
	}

	res := modelsGlobal.Collection{
		Name: "Похожие фильмы и сериалы",
		Films: []modelsGlobal.Film{{
			Name:      "Игра престолов",
			ProdDate:  "2013",
			EndYear:   "2014",
			ID:        123,
			Rating:    7.12332,
			PosterVer: "123",
			Genres:    []string{"фэнтези", "приключения"},
		}},
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	service.EXPECT().GetSimilarFilms(r.Context(), &params).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetSimilarFilmsHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewGetSimilarFilmsCollectionResponse(&res)

	var actualBody *models.GetSimilarFilmsCollectionResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestGetSimilarFilmsHandler_Action_BadRequest(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/10/similar", nil)

	vars := make(map[string]string)
	vars["id"] = "trash"
	r = mux.SetURLVars(r, vars)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetSimilarFilmsHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")
}

func TestGetSimilarFilmsHandler_Action_ServiceError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/10/similar", nil)

	vars := make(map[string]string)
	vars["id"] = "10"
	r = mux.SetURLVars(r, vars)

	params := constparams.GetSimilarFilmsParams{
		FilmID: 10,
	}

	expectedErr := errors.ErrNotFoundInDB

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	service.EXPECT().GetSimilarFilms(r.Context(), &params).Return(modelsGlobal.Collection{}, expectedErr)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetSimilarFilmsHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := &models.GetSimilarFilmsCollectionResponse{}

	var actualBody *models.GetSimilarFilmsCollectionResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
