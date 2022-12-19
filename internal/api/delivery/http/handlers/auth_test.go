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
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	mockAuthClient "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

func TestAuthHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/auth", nil)

	cookie := &http.Cookie{
		Name:     constparams.SessionCookieName,
		Value:    "c9QuR4KQR4RkXi_rbATHWITwQGDG9r801tHIA_AHkDt2JNiVWU8Tjg==",
		Expires:  time.Now().Add(constparams.TimeoutLiveCookie),
		Path:     constparams.GlobalCookiePath,
		HttpOnly: true,
	}

	resAuth := modelsGlobal.User{
		Nickname: "StepByyyy",
		Email:    "YasaPupkinEzji@top.world",
		Avatar:   "avatar",
	}

	service.EXPECT().GetUserBySession(r.Context(), &modelsGlobal.Session{
		ID: cookie.Value,
	}).Return(modelsGlobal.User{
		ID: 1,
	}, nil)

	service.EXPECT().Auth(r.Context(), &modelsGlobal.User{
		ID: 1,
	}).Return(resAuth, nil)

	r.AddCookie(cookie)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewAuthHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusOK, w.Code)

	// CheckNotificationSent Needed headers
	require.True(t, len(w.Header().Get("X-Csrf-Token")) > 0)

	// CheckNotificationSent body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewUserAuthResponse(&resAuth)

	var actualBody *models.UserAuthResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestAuthHandler_AuthWithoutCookie(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/auth", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNoCookie.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	authHandler := NewAuthHandler(authService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r)

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

func TestAuthHandler_AuthWithInvalidCookie(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/auth", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNoCookie.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	cookie := &http.Cookie{
		Expires:  time.Now().Add(constparams.TimeoutLiveCookie),
		Path:     constparams.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	authHandler := NewAuthHandler(authService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r)

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

func TestAuthHandler_Action_NotOK_AuthService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/auth", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrWorkDatabase.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	cookie := &http.Cookie{
		Name:     "session_id",
		Expires:  time.Now().Add(constparams.TimeoutLiveCookie),
		Path:     constparams.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	authService.EXPECT().GetUserBySession(r.Context(), &modelsGlobal.Session{}).Return(modelsGlobal.User{
		ID: 1,
	}, nil)

	authService.EXPECT().Auth(r.Context(), &modelsGlobal.User{
		ID: 1,
	}).Return(modelsGlobal.User{}, errors.ErrWorkDatabase)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	authHandler := NewAuthHandler(authService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r.WithContext(ctx))

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

func TestAuthHandler_Action_NotOK_SessionService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/auth", nil)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrSessionNotFound.Error(),
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	cookie := &http.Cookie{
		Name:     "session_id",
		Expires:  time.Now().Add(constparams.TimeoutLiveCookie),
		Path:     constparams.GlobalCookiePath,
		HttpOnly: true,
		Value:    "cryptorand",
	}

	r.AddCookie(cookie)

	authService.EXPECT().GetUserBySession(r.Context(), &modelsGlobal.Session{ID: cookie.Value}).Return(modelsGlobal.User{
		ID: 1,
	}, errors.ErrSessionNotFound)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	authHandler := NewAuthHandler(authService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r.WithContext(ctx))

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
