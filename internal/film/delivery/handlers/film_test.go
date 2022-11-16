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

	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/models"
	mockFilmService "go-park-mail-ru/2022_2_BugOverload/internal/film/service/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

func TestFilmHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1?count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	res := modelsGlobal.Film{
		Name: "Игра престолов",
		Actors: []modelsGlobal.FilmActor{{
			Name:      "Питер Динклэйдж",
			ID:        1,
			Character: "some",
			Avatar:    "1",
		}},
		Artists: []modelsGlobal.FilmPerson{{
			Name: "Питер Динклэйдж",
			ID:   1,
		}},
		Producers: []modelsGlobal.FilmPerson{{
			Name: "Питер Динклэйдж",
			ID:   1,
		}},
		Composers: []modelsGlobal.FilmPerson{{
			Name: "Питер Динклэйдж",
			ID:   1,
		}},
		Directors: []modelsGlobal.FilmPerson{{
			Name: "Питер Динклэйдж",
			ID:   1,
		}},
		Montage: []modelsGlobal.FilmPerson{{
			Name: "Питер Динклэйдж",
			ID:   1,
		}},
		Operators: []modelsGlobal.FilmPerson{{
			Name: "Питер Динклэйдж",
			ID:   1,
		}},
		Writers: []modelsGlobal.FilmPerson{{
			Name: "Питер Динклэйдж",
			ID:   1,
		}},
		AgeLimit:             "18+",
		BoxOfficeDollars:     60000000,
		Budget:               10000000,
		CountActors:          1,
		CountPositiveReviews: 1,
		CountNeutralReviews:  1,
		CountNegativeReviews: 1,
		CountRatings:         1,
		CountSeasons:         8,
		CurrencyBudget:       "USD",
		Description:          "Британская лингвистка Алетея прилетает из Лондона",
		ShortDescription:     "Много насилия и фэнтези",
		DurationMinutes:      55,
		EndYear:              "2019",
		Genres:               []string{"фантастика", "боевик"},
		Images:               []string{"1", "2"},
		OriginalName:         "Game of Thrones",
		PosterHor:            "1",
		ProdCountries:        []string{"США", "Великобритания"},
		ProdCompanies:        []string{"HBO"},
		ProdYear:             "2011",
		Rating:               9.2,
		Type:                 "serial",
		Slogan:               "Победа или смерть",
	}

	filmService.EXPECT().GetFilmByID(r.Context(), &modelsGlobal.Film{ID: 1}, &pkg.GetFilmParams{
		CountImages: 2,
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewFilmHandler(filmService)
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

	expectedBody := models.NewFilmResponse(&res)

	var actualBody *models.FilmResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestFilmHandler_Action_NotOKService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1?count_images=2", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	r = r.WithContext(ctx)

	filmService.EXPECT().GetFilmByID(r.Context(), &modelsGlobal.Film{ID: 1}, &pkg.GetFilmParams{
		CountImages: 2,
	}).Return(modelsGlobal.Film{}, errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewFilmHandler(filmService)
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

func TestFilmHandler_Action_ErrBind_ErrConvertQuery_Params(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1?count_images=ddas", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewFilmHandler(filmService)
	filmHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	filmHandler.Action(w, r.WithContext(ctx))

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

func TestFilmHandler_Action_ErrBind_ErrBadQueryParams(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/1?count_images=-1", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrBadQueryParams.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	filmHandler := NewFilmHandler(filmService)
	filmHandler.Configure(router, nil)

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	filmHandler.Action(w, r.WithContext(ctx))

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
