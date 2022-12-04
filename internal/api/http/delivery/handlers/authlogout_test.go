package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	mockAuthClient "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
)

func TestAuthLogoutHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodDelete, "/api/v1/auth/logout", nil)
	cookie := &http.Cookie{
		Name:     "session_id",
		Expires:  time.Now().Add(constparams.TimeoutLiveCookie),
		Path:     constparams.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	session := modelsGlobal.Session{
		ID: cookie.Value,
	}

	service.EXPECT().DeleteSession(r.Context(), &session)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewLogoutHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNoContent, w.Code)

	// Check body
	response := w.Result()

	_, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")
}
