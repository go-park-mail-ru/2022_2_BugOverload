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

	mockFilmService "go-park-mail-ru/2022_2_BugOverload/internal/film/service/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/httpwrapper"
)

func TestRecommendationHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/recommendation", nil)

	expectedBody := models.Film{
		Name:      "Игра престолов",
		ProdYear:  "2013",
		EndYear:   "2014",
		ID:        123,
		Rating:    7.12332,
		PosterHor: "123",
		Genres:    []string{"фэнтези", "приключения"},
	}

	filmService.EXPECT().GetRecommendation(r.Context()).Return(expectedBody, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	recommendationHandler := NewRecommendationFilmHandler(filmService)
	recommendationHandler.Configure(router, nil)

	recommendationHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody models.Film

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, actualBody, expectedBody, "Wrong body")
}

func TestRecommendationHandler_Action_NotOKService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	filmService := mockFilmService.NewMockFilmsService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/film/recommendation", nil)

	expectedBody := httpwrapper.ErrResponse{
		ErrMassage: errors.ErrWorkDatabase.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	r = r.WithContext(ctx)

	filmService.EXPECT().GetRecommendation(r.Context()).Return(models.Film{}, errors.ErrWorkDatabase)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	recommendationHandler := NewRecommendationFilmHandler(filmService)
	recommendationHandler.Configure(router, nil)

	recommendationHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusInternalServerError, w.Code, "Wrong StatusCode")

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
