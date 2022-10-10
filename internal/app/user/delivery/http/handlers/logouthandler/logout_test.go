package logouthandler_test

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/logouthandler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
)

// TestCase is structure for API testing
type TestCase struct {
	Method         string
	ContentType    string
	RequestBody    string
	Cookie         string
	ResponseCookie string
	ResponseBody   string
	StatusCode     int
}

func TestAuthHandler(t *testing.T) {
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

	us := memoryUser.NewUserRepo()
	cs := memoryCookie.NewCookieRepo()

	testUser := &models.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	}

	us.CreateUser(context.TODO(), testUser)
	cs.CreateSession(context.TODO(), testUser)

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

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))
	}
}
