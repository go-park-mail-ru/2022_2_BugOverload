package handlers

import (
	"context"
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
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
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockCollectionService "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/collection/service/mocks"
)

func TestTagCollectionHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=tag&key=popular&sort_param=date&count_films=1&delimiter=0", nil)

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

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	collectionService.EXPECT().GetStdCollection(r.Context(), &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		Key:        "popular",
		CountFilms: 1,
		Delimiter:  "0",
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
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

	expectedBody := models.NewStdCollectionResponse(&res)

	var actualBody *models.GetStdCollectionResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestTagCollectionHandler_Action_NotOKService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=tag&key=popular&sort_param=date&count_films=1&delimiter=0", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	collectionService.EXPECT().GetStdCollection(r.Context(), &constparams.GetStdCollectionParams{
		Target:     "tag",
		SortParam:  "date",
		Key:        "popular",
		CountFilms: 1,
		Delimiter:  "0",
	}).Return(modelsGlobal.Collection{}, errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
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

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestTagCollectionHandler_Action_ErrBind_ErrConvertQuery(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=tag&key=popular&sort_param=date&count_films=asd&delimiter=0", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

func TestTagCollectionHandler_Action_ErrBind_ErrBadQueryParams(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=tag&key=popular&sort_param=date&count_films=-1&delimiter=0", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParams.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

func TestTagCollectionHandler_Action_ErrBind_ErrBadQueryParamsEmpty_Target(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=&key=popular&sort_param=date&count_films=-1&delimiter=0", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

func TestTagCollectionHandler_Action_ErrBind_ErrBadQueryParamsEmpty_Key(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=asd&key=&sort_param=date&count_films=-1&delimiter=0", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

func TestTagCollectionHandler_Action_ErrBind_ErrBadQueryParamsEmpty_SortParam(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=asd&key=asd&sort_param=&count_films=-1&delimiter=0", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

func TestTagCollectionHandler_Action_ErrBind_ErrBadQueryParamsEmpty_CountFilms(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=asd&key=asd&sort_param=asd&count_films=&delimiter=0", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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

func TestTagCollectionHandler_Action_ErrBind_ErrBadQueryParamsEmpty_Delinmeter(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockCollectionService.NewMockCollectionService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/collection?target=asd&key=asd&sort_param=asd&count_films=12&delimiter=", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	tagCollectionHandler := NewStdCollectionHandler(collectionService)
	tagCollectionHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	tagCollectionHandler.Action(w, r.WithContext(ctx))

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
