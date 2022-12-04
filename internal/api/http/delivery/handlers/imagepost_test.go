package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	mockImageClient "go-park-mail-ru/2022_2_BugOverload/internal/image/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

// Multipart body
// Content-Type: multipart/form-data; boundary=---WebKitFormBoundary7MA4YWxkTrZu0gW
//
//    -----WebKitFormBoundary7MA4YWxkTrZu0gW
//    Content-Disposition: form-data; name=”file”; filename=”captcha”
//    Content-Type:
//
//    -----WebKitFormBoundary7MA4YWxkTrZu0gW
//    Content-Disposition: form-data; name=”action”
//
//    submit
//    -----WebKitFormBoundary7MA4YWxkTrZu0gW--

func TestPostImageHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	imageBin := []byte("some image")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Metadata part
	metadataHeader := textproto.MIMEHeader{}

	// Set the Content-Type header
	metadataHeader.Add("Content-Disposition", `form-data; name="object"; filename="some.jpeg"`)
	metadataHeader.Add("Content-Type", "image/jpeg")

	part, err := writer.CreatePart(metadataHeader)
	require.Nil(t, err, "Body.Close must be success")
	_, err = part.Write(imageBin)
	require.Nil(t, err, "part.Write(imageBin) must be success: ", err)

	err = writer.Close()
	require.Nil(t, err, "writer.Close() must be success: ", err)

	r := httptest.NewRequest(http.MethodPut, "/api/v1/image?object=user_avatar&key=10", body)

	input := modelsGlobal.Image{
		Object: "user_avatar",
		Key:    "10",
		Bytes:  imageBin,
	}

	r.Header.Set("Content-Type", writer.FormDataContentType())

	user := modelsGlobal.User{
		ID:      1,
		IsAdmin: true,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	// Settings mock
	service.EXPECT().UpdateImage(ctx, &input).Return(nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPostImageHandler(service)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusCreated, w.Code, "Wrong StatusCode")
}

func TestPostImageHandler_Action_NotOk(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	imageBin := []byte("some image")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Metadata part
	metadataHeader := textproto.MIMEHeader{}

	// Set the Content-Type header
	metadataHeader.Add("Content-Disposition", `form-data; name="object"; filename="some.jpeg"`)
	metadataHeader.Add("Content-Type", "image/jpeg")

	part, err := writer.CreatePart(metadataHeader)
	require.Nil(t, err, "Body.Close must be success")
	_, err = part.Write(imageBin)
	require.Nil(t, err, "part.Write(imageBin) must be success: ", err)

	err = writer.Close()
	require.Nil(t, err, "writer.Close() must be success: ", err)

	r := httptest.NewRequest(http.MethodPut, "/api/v1/image?object=user_avatar&key=10", body)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	input := modelsGlobal.Image{
		Object: "user_avatar",
		Key:    "10",
		Bytes:  imageBin,
	}

	r.Header.Set("Content-Type", writer.FormDataContentType())

	user := modelsGlobal.User{
		ID:      1,
		IsAdmin: true,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	// Settings mock
	service.EXPECT().UpdateImage(ctx, &input).Return(errors.ErrNotFoundInDB)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPostImageHandler(service)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	bodyResp, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResp, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestPostImageHandler_Action_Object_ContetUndef(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	imageBin := []byte("some image")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Metadata part
	metadataHeader := textproto.MIMEHeader{}

	// Set the Content-Type header
	metadataHeader.Add("Content-Disposition", `form-data; name="object"; filename="some.jpeg"`)
	metadataHeader.Add("Content-Type", "image/jpeg")

	part, err := writer.CreatePart(metadataHeader)
	require.Nil(t, err, "Body.Close must be success")
	_, err = part.Write(imageBin)
	require.Nil(t, err, "part.Write(imageBin) must be success: ", err)

	err = writer.Close()
	require.Nil(t, err, "writer.Close() must be success: ", err)

	r := httptest.NewRequest(http.MethodPut, "/api/v1/image?object=user_avatar&key=10", body)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrContentTypeUndefined.Error(),
	}

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPostImageHandler(service)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	bodyResp, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResp, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestPostImageHandler_Action_UserNotFound(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	imageBin := []byte("some image")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Metadata part
	metadataHeader := textproto.MIMEHeader{}

	// Set the Content-Type header
	metadataHeader.Add("Content-Disposition", `form-data; name="object"; filename="some.jpeg"`)
	metadataHeader.Add("Content-Type", "image/jpeg")

	part, err := writer.CreatePart(metadataHeader)
	require.Nil(t, err, "Body.Close must be success")
	_, err = part.Write(imageBin)
	require.Nil(t, err, "part.Write(imageBin) must be success: ", err)

	err = writer.Close()
	require.Nil(t, err, "writer.Close() must be success: ", err)

	r := httptest.NewRequest(http.MethodPut, "/api/v1/image?object=user_avatar&key=10", body)

	r.Header.Set("Content-Type", writer.FormDataContentType())

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPostImageHandler(service)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusInternalServerError, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	require.Nil(t, errResponse, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrGetUserRequest.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestPostImageHandler_Action_UserNotAdmin(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockImageClient.NewMockImageService(ctrl)

	// Data
	imageBin := []byte("some image")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Metadata part
	metadataHeader := textproto.MIMEHeader{}

	// Set the Content-Type header
	metadataHeader.Add("Content-Disposition", `form-data; name="object"; filename="some.jpeg"`)
	metadataHeader.Add("Content-Type", "image/jpeg")

	part, err := writer.CreatePart(metadataHeader)
	require.Nil(t, err, "Body.Close must be success")
	_, err = part.Write(imageBin)
	require.Nil(t, err, "part.Write(imageBin) must be success: ", err)

	err = writer.Close()
	require.Nil(t, err, "writer.Close() must be success: ", err)

	r := httptest.NewRequest(http.MethodPut, "/api/v1/image?object=user_avatar&key=10", body)

	r.Header.Set("Content-Type", writer.FormDataContentType())

	user := modelsGlobal.User{
		ID:      1,
		IsAdmin: false,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPostImageHandler(service)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusInternalServerError, w.Code, "Wrong StatusCode")

	// Check body
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	require.Nil(t, errResponse, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrGetUserRequest.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
