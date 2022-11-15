package handlers

import (
	"encoding/json"
	mockPersonService "go-park-mail-ru/2022_2_BugOverload/internal/person/service/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

func TestPersonHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockPersonService.NewMockPersonService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := models.Person{
		Name:         "Шон Коннери",
		OriginalName: "Sean Connery",
		Professions:  []string{"актер", "продюсер", "режиссер"},
		Images:       []string{"1", "2"},
		Growth:       1.9,
		Genres:       []string{"драма", "боевик", "триллер"},
		Gender:       "male",
		CountFilms:   218,
		Birthday:     "1930.08.25",
		Death:        "2020.10.31",
		BestFilms: []models.Film{{
			Name:      "Игра престолов",
			ProdYear:  "2013",
			EndYear:   "2014",
			ID:        123,
			Rating:    7.12332,
			PosterVer: "123",
			Genres:    []string{"фэнтези", "приключения"},
		}},
		Avatar: "12",
	}

	personService.EXPECT().GetPersonByID(r.Context(), &models.Person{ID: 1}, &pkg.GetPersonParams{
		CountImages: 2,
		CountFilms:  1,
	}).Return(expectedBody, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewPersonHandler(personService)
	filmHandler.Configure(router, nil)

	filmHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody models.Person

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestPersonHandler_Action_NotOKService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockPersonService.NewMockPersonService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	personService.EXPECT().GetPersonByID(r.Context(), &models.Person{ID: 1}, &pkg.GetPersonParams{
		CountImages: 2,
		CountFilms:  1,
	}).Return(models.Person{}, errors.ErrNotFoundInDB)

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

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestPersonHandler_Action_ErrBind_ErrConvertQuery_Params_CountFilms(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockPersonService.NewMockPersonService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=ds&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
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

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestPersonHandler_Action_ErrBind_ErrConvertQuery_CountImages(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockPersonService.NewMockPersonService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=dsd", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
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

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestPersonHandler_Action_ErrBind_ErrBadQueryParams_CountFilms(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockPersonService.NewMockPersonService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=-1&count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrBadQueryParams.Error(),
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

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestPersonHandler_Action_ErrBind_ErrBadQueryParams_CountImages(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	personService := mockPersonService.NewMockPersonService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/person/1?count_films=1&count_images=-2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrBadQueryParams.Error(),
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

	var actualBody httpwrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}
