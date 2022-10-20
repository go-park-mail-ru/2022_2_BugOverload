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
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/user/delivery/handlers"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/user/repository"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/user/service"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestSignupHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseCookie: "1=testmail@yandex.ru",
			ResponseBody:   `{"nickname":"testnickname","email":"testmail@yandex.ru","avatar":"asserts/img/invisibleMan.jpeg"}`,
			StatusCode:     http.StatusCreated,
		},
		// Such user exists
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Auth: [such a login exists]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Broken JSON
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","password":"testpassword"`,

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,

			ResponseBody: `{"error":"Def validation: [empty body]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body not JSON
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/xml",
			RequestBody: `<Name>Ellen Adams</Name>`,

			ResponseBody: `{"error":"Def validation: [unsupported media type]"}`,
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Empty required field - email
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Auth: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname"}`,

			ResponseBody: `{"error":"Auth: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - nickname
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","password": "testpassword"}`,

			ResponseBody: `{"error":"Auth: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		tests.TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Def validation: [content-type undefined]"}`,
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/signup"

	us := memoryUser.NewUserCache()
	cs := memoryCookie.NewCookieCache()

	userService := serviceUser.NewUserService(us)
	authService := serviceAuth.NewAuthService(cs)
	signupHandler := handlers.NewSingUpHandler(userService, authService)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		w := httptest.NewRecorder()

		signupHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseCookie != "" {
			respCookie := resp.Header.Get("Set-Cookie")

			cookieName := strings.Split(respCookie, ";")[0]

			ctx := context.WithValue(context.TODO(), innerPKG.CookieKey, cookieName)
			nameSession, err := authService.GetSession(ctx)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Result GetSession not error"))

			require.Equal(t, respCookie, nameSession, pkg.TestErrorMessage(caseNum, "Created and received cookie not equal"))
		}

		body, err := io.ReadAll(resp.Body)
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

		require.Equal(t, item.ResponseBody, string(body), pkg.TestErrorMessage(caseNum, "Wrong body"))
	}
}
