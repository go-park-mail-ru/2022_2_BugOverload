package handlers

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	authMock "go-park-mail-ru/2022_2_BugOverload/internal/auth/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/models"
	sessionMock "go-park-mail-ru/2022_2_BugOverload/internal/session/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler_Action(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := authMock.NewMockAuthService(ctrl)
	sessService := sessionMock.NewMockSessionService(ctrl)

	data := bytes.NewReader(
		[]byte(`{
			"email": "YasaPupkinEzji@top.world",
			"password": "Widget Adapter"
			}`),
	)

	r := httptest.NewRequest("POST", "/api/v1/login", data)

	authService.EXPECT().Login(r.Context(), &models.User{
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
	}).Return(models.User{
		Nickname: "StepByyyy",
		Email:    "YasaPupkinEzji@top.world",
	}, nil)

	sessService.EXPECT().CreateSession(r.Context(), models.User{
		Nickname: "StepByyyy",
		Email:    "YasaPupkinEzji@top.world",
	}).Return(models.Session{
		ID: "session",
	}, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	loginHandler := NewLoginHandler(authService, sessService)
	loginHandler.Configure(router, nil)

	loginHandler.Action(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, len(w.Header().Get("X-Csrf-Token")) > 0)
}
