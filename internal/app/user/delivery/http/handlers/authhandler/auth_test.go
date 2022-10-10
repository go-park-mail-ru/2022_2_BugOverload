package authhandler_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/authhandler"
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
			Method:       http.MethodGet,
			Cookie:       "1=YasaPupkinEzji@top.world",
			ResponseBody: `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"asserts/img/invisibleMan.jpeg"}`,
			StatusCode:   http.StatusOK,
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

	url := "http://localhost:8088/v1/auth"

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
	authService := serviceAuth.NewAuthService(us, cs, 2)
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

		body, err := io.ReadAll(resp.Body)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Body.Close must be success"))

		require.Equal(t, item.ResponseBody, string(body), utils.TestErrorMessage(caseNum, "Wrong body"))
	}
}
