package handlers_test

import (
	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	cases := []TestCase{
		// Success
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			RequestBody:  `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,
			ResponseBody: `{"nickname":"testnickname","email":"testmail@yandex.ru"}`,
			StatusCode:   http.StatusCreated,
		},
		// Such user exists
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			RequestBody:  `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,
			ResponseBody: "A user with such a mail already exists\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Broken JSON
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			RequestBody:  `{"email":"testmail@yandex.ru","password":"testpassword"`,
			ResponseBody: "unexpected end of JSON input\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			ResponseBody: "unexpected end of JSON input\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Body not JSON
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/xml",
			RequestBody:  `<Name>Ellen Adams</Name>`,
			ResponseBody: "unsupported media type\n",
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Empty required field - email
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			RequestBody:  `{"nickname":"testnickname","password": "testpassword"}`,
			ResponseBody: "request has empty fields (nickname | email | password)\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			RequestBody:  `{"email":"testmail@yandex.ru","nickname":"testnickname"}`,
			ResponseBody: "request has empty fields (nickname | email | password)\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - nickname
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			RequestBody:  `{"email":"testmail@yandex.ru","password": "testpassword"}`,
			ResponseBody: "request has empty fields (nickname | email | password)\n",
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		TestCase{
			Method:       http.MethodPost,
			RequestBody:  `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,
			ResponseBody: "content-type undefined\n",
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/signup"

	us := database.NewUserStorage()

	authHandler := handlers.NewHandlerAuth(us)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}
		w := httptest.NewRecorder()

		authHandler.Signup(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got [%d], expected [%d]", caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()

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
			t.Errorf("[%d] wrong Response: got [%+v], expected [%+v]", caseNum, bodyStr, item.ResponseBody)
		}
	}
}
