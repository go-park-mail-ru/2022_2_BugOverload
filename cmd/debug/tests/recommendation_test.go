package tests_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	"go-park-mail-ru/2022_2_BugOverload/internal/film/delivery/handlers"
	repoFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/repository"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/film/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
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

	// Base
	us := repoAuth.NewAuthCache()
	cs := repoSession.NewSessionCache()

	testUser := &models.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
	}

	_, err := us.CreateUser(context.TODO(), testUser)
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err create user for test"))

	var session models.Session
	session, err = cs.CreateSession(context.TODO(), testUser)
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err create session-cookie for test"))

	cases[0].Cookie = "session_id=" + session.ID + ";"

	authService := serviceSession.NewSessionService(cs)

	// Films
	pathPreview := "../../../test/data/preview.json"

	fs := repoFilms.NewFilmCache(pathPreview)

	filmsService := serviceFilms.NewFilmService(fs)

	recommendationHandler := handlers.NewRecommendationFilmHandler(filmsService, authService)

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
			var body []byte
			body, err = io.ReadAll(resp.Body)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

			err = resp.Body.Close()
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

			require.True(t, string(body) != "", pkg.TestErrorMessage(caseNum, "Wrong body"))
		}
	}
}
