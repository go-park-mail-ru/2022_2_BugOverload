package handlers

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestUserGetActivityOnFilmHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userActivityService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/user_activity", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	var userColections []modelsGlobal.NodeInUserCollection
	userColections = append(userColections, modelsGlobal.NodeInUserCollection{
		ID:     401,
		Name:   "Избранное",
		IsUsed: true,
	})

	userCollections := modelsGlobal.UserActivity{
		CountReviews: 44,
		Rating:       5,
		DateRating:   "2022.12.29",
		Collections:  userColections,
	}

	// Settings mock
	userActivityService.EXPECT().GetUserActivityOnFilm(r.Context(), &user, &constparams.GetUserActivityOnFilmParams{
		FilmID: 1,
	}).Return(userCollections, nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetActivityOnFilmHandler(userActivityService)
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

	expectedBody := models.NewGetUserActivityOnFilmResponse(&userCollections)

	var actualBody models.GetUserActivityOnFilmResponse

	err = easyjson.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, *expectedBody, actualBody, "Wrong body")
}

func TestUserGetActivityOnFilmHandler_Action_NotOK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userActivityService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/user_activity", nil)

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetActivityOnFilmHandler(userActivityService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusBadRequest, w.Code)

	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	var actualBody wrapper.ErrResponse

	err = easyjson.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetActivityOnFilmHandler_Action_UserNotFound(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userActivityService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/user_activity", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetActivityOnFilmHandler(userActivityService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusInternalServerError, w.Code, "Wrong StatusCode", pkg.GetResponseBody(*w))

	// Check body
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	require.Nil(t, errResponse, "io.ReadAll must be success")

	err := response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrGetUserRequest.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = easyjson.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestUserGetActivityOnFilmHandler_Action_ServiceError(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userActivityService := mockUserService.NewMockUserService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/user_activity", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedErr := errors.ErrUserNotFound

	// Settings mock
	userActivityService.EXPECT().GetUserActivityOnFilm(r.Context(), &user, &constparams.GetUserActivityOnFilmParams{
		FilmID: 1,
	}).Return(modelsGlobal.UserActivity{}, expectedErr)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetActivityOnFilmHandler(userActivityService)
	handler.Configure(router, nil)

	// Check result
	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNotFound, w.Code, "Wrong StatusCode", pkg.GetResponseBody(*w))

	// Check body
	response := w.Result()

	bodyResponse, errResponse := io.ReadAll(response.Body)
	require.Nil(t, errResponse, "io.ReadAll must be success")

	err := response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: expectedErr.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = easyjson.Unmarshal(bodyResponse, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
