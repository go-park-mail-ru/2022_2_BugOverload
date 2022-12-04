package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	mockImageClient "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
)

func TestGetImageHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodGet, "/api/v1/image?object=connection_poster&key=4", nil)

	input := modelsGlobal.Image{
		Object: "connection_poster",
		Key:    "4",
	}

	res := modelsGlobal.Image{
		Bytes: []byte("some image"),
	}

	// Settings mock
	service.EXPECT().GetImage(r.Context(), &input).Return(res, nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetImageHandler(service)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	actual, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	require.Equal(t, res.Bytes, actual, "Wrong body")
}
