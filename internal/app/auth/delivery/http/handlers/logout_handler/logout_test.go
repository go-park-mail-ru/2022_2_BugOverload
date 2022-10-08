package logout_handler_test

import (
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/delivery/http/handlers/logout_handler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/database"
)

// TestCase is structure for API testing
type TestCase struct {
	Method          string
	ContentType     string
	RequestBody     string
	CookieUserEmail string
	Cookie          string
	ResponseCookie  string
	ResponseBody    string
	StatusCode      int
}

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

	url := "http://localhost:8088/v1/auth/logout"

	us := database.NewUserStorage()
	user := models.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	}
	us.Create(user)

	cs := database.NewCookieStorage()
	cs.Create("YasaPupkinEzji@top.world")

	logoutHandler := logout_handler.NewHandler(us, cs)

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

		logoutHandler.Action(w, req)

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
