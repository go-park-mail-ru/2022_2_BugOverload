package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	mockUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
)

func TestUserProfileHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/profile/1", nil)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 0,
	}

	resUser := modelsGlobal.User{
		ID: 0,
		Nickname: "Mike",
		Avatar: "201",
		CountCollections: 2,
		CountReviews: 3,
		CountRatings: 1,
		JoinedDate: "2022.12.02",
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	// Settings mock
	rateService.EXPECT().GetUserProfileByID(r.Context(), &user).Return(resUser, nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewUserProfileHandler(rateService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusOK, w.Code)

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewUserProfileResponse(&resUser)

	var actualBody models.UserProfileResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, &actualBody, "Wrong body")
}

func TestUserProfileHandler_Action_NotOK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rateService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/user/profile/1", nil)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 0,
	}

	resUser := modelsGlobal.User{
		ID: 0,
		Nickname: "Mike",
		Avatar: "201",
		CountCollections: 2,
		CountReviews: 3,
		CountRatings: 1,
		JoinedDate: "2022.12.02",
	}

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	// Settings mock
	rateService.EXPECT().GetUserProfileByID(r.Context(), &user).Return(resUser, errors.ErrNotFoundInDB)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewUserProfileHandler(rateService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNotFound, w.Code)

	// Check body
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
