package handlers

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	mockWarehouseClient "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client/mocks"
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
)

func TestPersonHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	res := modelsGlobal.Person{
		Name:         "Шон Коннери",
		OriginalName: "Sean Connery",
		Professions:  []string{"актер", "продюсер", "режиссер"},
		Images:       []string{"1", "2"},
		GrowthMeters: 1.9,
		Genres:       []string{"драма", "боевик", "триллер"},
		Gender:       "male",
		CountFilms:   218,
		Birthday:     "1930.08.25",
		Death:        "2020.10.31",
		BestFilms: []modelsGlobal.Film{{
			Name:      "Игра престолов",
			ProdDate:  "2013",
			EndYear:   "2014",
			ID:        123,
			Rating:    7.12332,
			PosterVer: "123",
			Genres:    []string{"фэнтези", "приключения"},
		}},
		Avatar: "12",
	}

	service.EXPECT().GetPersonByID(r.Context(), &modelsGlobal.Person{ID: 1}, &constparams.GetPersonParams{
		CountImages: 2,
		CountFilms:  1,
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPersonHandler(service)
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

	expectedBody := models.NewPersonResponse(&res)

	var actualBody *models.PersonResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestPersonHandler_Action_NotOKService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	personService.EXPECT().GetPersonByID(r.Context(), &modelsGlobal.Person{ID: 1}, &constparams.GetPersonParams{
		CountImages: 2,
		CountFilms:  1,
	}).Return(modelsGlobal.Person{}, errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

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

func TestPersonHandler_Action_ErrBind_ErrConvertQuery_Params_CountFilms(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=ds&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

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

func TestPersonHandler_Action_ErrBind_ErrConvertQuery_Params_CountFilms_Empty(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

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

func TestPersonHandler_Action_ErrBind_ErrConvertQuery_CountImages(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=dsd", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

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

func TestPersonHandler_Action_ErrBind_ErrConvertQuery_Params_CountImages_Empty(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=2&count_images=", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParamsEmptyRequiredFields.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

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

func TestPersonHandler_Action_ErrBind_ErrBadQueryParams_CountFilms(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=-1&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParams.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

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

func TestPersonHandler_Action_ErrBind_ErrBadQueryParams_CountImages(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=-2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrBadRequestParams.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

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
