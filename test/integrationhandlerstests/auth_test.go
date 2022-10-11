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
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/authhandler"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
	"go-park-mail-ru/2022_2_BugOverload/test/integrationhandlerstests"
)

func TestAuthHandler(t *testing.T) {
	cases := []integrationhandlerstests.TestCase{
		// Success
		integrationhandlerstests.TestCase{
			Method:       http.MethodGet,
			Cookie:       "GeneratedData",
			ResponseBody: `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"asserts/img/invisibleMan.jpeg"}`,
			StatusCode:   http.StatusOK,
		},
		// Wrong cookie
		integrationhandlerstests.TestCase{
			Method:       http.MethodGet,
			Cookie:       "2=YasaPupkinEzji@top.world",
			ResponseBody: `{"error":"Action: [no such cookie]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
		// Cookie is missing
		integrationhandlerstests.TestCase{
			Method:       http.MethodGet,
			ResponseBody: `{"error":"Action: [request has no cookies]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
	}

	url := "http://localhost:8088/v1/auth"

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

	var cookie string
	cookie, err = cs.CreateSession(context.TODO(), testUser)
	require.Nil(t, err, utils.TestErrorMessage(-1, "Err create session-cookie for test"))

	cases[0].Cookie = strings.Split(cookie, ";")[0]

	userService := serviceUser.NewUserService(us, params.ContextTimeout)
	authService := serviceAuth.NewAuthService(cs, params.ContextTimeout)
	authHandler := authhandler.NewHandler(userService, authService)

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		if item.Cookie != "" {
			req.Header.Set("Cookie", item.Cookie)
		}

		w := httptest.NewRecorder()

		authHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))

		var body []byte
		body, err = io.ReadAll(resp.Body)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Body.Close must be success"))

		require.Equal(t, item.ResponseBody, string(body), utils.TestErrorMessage(caseNum, "Wrong body"))
	}
}
