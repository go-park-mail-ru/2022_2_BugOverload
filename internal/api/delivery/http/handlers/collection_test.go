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

	"go-park-mail-ru/2022_2_BugOverload/internal/api/delivery/http/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/wrapper"
	mockWarehouseClient "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client/mocks"
)

func TestCollectionHandler_Action_OK(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodGet, "/api/v1/collections/413?sort_param=date", nil)
	vars := make(map[string]string)
	vars["id"] = "413"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	inputParams := constparams.CollectionGetFilmsRequestParams{
		CollectionID: 413,
		SortParam:    "date",
	}

	output := modelsGlobal.Collection{
		ID:   413,
		Name: "Избранное",
		Films: []modelsGlobal.Film{
			{
				ID:   1,
				Name: "testname film",
			},
		},
	}

	// Settings mock
	collectionService.EXPECT().GetCollectionFilmsAuthorized(r.Context(), &user, &inputParams).Return(output, nil)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetCollectionHandler(collectionService)
	handler.Configure(router, nil)

	// CheckNotificationSent result
	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusOK, w.Code)

	// CheckNotificationSent body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewCollectionResponse(&output)

	var actualBody *models.CollectionResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestCollectionHandler_Action_BindError(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodGet, "/api/v1/collections/413?sort_param=date", nil)
	vars := make(map[string]string)
	vars["id"] = "trash"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetCollectionHandler(collectionService)
	handler.Configure(router, nil)

	// CheckNotificationSent result
	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusBadRequest, w.Code)

	// CheckNotificationSent body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: errors.ErrConvertQueryType.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestCollectionHandler_Action_ServiceAuthorizedError(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodGet, "/api/v1/collections/413?sort_param=date", nil)
	vars := make(map[string]string)
	vars["id"] = "413"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	// Create required setup for handling
	user := modelsGlobal.User{
		ID: 1,
	}

	ctx := context.WithValue(r.Context(), constparams.CurrentUserKey, user)
	r = r.WithContext(ctx)

	inputParams := constparams.CollectionGetFilmsRequestParams{
		CollectionID: 413,
		SortParam:    "date",
	}

	expectedErr := errors.ErrCollectionIsNotPublic

	// Settings mock
	collectionService.EXPECT().GetCollectionFilmsAuthorized(r.Context(), &user, &inputParams).Return(modelsGlobal.Collection{}, expectedErr)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetCollectionHandler(collectionService)
	handler.Configure(router, nil)

	// CheckNotificationSent result
	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusForbidden, w.Code)

	// CheckNotificationSent body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: expectedErr.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}

func TestCollectionHandler_Action_ServiceNotAuthorizedError(t *testing.T) {
	// Init mock
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	// Data
	r := httptest.NewRequest(http.MethodGet, "/api/v1/collections/413?sort_param=date", nil)
	vars := make(map[string]string)
	vars["id"] = "413"
	r = mux.SetURLVars(r, vars)

	r.Header.Set("Content-Type", "application/json")

	inputParams := constparams.CollectionGetFilmsRequestParams{
		CollectionID: 413,
		SortParam:    "date",
	}

	expectedErr := errors.ErrCollectionIsNotPublic

	// Settings mock
	collectionService.EXPECT().GetCollectionFilmsNotAuthorized(r.Context(), &inputParams).Return(modelsGlobal.Collection{}, expectedErr)

	// Init
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewGetCollectionHandler(collectionService)
	handler.Configure(router, nil)

	// CheckNotificationSent result
	handler.Action(w, r)

	// CheckNotificationSent code
	require.Equal(t, http.StatusForbidden, w.Code)

	// CheckNotificationSent body
	response := w.Result()

	body, err := io.ReadAll(response.Body)
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := wrapper.ErrResponse{
		ErrMassage: expectedErr.Error(),
	}

	var actualBody wrapper.ErrResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
