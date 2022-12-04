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
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	mocUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
)

func TestUserFilmRateDropHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mocUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/rate/drop", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	resRate := modelsGlobal.Film{
		CountRatings: 12,
		Rating:       6,
	}

	// Settings mock
	rateService.EXPECT().FilmRateDrop(r.Context(), &user, &constparams.FilmRateDropParams{
		FilmID: 1,
	}).Return(resRate, nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewFilmRateDropHandler(rateService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code)

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewFilmRateDropResponse(&resRate)

	var actualBody *models.FilmRateDropResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
