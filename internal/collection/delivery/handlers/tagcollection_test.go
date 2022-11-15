package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	mockCollectionService "go-park-mail-ru/2022_2_BugOverload/internal/collection/service/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

func TestTagCollectionHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/collection/popular?count_films=1&delimiter=10", nil)
	vars := make(map[string]string)
	vars["tag"] = "popular"
	r = mux.SetURLVars(r, vars)

	expectedBody := models.Collection{
		Name: "популярное",
		Films: []models.Film{models.Film{
			Name:      "Игра престолов",
			ProdYear:  "2013",
			EndYear:   "2014",
			ID:        123,
			Rating:    7.12332,
			PosterVer: "123",
			Genres:    []string{"фэнтези", "приключения"},
		}},
	}

	collectionService.EXPECT().GetCollectionByTag(r.Context(), &pkg.GetCollectionTagParams{
		Tag:        "popular",
		CountFilms: 1,
		Delimiter:  "10",
	}).Return(expectedBody, nil)

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

	var actualBody models.Collection

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestTagCollectionHandler_Action_NotOk(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/collection/popular?count_films=1&delimiter=10", nil)
	vars := make(map[string]string)
	vars["tag"] = "popular"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	collectionService.EXPECT().GetCollectionByTag(r.Context(), &pkg.GetCollectionTagParams{
		Tag:        "popular",
		CountFilms: 1,
		Delimiter:  "10",
	}).Return(models.Collection{}, errors.ErrNotFoundInDB)

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

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestTagCollectionHandler_Action_ErrBind_ErrConvertQuery(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/collection/popular?count_films=вфы&delimiter=10", nil)
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

	tagCollectionHandler.Action(w, r)

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

func TestTagCollectionHandler_Action_ErrBind_ErrQueryBad(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/collection/popular?count_films=-1&delimiter=10", nil)
	vars := make(map[string]string)
	vars["tag"] = "popular"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrQueryBad.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewTagCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	tagCollectionHandler.Action(w, r)

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
