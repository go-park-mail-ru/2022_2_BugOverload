package handlers

import (
	"bytes"
	"context"
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
)

// --c833dba25c1bac4ac0a91d0ce81b3a873fa0a49fe912808474b9bd718ab9\r\n
//Content-Disposition: form-data; name=\"object\"; filename=\"some\"\r\n
//Content-Type: application/octet-stream\r\n\r\n
//some image\r\n
//--c833dba25c1bac4ac0a91d0ce81b3a873fa0a49fe912808474b9bd718ab9--\r\n

func TestPutImageHandler_Action_OK(t *testing.T) {
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
	part.Write(imageBin)

	err = writer.Close()
	require.Nil(t, err, "writer.Close() must be success: ", err)

	r := httptest.NewRequest(http.MethodPut, "/api/v1/image?object=user_avatar&key=10", body)

	input := modelsGlobal.Image{
		Object: "user_avatar",
		Key:    "1",
		Bytes:  imageBin,
	}

	//
	r.Header.Set("Content-Type", writer.FormDataContentType())

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	// Settings mock
	service.EXPECT().UpdateImage(ctx, &input).Return(nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPutImageHandler(service)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNoContent, w.Code, "Wrong StatusCode")
}
