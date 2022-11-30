package handlers

import (
	"bytes"
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

func TestUserFilmRateHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mocUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]int{
		"score": 6,
	}
    body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/rate", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	resRate := modelsGlobal.Film{
		CountRatings: 12,
		Rating: 6,
	}

	rateService.EXPECT().FilmRate(r.Context(), &user, &constparams.FilmRateParams{
		FilmID: 1,
		Score:  6,
	}).Return(resRate, nil)


	w := httptest.NewRecorder()

	router := mux.NewRouter()
	rateHandler := NewFilmRateHandler(rateService)
	rateHandler.Configure(router, nil)

	rateHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code)

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewFilmRateResponse(&resRate)

	var actualBody *models.FilmRateResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
