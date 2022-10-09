package signuphandler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/user/delivery/http/handlers/signuphandler"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	serviceUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/service"
)

// TestCase is structure for API testing
type TestCase struct {
	Method         string
	ContentType    string
	RequestBody    string
	User           models.User
	Cookie         string
	ResponseCookie string
	ResponseBody   string
	StatusCode     int
}

func TestSignupHandler(t *testing.T) {
	cases := []TestCase{
		// Success
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,
			User: models.User{
				Email: "YasaPupkinEzji@top.world",
			},

			ResponseCookie: "1=YasaPupkinEzji@top.world",
			ResponseBody:   `{"nickname":"testnickname","email":"testmail@yandex.ru"}`,
			StatusCode:     http.StatusCreated,
		},
		// Such user exists
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Action: [such a login exists]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Broken JSON
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","password":"testpassword"`,

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body not JSON
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/xml",
			RequestBody: `<Name>Ellen Adams</Name>`,

			ResponseBody: `{"error":"Def validation: [unsupported media type]"}`,
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Empty required field - email
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - password
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Empty required field - nickname
		TestCase{
			Method:      http.MethodPost,
			ContentType: "application/json",
			RequestBody: `{"email":"testmail@yandex.ru","password": "testpassword"}`,

			ResponseBody: `{"error":"Action: [request has empty fields (nickname | email | password)]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Content-Type not set
		TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"email":"testmail@yandex.ru","nickname":"testnickname","password": "testpassword"}`,

			ResponseBody: `{"error":"Def validation: [content-type undefined]"}`,
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/auth/signup"

	us := memoryUser.NewUserRepo()
	cs := memoryCookie.NewCookieRepo()

	serviceUser.NewUserService(us, 2)
	serviceAuth.NewAuthService(cs, 2)
	authHandler := signuphandler.NewHandler(us, cs)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}
		w := httptest.NewRecorder()

		authHandler.Action(w, req)

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got [%d], expected [%d]", caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()

		if item.ResponseCookie != "" {
			respCookie := resp.Header.Get("Set-Cookie")

			fullCookieStr := cs.CreateSession(item.CookieUserEmail)

			if strings.HasPrefix(fullCookieStr, item.ResponseCookie) {
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
			t.Errorf("[%d] wrong Response: got [%+v], expected [%+v]", caseNum, bodyStr, item.ResponseBody)
		}
	}
}
