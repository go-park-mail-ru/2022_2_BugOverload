package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockUserService "go-park-mail-ru/2022_2_BugOverload/internal/user/service/mocks"
)

func TestUserfilmAddtocollectionHandler_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]int{
		"collection_id": 401,
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/save", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	service.EXPECT().AddFilmToUserCollection(r.Context(), &user, &constparams.UserCollectionFilmsUpdateParams{
		FilmID:       1,
		CollectionID: 401,
	}).Return(nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewAddFilmToUserCollectionHandler(service)
	handler.Configure(router, nil)

	handler.Action(w, r)

	// Check code
	require.Equal(t, http.StatusNoContent, w.Code)
}

func TestUserfilmAddtocollectionHandler_Action_NotOK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	mcPostBody := map[string]int{
		"collection_id": 401,
	}
	body, _ := json.Marshal(mcPostBody)
	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/save", bytes.NewReader(body))
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrNotFoundInDB.Error(),
	}

	service.EXPECT().AddFilmToUserCollection(r.Context(), &user, &constparams.UserCollectionFilmsUpdateParams{
		FilmID:       1,
		CollectionID: 401,
	}).Return(errors.ErrNotFoundInDB)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewAddFilmToUserCollectionHandler(service)
	handler.Configure(router, nil)

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

func TestUserfilmAddtocollectionHandler_Action_InvBody(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mockUserService.NewMockUserService(ctrl)

	r := httptest.NewRequest(http.MethodPost, "/api/v1/film/1/save", nil)
	vars := make(map[string]string)
	vars["id"] = "1"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrEmptyBody.Error(),
	}

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewAddFilmToUserCollectionHandler(service)
	handler.Configure(router, nil)

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

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
