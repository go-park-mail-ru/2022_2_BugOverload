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
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/models"
	mockCollectionService "go-park-mail-ru/2022_2_BugOverload/internal/collection/service/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

func TestTagCollectionHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection/popular?count_films=1&delimiter=10", nil)
	vars := make(map[string]string)
	vars["tag"] = "popular"
	r = mux.SetURLVars(r, vars)

	res := modelsGlobal.Collection{
		Name: "популярное",
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

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	r = r.WithContext(ctx)

	collectionService.EXPECT().GetCollectionByTag(r.Context(), &pkg.GetCollectionTagParams{
		Tag:        "popular",
		CountFilms: 1,
		Delimiter:  "10",
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewTagCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	tagCollectionHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewTagCollectionResponse(&res)

	var actualBody *models.TagCollectionResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestTagCollectionHandler_Action_NotOKService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection/popular?count_films=1&delimiter=10", nil)
	vars := make(map[string]string)
	vars["tag"] = "popular"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	r = r.WithContext(ctx)

	collectionService.EXPECT().GetCollectionByTag(r.Context(), &pkg.GetCollectionTagParams{
		Tag:        "popular",
		CountFilms: 1,
		Delimiter:  "10",
	}).Return(modelsGlobal.Collection{}, errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewTagCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	tagCollectionHandler.Action(w, r)

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

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestTagCollectionHandler_Action_ErrBind_ErrConvertQuery(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection/popular?count_films=вфы&delimiter=10", nil)
	vars := make(map[string]string)
	vars["tag"] = "popular"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewTagCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestTagCollectionHandler_Action_ErrBind_ErrBadQueryParams(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection/popular?count_films=-1&delimiter=10", nil)
	vars := make(map[string]string)
	vars["tag"] = "popular"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrBadQueryParams.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewTagCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
