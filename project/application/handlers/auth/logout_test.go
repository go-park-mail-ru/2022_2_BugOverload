package auth_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers/auth"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

func TestLogoutHandler(t *testing.T) {
	cases := []TestCase{
		// Success
		TestCase{
			Method:         http.MethodGet,
			Cookie:         "1=YasaPupkinEzji@top.world",
			ResponseBody:   `""`,
			ResponseCookie: "1=YasaPupkinEzji@top.world",
			StatusCode:     http.StatusOK,
		},
		// Cookie has been deleted
		TestCase{
			Method:       http.MethodGet,
			Cookie:       "1=YasaPupkinEzji@top.world",
			ResponseBody: `{"error":"Auth: [no such cookie]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
		// Wrong cookie
		TestCase{
			Method:       http.MethodGet,
			Cookie:       "2=YasaPupkinEzji@top.world",
			ResponseBody: `{"error":"Auth: [no such cookie]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
		// Cookie is missing
		TestCase{
			Method:       http.MethodGet,
			ResponseBody: `{"error":"Auth: [request has no cookies]"}`,
			StatusCode:   http.StatusUnauthorized,
		},
	}

	url := "http://localhost:8088/v1/auth/logout"

	us := database.NewUserStorage()
	user := structs.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	}
	us.Create(user)

	cs := database.NewCookieStorage()
	cs.Create("YasaPupkinEzji@top.world")

	authHandler := auth.NewHandlerAuth(us, cs)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		if item.Cookie != "" {
			req.Header.Set("Cookie", item.Cookie)
		}

		w := httptest.NewRecorder()

		authHandler.Logout(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got [%d], expected [%d]", caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()

		if item.ResponseCookie != "" {
			respCookie := resp.Header.Get("Cookie")
			if strings.HasPrefix(respCookie, item.ResponseCookie) {
				t.Errorf("[%d] wrong cookie: got [%s], cookie must be [%s]", caseNum, respCookie, item.ResponseCookie)
			}
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("[%d] err: [%s], expected: nil", caseNum, err)
		}
		err = resp.Body.Close()
		if err != nil {
			t.Errorf("[%d] err: [%s], expected: nil", caseNum, err)
		}

		bodyStr := string(body)
		if bodyStr != item.ResponseBody {
			t.Errorf("[%d] wrong Response: got [%s], expected [%s]", caseNum, bodyStr, item.ResponseBody)
		}
	}
}
