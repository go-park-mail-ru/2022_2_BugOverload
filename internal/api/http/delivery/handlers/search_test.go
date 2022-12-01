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

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	mockWarehouseClient "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client/mocks"
)

func TestSearchHandlerPremiere_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	searchService := mockWarehouseClient.NewMockWarehouseService(ctrl)

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

	searchService.EXPECT().Search(r.Context(), &constparams.SearchParams{ //
		Query: "abc", //
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSearchHandler(searchService) //
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

	expectedBody, _ := models.NewSearchResponse(res) //

	var actualBody *models.SearchResponse //

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, &expectedBody, actualBody, "Wrong body")
}
