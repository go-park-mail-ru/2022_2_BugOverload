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
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/api/http/delivery/models"
	modelsGlobal "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	mockWarehouseClient "go-park-mail-ru/2022_2_BugOverload/internal/warehouse/delivery/grpc/client/mocks"
)

func TestTCollectionHandlerPremiere_Action_OK(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	collectionService := mockWarehouseClient.NewMockWarehouseService(ctrl)

	r := httptest.NewRequest(http.MethodGet, "/api/v1/premieres&count_films=1&delimiter=1", nil)

	res := modelsGlobal.Collection{
		Name: "популярное",
		Films: []modelsGlobal.Film{{
			Name:      "Игра престолов",
			ProdDate:  "2013",
			EndYear:   "2014",
			ID:        123,
			Rating:    7.12332,
			PosterVer: "123",
			Genres:    []string{"фэнтези", "приключения"},
		}},
	}

	oldLogger := logrus.New()
	logger := logrus.NewEntry(oldLogger)

	ctx := context.WithValue(r.Context(), constparams.LoggerKey, logger)

	r = r.WithContext(ctx)

	collectionService.EXPECT().GetPremieresCollection(r.Context(), &constparams.PremiersCollectionParams{
		CountFilms: 1,
		Delimiter:  1,
	}).Return(res, nil)

	w := httptest.NewRecorder()

	router := mux.NewRouter()
	handler := NewPremieresCollectionHandler(collectionService)
	handler.Configure(router, nil)

	handler.Action(w, r)


	// Check body
	response := w.Result()

	body, err := io.ReadAll(response.Body)

	// Check code
	require.Equal(t, http.StatusOK, w.Code, "Wrong StatusCode ")
	require.Nil(t, err, "io.ReadAll must be success")

	err = response.Body.Close()
	require.Nil(t, err, "Body.Close must be success")

	expectedBody := models.NewPremieresCollectionResponse(&res)

	var actualBody *models.GetStdCollectionResponse

	err = json.Unmarshal(body, &actualBody)
	require.Nil(t, err, "json.Unmarshal must be success")

	require.Equal(t, expectedBody, actualBody, "Wrong body")
}
