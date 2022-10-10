package integrationhandlerstests

import (
	"github.com/stretchr/testify/require"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/signuphandler"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
)

func TestSignupHandler(t *testing.T) {
	cases := []TestCase{
		// Success
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseCookie: "1=testmail@yandex.ru",
			ResponseBody:   `{"nickname":"testnickname","email":"testmail@yandex.ru","avatar":"asserts/img/invisibleMan.jpeg"}`,
			StatusCode:     http.StatusCreated,
		},
		// Such user exists
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Action: [such a login exists]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Broken JSON
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","password":"testpassword"`,

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body not JSON
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/xml",
			RequestBody: `<Name>Ellen Adams</Name>`,

			ResponseBody: `{"error":"Def validation: [unsupported media type]"}`,
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Empty required field - email
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - nickname
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","password": "testpassword"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Def validation: [content-type undefined]"}`,
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/signup"

	us := memoryUser.NewUserRepo()
	cs := memoryCookie.NewCookieRepo()

	userService := serviceUser.NewUserService(us, 2)
	authService := serviceAuth.NewAuthService(cs, 2)
	signupHandler := signuphandler.NewHandler(userService, authService)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		w := httptest.NewRecorder()

		signupHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseCookie != "" {
			respCookie := resp.Header.Get("Set-Cookie")

			require.Contains(t, respCookie, item.Cookie, utils.TestErrorMessage(caseNum, "Created and received cookie not equal"))
		}

		body, err := io.ReadAll(resp.Body)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Body.Close must be success"))

		require.Equal(t, item.ResponseBody, string(body), utils.TestErrorMessage(caseNum, "Wrong body"))
	}
}
