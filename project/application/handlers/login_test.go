package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// TestCase is structure for API testing
type TestCase struct {
	Method       string
	RequestBody  string
	ResponseBody string
	ContentType  string
	StatusCode   int
}

func TestLoginHandler(t *testing.T) {
	cases := []TestCase{
		//  Success
		TestCase{
			Method:       http.MethodPost,
			ContentType:  "application/json",
			RequestBody:  `{"email":"YasaPupkinEzji@top.world","password":"Widget Adapter"}`,
			ResponseBody: `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"URL"}`,
			StatusCode:   http.StatusOK,
		},
		//  BrokenJSON
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email": 123, "password": "Widget Adapter"`,

			ResponseBody: "unexpected end of JSON input\n",
			StatusCode:   http.StatusBadRequest,
		},
		////  Method not POST
		//TestCase{
		//	Method:      http.MethodGet,
		//	ContentType: "application/json",
		//	RequestBody: `{"email":"YasaPupkinEzji@top.world":"password":"Widget Adapter"}`,
		//	StatusCode:  http.StatusMethodNotAllowed,
		//},
		////  Body is empty
		//TestCase{
		//	Method:      http.MethodPost,
		//	ContentType: "application/json",
		//
		//	ResponseBody: "BadRequest\nContent-Type must be application/json\n",
		//	StatusCode:   http.StatusBadRequest,
		//},
		//// body not JSON
		//TestCase{
		//	Method:      http.MethodPost,
		//	ContentType: "application/xml",
		//	RequestBody: `<Name>Ellen Adams</Name>`,
		//
		//	ResponseBody: "Unsupported Media Type\nContent-Type must be application/json\n",
		//	StatusCode:   http.StatusUnsupportedMediaType,
		//},
	}

	url := "http://localhost:8088/v1/auth/login"

	us := database.NewUserStorage()
	us.Create(structs.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	})

	authHandler := handlers.NewHandlerAuth(us)

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
