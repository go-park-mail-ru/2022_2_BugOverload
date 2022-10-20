package dependenttest_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/handlers"
	S3Image "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestGetImageHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:      http.MethodGet,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"bucket":"default/","item":"test.jpeg"}`,

			ResponseBody: "GeneratedData",
			StatusCode:   http.StatusOK,
		},
		// Not such image
		tests.TestCase{
			Method:      http.MethodGet,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"bucket":"default/","item":"test.jpeg123"}`,

			ResponseBody: `{"error":"Auth: [no such combination of login and password]"}`,
			StatusCode:   http.StatusNotFound,
		},
		// Broken JSON
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"email": 123, "password": "Widget Adapter"`,

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body is empty
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,

			ResponseBody: `{"error":"Def validation: [unexpected end of JSON input]"}`,
			StatusCode:   http.StatusBadRequest,
		},
		// Body not image/jpeg
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: "application/xml",
			RequestBody: `<Name>Ellen Adams</Name>`,

			ResponseBody: `{"error":"Def validation: [unsupported media type]"}`,
			StatusCode:   http.StatusUnsupportedMediaType,
		},
		// Content-Type not set
		tests.TestCase{
			Method:      http.MethodPost,
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: `{"error":"Def validation: [content-type undefined]"}`,
			StatusCode:   http.StatusBadRequest,
		},
	}

	url := "http://localhost:8088/v1/static"
	config := innerPKG.NewConfig()

	config.S3.Region = "us-east-1"
	config.S3.Endpoint = "http://localhost:4566"
	config.S3.Secret = "bar"
	config.S3.ID = "foo"

	is, err := S3Image.NewImageS3(config)
	require.Nil(t, err, pkg.TestErrorMessage(-1, "NewImageS3 must be success"))

	imageService := serviceImage.NewImageService(is)
	getImageHandler := handlers.NewGetImageHandler(imageService)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		w := httptest.NewRecorder()

		getImageHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		var body []byte
		body, err = io.ReadAll(resp.Body)
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))
		require.NotNil(t, body, pkg.TestErrorMessage(caseNum, "body must be not nil"))

		err = resp.Body.Close()
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))
	}
}
