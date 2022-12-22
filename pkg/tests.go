package pkg

import (
	"io"
	"net/http/httptest"

	"github.com/sirupsen/logrus"
)

func GetResponseBody(w httptest.ResponseRecorder) string {
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	if errResponse != nil {
		logrus.Fatal(errResponse.Error())
	}

	err := response.Body.Close()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return string(bodyResponse)
}
