package integrationhandlerstests_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/loginhandler"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
	"go-park-mail-ru/2022_2_BugOverload/test/integrationhandlerstests"
)

func TestLoginHandler(t *testing.T) {
	cases := []integrationhandlerstests.TestCase{
		// Success
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"YasaPupkinEzji@top.world","password":"Widget Adapter"}`,

			ResponseCookie: "GeneratedData",
			ResponseBody:   `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"asserts/img/invisibleMan.jpeg"}`,
			StatusCode:     http.StatusOK,
		},
		// Wrong password
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"YasaPupkinEzji@top.world","password":"Widget 123123123Adapter"}`,

			ResponseBody: `{"error":"Action: [no such combination of login and password]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
		// Broken JSON
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email": 123, "password": "Widget Adapter"`,

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body not JSON
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/xml",
			RequestBody: `<Name>Ellen Adams</Name>`,

			ResponseBody: `{"error":"Def validation: [unsupported media type]"}`,
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Empty required field - email
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"YasaPupkinEzji@top.world"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		integrationhandlerstests.TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: `{"error":"Def validation: [content-type undefined]"}`,
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/signup"

	userMutex := &sync.Mutex{}
	authMutex := &sync.Mutex{}

	us := memoryUser.NewUserRepo(userMutex)
	cs := memoryCookie.NewCookieRepo(authMutex)

	testUser := &models.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	}

	_, err := us.CreateUser(context.TODO(), testUser)
	require.Nil(t, err, utils.TestErrorMessage(-1, "Err create user for test"))

	userService := serviceUser.NewUserService(us, params.ContextTimeout)
	authService := serviceAuth.NewAuthService(cs, params.ContextTimeout)
	loginHandler := loginhandler.NewHandler(userService, authService)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		w := httptest.NewRecorder()

		loginHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseCookie != "" {
			respCookie := resp.Header.Get("Set-Cookie")

			cookieName := strings.Split(respCookie, ";")[0]

			ctx := context.WithValue(context.TODO(), params.CookieKey, cookieName)

			var nameSession string
			nameSession, err = authService.GetSession(ctx)
			require.Nil(t, err, utils.TestErrorMessage(caseNum, "Result GetSession not error"))

			require.Equal(t, respCookie, nameSession, utils.TestErrorMessage(caseNum, "Created and received cookie not equal"))
		}

		var body []byte
		body, err = io.ReadAll(resp.Body)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Body.Close must be success"))

		require.Equal(t, item.ResponseBody, string(body), utils.TestErrorMessage(caseNum, "Wrong body"))
	}
}
