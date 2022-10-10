package integrationhandlerstests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/logouthandler"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
)

func TestLogoutHandler(t *testing.T) {
	cases := []TestCase{
		// Success
		TestCase{
			Method:         http.MethodGet,
			Cookie:         "1=YasaPupkinEzji@top.world",
			ResponseCookie: "1=YasaPupkinEzji@top.world",
			StatusCode:     http.StatusNoContent,
		},
		// Cookie has been deleted
		TestCase{
			Method:       http.MethodGet,
			Cookie:       "1=YasaPupkinEzji@top.world",
			ResponseBody: `{"error":"Action: [no such cookie]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
		// Wrong cookie
		TestCase{
			Method:       http.MethodGet,
			Cookie:       "2=YasaPupkinEzji@top.world",
			ResponseBody: `{"error":"Action: [no such cookie]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
		// Cookie is missing
		TestCase{
			Method:       http.MethodGet,
			ResponseBody: `{"error":"Action: [request has no cookies]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
	}

	url := "http://localhost:8088/v1/auth/logput"

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

	us.CreateUser(context.TODO(), testUser)
	cookie, _ := cs.CreateSession(context.TODO(), testUser)

	cases[0].Cookie = strings.Split(cookie, ";")[0]

	userService := serviceUser.NewUserService(us, 2)
	authService := serviceAuth.NewAuthService(cs, 2)
	logoutHandler := logouthandler.NewHandler(userService, authService)

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		if item.Cookie != "" {
			req.Header.Set("Cookie", item.Cookie)
		}

		w := httptest.NewRecorder()

		logoutHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseCookie != "" {
			respCookie := resp.Header.Get("Set-Cookie")

			require.Contains(t, respCookie, item.Cookie, utils.TestErrorMessage(caseNum, "Created and received cookie not equal"))
		}
	}
}