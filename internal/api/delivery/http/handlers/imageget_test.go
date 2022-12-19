package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	mockImageClient "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
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

	// CheckNotificationSent result
	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode")

	// CheckNotificationSent body
	response := w.Result()

	actual, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	require.Equal(t, res.Bytes, actual, "Wrong body")
}

func TestGetImageHandler_Action_NotOK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodGet, "/api/v1/image?object=connection_poster&key=4", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	input := modelsGlobal.Image{
		Object: "connection_poster",
		Key:    "4",
	}

	res := modelsGlobal.Image{
		Bytes: []byte("some image"),
	}

	// Settings mock
	service.EXPECT().GetImage(r.Context(), &input).Return(res, errors.ErrNotFoundInDB)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetImageHandler(service)
	handler.Configure(router, nil)

	// CheckNotificationSent result
	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode")

	// CheckNotificationSent body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestGetImageHandler_Action_BindError(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodGet, "/api/v1/image?object=connection_poster&key=4", nil)
	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrUnsupportedMediaType.Error(),
	}

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetImageHandler(service)
	handler.Configure(router, nil)

	// CheckNotificationSent result
	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusUnsupportedMediaType, w.Code, "Wrong StatusCode")

	// CheckNotificationSent body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
