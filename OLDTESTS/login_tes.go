package OLDTESTS_test

//
//import (
//	memory2 "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
//	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//
//	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
//	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/loginhandler"
//)
//
//// TestCase is structure for API testing
//type TestCase struct {
//	Method          string
//	ContentType     string
//	RequestBody     string
//	CookieUserEmail string
//	Cookie          string
//	ResponseCookie  string
//	ResponseBody    string
//	StatusCode      int
//}
//
//func TestLoginHandler(t *testing.T) {
//	cases := []TestCase{
//		// Success
//		TestCase{
//			Method:          http.MethodPost,
//			ContentType:     "application/json",
//			RequestBody:     `{"email":"YasaPupkinEzji@top.world","password":"Widget Adapter"}`,
//			CookieUserEmail: "YasaPupkinEzji@top.world",
//
//			ResponseCookie: "1=YasaPupkinEzji@top.world",
//			ResponseBody:   `{"nickname":"Andeo","email":"YasaPupkinEzji@top.world","avatar":"asserts/img/invisibleMan.jpeg"}`,
//			StatusCode:     http.StatusOK,
//		},
//		// Wrong password
//		TestCase{
//			Method:      http.MethodPost,
//			ContentType: "application/json",
//			RequestBody: `{"email":"YasaPupkinEzji@top.world","password":"Widget 123123123Adapter"}`,
//
//			ResponseBody: `{"error":"Action: [no such combination of login and password]"}`,
//			StatusCode:   http.StatusUnauthorized,
//		},
//		// Broken JSON
//		TestCase{
//			Method:      http.MethodPost,
//			ContentType: "application/json",
//			RequestBody: `{"email": 123, "password": "Widget Adapter"`,
//
//			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
//			StatusCode:   http.StatusBadRequest,
//		},
//		// Body is empty
//		TestCase{
//			Method:      http.MethodPost,
//			ContentType: "application/json",
//
//			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
//			StatusCode:   http.StatusBadRequest,
//		},
//		// Body not JSON
//		TestCase{
//			Method:      http.MethodPost,
//			ContentType: "application/xml",
//			RequestBody: `<Name>Ellen Adams</Name>`,
//
//			ResponseBody: `{"error":"Def validation: [unsupported media type]"}`,
//			StatusCode:   http.StatusUnsupportedMediaType,
//		},
//		// Empty required field - email
//		TestCase{
//			Method:      http.MethodPost,
//			ContentType: "application/json",
//			RequestBody: `{"password":"Widget Adapter"}`,
//
//			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
//			StatusCode:   http.StatusBadRequest,
//		},
//		// Empty required field - password
//		TestCase{
//			Method:      http.MethodPost,
//			ContentType: "application/json",
//			RequestBody: `{"email":"YasaPupkinEzji@top.world"}`,
//
//			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
//			StatusCode:   http.StatusBadRequest,
//		},
//		// Content-Type not set
//		TestCase{
//			Method:      http.MethodPost,
//			RequestBody: `{"password":"Widget Adapter"}`,
//
//			ResponseBody: `{"error":"Def validation: [content-type undefined]"}`,
//			StatusCode:   http.StatusBadRequest,
//		},
//	}
//
//	url := "http://localhost:8088/v1/auth/login"
//
//	us := memory.NewUserRepo()
//	user := models.User{
//		Nickname: "Andeo",
//		Email:    "YasaPupkinEzji@top.world",
//		Password: "Widget Adapter",
//		Avatar:   "URL",
//	}
//	us.Signup(user)
//
//	cs := memory2.NewCookieRepo()
//
//	authHandler := loginhandler.NewHandler(us, cs)
//
//	for caseNum, item := range cases {
//		var reader = strings.NewReader(item.RequestBody)
//
//		req := httptest.NewRequest(item.Method, url, reader)
//		if item.ContentType != "" {
//			req.Header.Set("Content-Type", item.ContentType)
//		}
//
//		w := httptest.NewRecorder()
//
//		authHandler.Action(w, req)
//
//		if w.Code != item.StatusCode {
//			t.Errorf("[%d] wrong StatusCode: got [%d], expected [%d]", caseNum, w.Code, item.StatusCode)
//		}
//
//		resp := w.Result()
//
//		if item.ResponseCookie != "" {
//			respCookie := resp.Header.Get("Set-Cookie")
//
//			fullCookieStr := cs.CreateSession(item.CookieUserEmail)
//
//			if strings.HasPrefix(fullCookieStr, item.ResponseCookie) {
//				t.Errorf("[%d] wrong cookie: got [%s], cookie must be [%s]", caseNum, respCookie, item.ResponseCookie)
//			}
//		}
//
//		body, err := io.ReadAll(resp.Body)
//		if err != nil {
//			t.Errorf("[%d] err: [%s], expected: nil", caseNum, err)
//		}
//		err = resp.Body.Close()
//		if err != nil {
//			t.Errorf("[%d] err: [%s], expected: nil", caseNum, err)
//		}
//
//		bodyStr := string(body)
//		if bodyStr != item.ResponseBody {
//			t.Errorf("[%d] wrong Response: got [%s], expected [%s]", caseNum, bodyStr, item.ResponseBody)
//		}
//	}
//}
