package dependenttest_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	urlNet "net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	"go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/handlers"
	S3Image "go-park-mail-ru/2022_2_BugOverload/internal/image/repository"
	serviceImage "go-park-mail-ru/2022_2_BugOverload/internal/image/service"
	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestUploadImageHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:      http.MethodGet,
			ContentType: innerPKG.ContentTypeJPEG,
			RequestBody: `image`,
			Keys:        []string{"default", "test123123"},
			Values:      []string{"object", "key"},

			StatusCode: http.StatusNoContent,
		},
		// Content-Type is not for get image
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJSON,
			RequestBody: `image`,

			StatusCode: http.StatusUnsupportedMediaType,
		},
		// Body is empty
		tests.TestCase{
			Method:      http.MethodPost,
			ContentType: innerPKG.ContentTypeJPEG,

			StatusCode: http.StatusBadRequest,
		},
	}

	file, err := os.Open("testup.jpeg")
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err open file for test"))

	buf := make([]byte, innerPKG.BufSizeImage)
	numBytes, err := file.Read(buf)
	require.Nil(t, err, pkg.TestErrorMessage(-1, "Err read file for test"))

	buf = buf[:numBytes]

	url := "http://localhost:8088/v1/image"
	config := innerPKG.NewConfig()

	config.S3.Region = "us-east-1"
	config.S3.Endpoint = "http://localhost:4566"
	config.S3.Secret = "bar"
	config.S3.ID = "foo"

	is := S3Image.NewImageS3(config)

	imageService := serviceImage.NewImageService(is)
	getImageHandler := handlers.NewUploadImageHandler(imageService)

	for caseNum, item := range cases {
		var reader = bytes.NewReader([]byte{})

		if item.RequestBody != "" {
			reader.Reset(buf)
		}

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

			var body []byte

			body, err = io.ReadAll(resp.Body)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

			err = resp.Body.Close()
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

			require.Equal(t, item.ResponseBody, string(body), pkg.TestErrorMessage(caseNum, "Wrong body"))
		}
	}
}
