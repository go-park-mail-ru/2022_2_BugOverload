package dependenttest_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/handlers"
	repoFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/sqltools"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestRecommendationHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method: http.MethodGet,
			Cookie: "GeneratedData",

			ResponseBody: "GeneratedData",
			StatusCode:   http.StatusOK,
		},
	}

	url := "http://localhost:8088/v1/auth"

	// Films
	postgres := sqltools.NewPostgresRepository()
	fs := repoFilms.NewFilmPostgres(postgres)

	filmsService := serviceFilms.NewFilmService(fs)

	recommendationHandler := handlers.NewRecommendationFilmHandler(filmsService)

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		if item.Cookie != "" {
			req.Header.Set("Cookie", item.Cookie)
		}

		w := httptest.NewRecorder()

		recommendationHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseBody != "" {
			body, err := io.ReadAll(resp.Body)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

			err = resp.Body.Close()
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

			require.True(t, string(body) != "", pkg.TestErrorMessage(caseNum, "Wrong body"))
		}
	}
}
