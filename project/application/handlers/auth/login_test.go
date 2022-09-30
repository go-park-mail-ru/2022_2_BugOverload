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

// TestCase is structure for API testing
type TestCase struct {
	Method         string
	ContentType    string
	RequestBody    string
	ResponseCookie string
	ResponseBody   string
	StatusCode     int
}

func TestLoginHandler(t *testing.T) {
	cases := []TestCase{
		// Success
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"YasaPupkinEzji@top.world","password":"Widget Adapter"}`,

			ResponseCookie: "1" + "YasaPupkinEzji@top.world",
			ResponseBody:   `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"URL"}`,
			StatusCode:     http.StatusOK,
		},
		// Wrong password
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"YasaPupkinEzji@top.world","password":"Widget 123123123Adapter"}`,

			ResponseBody: "no such combination of user and password\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Broken JSON
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email": 123, "password": "Widget Adapter"`,

			ResponseBody: "unexpected end of JSON input\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",

			ResponseBody: "unexpected end of JSON input\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Body not JSON
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/xml",
			RequestBody: `<Name>Ellen Adams</Name>`,

			ResponseBody: "unsupported media type\n",
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Empty required field - email
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: "request has empty fields (nickname | email | password)\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"YasaPupkinEzji@top.world"}`,

			ResponseBody: "request has empty fields (nickname | email | password)\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: "content-type undefined\n",
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/login"

	us := database.NewUserStorage()
	user := structs.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	}
	us.Create(user)

	cs := database.NewCookieStorage()
	cs.Create(user.Email)

	authHandler := auth.NewHandlerAuth(us, cs)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		w := httptest.NewRecorder()

		authHandler.Login(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got [%d], expected [%d]", caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()

		if item.ResponseCookie != "" {
			respCookie := resp.Header.Get("Cookie")

			if strings.HasPrefix(respCookie, item.ResponseCookie) {
				t.Errorf("[%d] wrong cookie: got [%s], cookie must have [%s]", caseNum, respCookie, item.ResponseCookie)
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
