package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	mockAuthClient "go-park-mail-ru/2022_2_BugOverload/internal/auth/delivery/grpc/client/mocks"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

func TestAuthSignUpHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockAuthClient.NewMockAuthService(ctrl)

	mcPostBody := map[string]string{
		"nickname": "StepByyyy",
		"email":    "YasaPupkinEzji@top.world",
		"password": "Widget Adapter",
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))

	r.Header.Set("Content-Type", "application/json")

	resSignup := modelsGlobal.User{
		Email:    "YasaPupkinEzji@top.world",
		Nickname: "StepByyyy",
		Avatar:   "123",
	}

	resSession := modelsGlobal.Session{
		ID:   "1",
		User: &resSignup,
	}

	service.EXPECT().Signup(r.Context(), &modelsGlobal.User{
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Nickname: "StepByyyy",
	}).Return(resSignup, nil)

	service.EXPECT().CreateSession(r.Context(), &resSignup).Return(resSession, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSingUpHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// check code
	require.Equal(t, http.StatusCreated, w.Code)

	// check Needed headers
	require.True(t, len(w.Header().Get("X-Csrf-Token")) > 0)

	// check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewUserSignUpResponse(&resSignup)

	var actualBody *models.UserSignupResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestAuthSignUpHandler_Action_InvBody(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockAuthClient.NewMockAuthService(ctrl)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)

	r.Header.Set("Content-Type", "application/json")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrEmptyBody.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSingUpHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// check code
	require.Equal(t, http.StatusBadRequest, w.Code)

	// check body
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

func TestAuthSignupHandler_Action_ServiceSignupError(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	mcPostBody := map[string]string{
		"email":    "YasaPupkinEzji@top.world",
		"password": "Widget Adapter",
		"nickname": "StepByyyy",
	}

	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewReader(body))

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	inputUser := modelsGlobal.User{
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Nickname: "StepByyyy",
	}

	service.EXPECT().Signup(r.Context(), &inputUser).Return(modelsGlobal.User{}, errors.ErrUserExist)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrUserExist.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSingUpHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// check code
	require.Equal(t, http.StatusBadRequest, w.Code)

	// check body
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

func TestAuthSignupHandler_Action_CreateSessionError(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockAuthClient.NewMockAuthService(ctrl)

	// Data
	mcPostBody := map[string]string{
		"email":    "YasaPupkinEzji@top.world",
		"password": "Widget Adapter",
		"nickname": "StepByyyy",
	}

	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	inputUser := modelsGlobal.User{
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Nickname: "StepByyyy",
	}

	outputUser := modelsGlobal.User{
		ID:       1,
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
	}

	service.EXPECT().Signup(r.Context(), &inputUser).Return(outputUser, nil)
	service.EXPECT().CreateSession(r.Context(), &outputUser).Return(modelsGlobal.Session{}, errors.ErrCreateSession)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrCreateSession.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewSingUpHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// check code
	require.Equal(t, http.StatusInternalServerError, w.Code)

	// check body
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
