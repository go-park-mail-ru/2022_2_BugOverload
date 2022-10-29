package tests_test

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/handlers"
	repoAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	repoSession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestLogoutHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:         http.MethodGet,
			Cookie:         "GeneratedData",
			ResponseCookie: "1=YasaPupkinEzji@top.world",
			StatusCode:     http.StatusNoContent,
		},
		// Cookie has been deleted
		tests.TestCase{
			Method:       http.MethodGet,
			Cookie:       "1=YasaPupkinEzji@top.world",
			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrSessionNotExist)),
			StatusCode:   http.StatusNotFound,
		},
		// Wrong session
		tests.TestCase{
			Method:       http.MethodGet,
			Cookie:       "2=YasaPupkinEzji@top.world",
			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrSessionNotExist)),
			StatusCode:   http.StatusNotFound,
		},
		// Cookie is missing
		tests.TestCase{
			Method:       http.MethodGet,
			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrNoCookie)),
			StatusCode:   http.StatusUnauthorized,
		},
	}

	url := "http://localhost:8088/v1/auth/logput"

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
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err create session-session for test"))

	cases[0].Cookie = "session_id=" + session.ID + ";"
	cases[1].Cookie = "session_id=" + session.ID + ";"

	userService := serviceAuth.NewUserService(us)
	authService := serviceSession.NewSessionService(cs)
	logoutHandler := handlers.NewLogoutHandler(userService, authService)

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		if item.Cookie != "" {
			req.Header.Set("Cookie", item.Cookie)
		}

		w := httptest.NewRecorder()

		logoutHandler.Action(w, req)

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseCookie != "" {
			resp := w.Result()

			cookieResp := resp.Cookies()[0]

			require.Contains(t, cookieResp.Name, "session_id",
				pkg.TestErrorMessage(caseNum, "Created and received session not equal"))
		}
	}
}
