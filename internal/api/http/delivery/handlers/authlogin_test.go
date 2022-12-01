package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	mockAuthClient "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
)

func TestAuthLoginHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authService := mockAuthClient.NewMockAuthService(ctrl)

	mcPostBody := map[string]string{
		"email":    "YasaPupkinEzji@top.world",
		"password": "Widget Adapter",
	}

	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))

	r.Header.Set("Content-Type", "application/json")

	resLogin := modelsGlobal.User{
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
	}

	resSession := modelsGlobal.Session{
		ID:   "1",
		User: &resLogin,
	}

	authService.EXPECT().Login(r.Context(), &modelsGlobal.User{
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
	}).Return(resLogin, nil)

	authService.EXPECT().CreateSession(r.Context(), &resLogin).Return(resSession, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	loginHandler := NewLoginHandler(authService)
	loginHandler.Configure(router, nil)

	loginHandler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code)

	// Check Needed headers
	require.True(t, len(w.Header().Get("X-Csrf-Token")) > 0)

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewUserLoginResponse(&resLogin)

	var actualBody *models.UserLoginResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
