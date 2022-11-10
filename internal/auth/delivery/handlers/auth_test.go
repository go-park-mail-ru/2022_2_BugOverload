package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	authMock "go-park-mail-ru/2022_2_BugOverload/internal/auth/mocks"
	sessionMock "go-park-mail-ru/2022_2_BugOverload/internal/session/mocks"
)

func TestAuthHandler_AuthSuccess(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := authMock.NewMockAuthService(ctrl)
	sessService := sessionMock.NewMockSessionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/auth", nil)

	authService.EXPECT().Auth(r.Context(), &models.User{
		ID: 1,
	}).Return(models.User{
		Nickname: "StepByyyy",
		Email:    "YasaPupkinEzji@top.world",
		Profile: models.Profile{
			Avatar: "avatar",
		},
	}, nil)

	sessService.EXPECT().GetUserBySession(r.Context(), models.Session{
		ID: "c9QuR4KQR4RkXi_rbATHWITwQGDG9r801tHIA_AHkDt2JNiVWU8Tjg==",
	}).Return(models.User{
		ID: 1,
	}, nil)

	cookie := &http.Cookie{
		Name:     pkg.SessionCookieName,
		Value:    "c9QuR4KQR4RkXi_rbATHWITwQGDG9r801tHIA_AHkDt2JNiVWU8Tjg==",
		Expires:  time.Now().Add(pkg.TimeoutLiveCookie),
		Path:     pkg.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	w := httptest.NewRecorder()

	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	authHandler := NewAuthHandler(authService, sessService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, len(w.Header().Get("X-Csrf-Token")) > 0)
}

func TestAuthHandler_AuthWithoutCookie(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := authMock.NewMockAuthService(ctrl)
	sessService := sessionMock.NewMockSessionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/auth", nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	authHandler := NewAuthHandler(authService, sessService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthHandler_AuthWithInvalidCookie(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := authMock.NewMockAuthService(ctrl)
	sessService := sessionMock.NewMockSessionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/auth", nil)

	cookie := &http.Cookie{
		Expires:  time.Now().Add(pkg.TimeoutLiveCookie),
		Path:     pkg.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	w := httptest.NewRecorder()

	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	authHandler := NewAuthHandler(authService, sessService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestAuthHandler_AuthFallAuthService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := authMock.NewMockAuthService(ctrl)
	sessService := sessionMock.NewMockSessionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/auth", nil)

	cookie := &http.Cookie{
		Name:     "session_id",
		Expires:  time.Now().Add(pkg.TimeoutLiveCookie),
		Path:     pkg.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	oldlogger := logrus.New()
	logger := logrus.NewEntry(oldlogger)

	ctx := context.WithValue(r.Context(), pkg.LoggerKey, logger)

	sessService.EXPECT().GetUserBySession(context.WithValue(r.Context(), pkg.LoggerKey, logger), models.Session{ID: ""}).Return(models.User{
		ID: 0,
	}, nil)
	authService.EXPECT().Auth(context.WithValue(r.Context(), pkg.LoggerKey, logger), &models.User{
		ID: 0,
	}).Return(models.User{}, errors.ErrPostgresRequest)

	w := httptest.NewRecorder()

	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	authHandler := NewAuthHandler(authService, sessService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r.WithContext(ctx))
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAuthHandler_AuthFallSessionService(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := authMock.NewMockAuthService(ctrl)
	sessService := sessionMock.NewMockSessionService(ctrl)

	r := httptest.NewRequest("GET", "/api/v1/auth", nil)

	cookie := &http.Cookie{
		Name:     "session_id",
		Expires:  time.Now().Add(pkg.TimeoutLiveCookie),
		Path:     pkg.GlobalCookiePath,
		HttpOnly: true,
	}

	r.AddCookie(cookie)

	sessService.EXPECT().GetUserBySession(r.Context(), models.Session{ID: ""}).Return(models.User{}, errors.ErrSessionNotExist)

	w := httptest.NewRecorder()

	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	authHandler := NewAuthHandler(authService, sessService)
	authHandler.Configure(router, nil)

	authHandler.Action(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
