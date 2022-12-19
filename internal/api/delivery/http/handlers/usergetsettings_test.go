package handlers

import (
	"context"
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	mockUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
)

func TestUserGetSettingsHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/settings", nil)

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	resUser := modelsGlobal.User{
		CountViewsFilms:  1,
		CountCollections: 2,
		CountReviews:     44,
		CountRatings:     1,
		JoinedDate:       "2022.12.02",
	}

	service.EXPECT().GetUserProfileSettings(r.Context(), &user).Return(resUser, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetSettingsHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusOK, w.Code)

	// CheckNewNotification body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewGetUserSettingsResponse(&resUser)

	var actualBody *models.GetUserSettingsResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetSettingsHandler_Action_UserNotFound(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/settings", nil)

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetSettingsHandler(service)
	handler.Configure(router, nil)

	// CheckNewNotification result
	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusInternalServerError, w.Code, "Wrong StatusCode")

	// CheckNewNotification body
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	require.Nil(t, errResponse, "io.ReadAll must be success")

	err := response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrGetUserRequest.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetSettingsHandler_Action_ServiceError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/settings", nil)

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedErr := errors.ErrUserNotFound

	service.EXPECT().GetUserProfileSettings(r.Context(), &user).Return(modelsGlobal.User{}, expectedErr)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetSettingsHandler(service)
	handler.Configure(router, nil)

	// CheckNewNotification result
	handler.Action(w, r)

	// CheckNewNotification code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode")

	// CheckNewNotification body
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	require.Nil(t, errResponse, "io.ReadAll must be success")

	err := response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: expectedErr.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
