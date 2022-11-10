package tests_test

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	"go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/handlers"
	memoryAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/repository"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/auth/service"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	memorySession "go-park-mail-ru/2022_2_BugOverload/internal/session/repository"
	serviceSession "go-park-mail-ru/2022_2_BugOverload/internal/session/service"
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
			ResponseBody:   `{"nickname":"testnickname","email":"testmail@yandex.ru","avatar":"avatar"}`,
			StatusCode:     http.StatusCreated,
		},
		// Such user exists
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrSignupUserExist)),
			StatusCode:   http.StatusBadRequest,
		},
		// Broken JSON
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","password":"testpassword"`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrValidation(errors.ErrJSONUnexpectedEnd)),
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrValidation(errors.ErrEmptyBody)),
			StatusCode:   http.StatusBadRequest,
		},
		// Body not JSON
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/xml",
			RequestBody: `<Name>Ellen Adams</Name>`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrValidation(errors.ErrUnsupportedMediaType)),
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Empty required field - email
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrEmptyFieldAuth)),
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrEmptyFieldAuth)),
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - nickname
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"testmail@yandex.ru","password": "testpassword"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrEmptyFieldAuth)),
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		tests.TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrValidation(errors.ErrContentTypeUndefined)),
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/signup"

	us := memoryAuth.NewAuthCache()
	cs := memorySession.NewSessionCache()

	authService := serviceAuth.NewAuthService(us)
	sessionService := serviceSession.NewSessionService(cs)
	signupHandler := handlers.NewSingUpHandler(authService, sessionService)

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
			cookie := resp.Cookies()[0]

			require.Equal(t, "session_id", cookie.Name, pkg.TestErrorMessage(caseNum, "Created and received cookie not equal"))
		}

		body, err := io.ReadAll(resp.Body)
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

		require.Equal(t, item.ResponseBody, string(body), pkg.TestErrorMessage(caseNum, "Wrong body"))
	}
}
