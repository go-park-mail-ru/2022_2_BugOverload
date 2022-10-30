package tests_test

import (
	"context"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestLoginHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"YasaPupkinEzji@top.world","password":"Widget Adapter"}`,

			ResponseCookie: "GeneratedData",
			ResponseBody:   `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"avatar"}`,
			StatusCode:     http.StatusOK,
		},
		// No such user
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"YasaPupkinEzji@top.world123","password":"Widget Adapter"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrUserNotExist)),
			StatusCode:   http.StatusNotFound,
		},
		// Wrong password
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"YasaPupkinEzji@top.world","password":"Widget 123123123Adapter"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrLoginCombinationNotFound)),
			StatusCode:   http.StatusForbidden,
		},
		// Broken JSON
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email": 123, "password": "Widget Adapter"`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrValidation(errors.ErrCJSONUnexpectedEnd)),
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
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrEmptyFieldAuth)),
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email":"YasaPupkinEzji@top.world"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrAuth(errors.ErrEmptyFieldAuth)),
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		tests.TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: pkg.NewTestErrorResponse(errors.NewErrValidation(errors.ErrContentTypeUndefined)),
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/signup"
	us := repoAuth.NewAuthCache()
	cs := repoSession.NewSessionCache()

	testUser := &models.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
	}

	_, err := us.CreateUser(context.TODO(), testUser)
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err create user for test"))

	userService := serviceAuth.NewUserService(us)
	authService := serviceSession.NewSessionService(cs)
	loginHandler := handlers.NewLoginHandler(userService, authService)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		w := httptest.NewRecorder()

		loginHandler.Action(w, req)

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		resp := w.Result()

		if item.ResponseCookie != "" {
			cookie := resp.Cookies()[0]

			require.Equal(t, "session_id", cookie.Name, pkg.TestErrorMessage(caseNum, "Created and received cookie not equal"))
		}

		if item.ResponseBody != "" {
			var body []byte

			body, err = io.ReadAll(resp.Body)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

			err = resp.Body.Close()
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

			require.Equal(t, item.ResponseBody, string(body), pkg.TestErrorMessage(caseNum, "Wrong body"))
		}
	}
}
