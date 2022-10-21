package dependenttest_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	urlNet "net/url"
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

func TestDownloadImageHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:      http.MethodGet,
			Keys:        []string{"default", "test"},
			Values:      []string{"object", "key"},
			RequestBody: `{"object":"default","key":"test"}`,

			StatusCode: http.StatusOK,
		},
		// Not such image
		tests.TestCase{
			Method:      http.MethodGet,
			RequestBody: `{"object":"default","key":"test123"}`,

			StatusCode: http.StatusNotFound,
		},
		// Content-Type is not for get image
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `{"password":"Widget Adapter"}`,

			ResponseBody: `{"error":"Def validation: [unsupported media type]"}`,
			StatusCode:   http.StatusUnsupportedMediaType,
		},
	}

	url := "http://localhost:8088/v1/image"
	config := innerPKG.NewConfig()

	config.S3.Region = "us-east-1"
	config.S3.Endpoint = "http://localhost:4566"
	config.S3.Secret = "bar"
	config.S3.ID = "foo"

	is := S3Image.NewImageS3(config)

	imageService := serviceImage.NewImageService(is)
	getImageHandler := handlers.NewDownloadImageHandler(imageService)

	for caseNum, item := range cases {
		var reader = strings.NewReader(item.RequestBody)

		req := httptest.NewRequest(item.Method, url, reader)
		if item.ContentType != "" {
			req.Header.Set("Content-Type", item.ContentType)
		}

		if len(item.Values) > 0 {
			v := urlNet.Values{}

			for i := range item.Values {
				v.Add(item.Values[i], item.Keys[i])
			}

			req.Form = v
		}

		w := httptest.NewRecorder()

		getImageHandler.Action(w, req)

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseBody != "" {
			resp := w.Result()

			body, err := io.ReadAll(resp.Body)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

			err = resp.Body.Close()
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

			require.Equal(t, item.ResponseBody, string(body), pkg.TestErrorMessage(caseNum, "Wrong body"))
		}
	}
}
