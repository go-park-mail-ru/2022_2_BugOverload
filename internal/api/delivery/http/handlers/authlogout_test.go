package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	mockAuthClient "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
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

	// CheckNotificationSent code
	require.Equal(t, http.StatusNoContent, w.Code)

	// CheckNotificationSent body
	response := w.Result()

	_, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")
}

func TestAuthLogoutHandler_Action_CookieError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodDelete, "/api/v1/auth/logout", nil)
	cookie := &http.Cookie{
		Name:     "invalid_cookie_name",
		Expires:  time.Now().Add(constparams.TimeoutLiveCookie),
		Path:     constparams.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	_, errCookie := r.Cookie(constparams.SessionCookieName)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: stdErrors.Wrap(errCookie, "Undefined error").Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewLogoutHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusInternalServerError, w.Code)

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

func TestAuthLogoutHandler_Action_SessionError(t *testing.T) {
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

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrSessionNotFound.Error(),
	}

	session := modelsGlobal.Session{
		ID: cookie.Value,
	}

	service.EXPECT().DeleteSession(r.Context(), &session).Return(modelsGlobal.Session{}, errors.ErrSessionNotFound)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewLogoutHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusNotFound, w.Code)

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
