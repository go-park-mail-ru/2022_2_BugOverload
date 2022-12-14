package handlers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockWarehouseClient "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client/mocks"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestSearchHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/search?q=abc", nil) //

	res := modelsGlobal.Search{ //
		Persons: []modelsGlobal.Person{{
			ID:           1,
			Name:         "Шон Коннери",
			OriginalName: "Sean Connery",
			Professions:  []string{"актер", "продюсер", "режиссер"},
			CountFilms:   218,
			Birthday:     "1930.08.25",
			Death:        "2020.10.31",
			Avatar:       "12",
		}},

		Serials: []modelsGlobal.Film{{
			Name:          "Игра престолов1",
			ProdDate:      "2013",
			EndYear:       "2014",
			ID:            123,
			Rating:        7.12332,
			PosterVer:     "123",
			ProdCountries: []string{"США"},
			Genres:        []string{"фэнтези", "приключения"},
			Directors:     []modelsGlobal.FilmPerson{{ID: 1, Name: "qqq"}},
		}},

		Films: []modelsGlobal.Film{{
			Name:          "Игра престолов2",
			ProdDate:      "2013",
			EndYear:       "2014",
			ID:            123,
			Rating:        7.12332,
			PosterVer:     "123",
			ProdCountries: []string{"США"},
			Genres:        []string{"фэнтези", "приключения"},
			Directors:     []modelsGlobal.FilmPerson{{ID: 1, Name: "qqq"}},
		}},
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	service.EXPECT().Search(r.Context(), &constparams.SearchParams{ //
		Query: "abc", //
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSearchHandler(service) //
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode", pkg.GetResponseBody(*w))

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody, _ := models.NewSearchResponse(res) //

	var actualBody models.SearchResponse //

	err = easyjson.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, &expectedBody, &actualBody, "Wrong body")
}

func TestSearchHandler_Action_NotOK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/search?q=abc", nil) //

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	res := modelsGlobal.Search{ //
		Persons: []modelsGlobal.Person{{
			ID:           1,
			Name:         "Шон Коннери",
			OriginalName: "Sean Connery",
			Professions:  []string{"актер", "продюсер", "режиссер"},
			CountFilms:   218,
			Birthday:     "1930.08.25",
			Death:        "2020.10.31",
			Avatar:       "12",
		}},

		Serials: []modelsGlobal.Film{{
			Name:          "Игра престолов1",
			ProdDate:      "2013",
			EndYear:       "2014",
			ID:            123,
			Rating:        7.12332,
			PosterVer:     "123",
			ProdCountries: []string{"США"},
			Genres:        []string{"фэнтези", "приключения"},
			Directors:     []modelsGlobal.FilmPerson{{ID: 1, Name: "qqq"}},
		}},

		Films: []modelsGlobal.Film{{
			Name:          "Игра престолов2",
			ProdDate:      "2013",
			EndYear:       "2014",
			ID:            123,
			Rating:        7.12332,
			PosterVer:     "123",
			ProdCountries: []string{"США"},
			Genres:        []string{"фэнтези", "приключения"},
			Directors:     []modelsGlobal.FilmPerson{{ID: 1, Name: "qqq"}},
		}},
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	service.EXPECT().Search(r.Context(), &constparams.SearchParams{ //
		Query: "abc", //
	}).Return(res, errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSearchHandler(service) //
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode", pkg.GetResponseBody(*w))

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = easyjson.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestSearchHandler_Action_Emp(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/search?q=", nil) //

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrQueryRequiredEmpty.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)
	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)
	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSearchHandler(service) //
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode", pkg.GetResponseBody(*w))

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = easyjson.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestSearchHandler_Action_EmptyResult(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/search?q=123", nil) //

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)
	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)
	r = r.WithContext(ctx)

	inputParams := constparams.SearchParams{
		Query: "123",
	}

	output := modelsGlobal.Search{
		Films:   []modelsGlobal.Film{},
		Serials: []modelsGlobal.Film{},
		Persons: []modelsGlobal.Person{},
	}

	service.EXPECT().Search(r.Context(), &inputParams).Return(output, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSearchHandler(service) //
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode", pkg.GetResponseBody(*w))

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = easyjson.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
