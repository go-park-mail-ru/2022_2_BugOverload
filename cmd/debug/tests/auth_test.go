package tests_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestAuthHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:       http.MethodGet,
			Cookie:       "GeneratedData",
			ResponseBody: `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"asserts/img/invisibleMan.jpeg"}`,
			StatusCode:   http.StatusOK,
		},
		// Wrong cookie
		tests.TestCase{
			Method:       http.MethodGet,
			Cookie:       "2=YasaPupkinEzji@top.world",
			ResponseBody: `{"error":"Auth: [no such cookie]"}`,
			StatusCode:   http.StatusNotFound,
		},
		// Cookie is missing
		tests.TestCase{
			Method:       http.MethodGet,
			ResponseBody: `{"error":"Auth: [request has no cookies]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
	}

	url := "http://localhost:8088/v1/auth"

	us := memoryUser.NewUserCache()
	cs := memoryCookie.NewCookieCache()

	testUser := &models.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	}

	_, err := us.CreateUser(context.TODO(), testUser)
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err create user for test"))

	var cookie string
	cookie, err = cs.CreateSession(context.TODO(), testUser)
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err create session-cookie for test"))

	cases[0].Cookie = strings.Split(cookie, ";")[0]

	userService := serviceUser.NewUserService(us)
	authService := serviceAuth.NewAuthService(cs)
	authHandler := handlers.NewAuthHandler(userService, authService)

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		if item.Cookie != "" {
			req.Header.Set("Cookie", item.Cookie)
		}

		w := httptest.NewRecorder()

		authHandler.Action(w, req)

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseBody != "" {
			resp := w.Result()

			var body []byte

			body, err = io.ReadAll(resp.Body)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

			err = resp.Body.Close()
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

			require.Equal(t, item.ResponseBody, string(body), pkg.TestErrorMessage(caseNum, "Wrong body"))
		}
	}
}
