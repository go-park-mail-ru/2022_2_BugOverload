package tests_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/handlers"
	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	models2 "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestPopularHandler(t *testing.T) {
	cases := []tests.TestCase{
		// Success
		tests.TestCase{
			Method:     http.MethodGet,
			StatusCode: http.StatusOK,
		},
	}

	pathInCinema := "../../../test/data/incinema.json"
	pathPopular := "../../../test/data/popular.json"

	//  init
	cs := memoryCollection.NewCollectionCache(pathPopular, pathInCinema)

	collectionService := serviceCollection.NewCollectionService(cs)
	popularHandler := handlers.NewPopularFilmsHandler(collectionService)

	url := "http://localhost:8088/api/v1/collections/popular"

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		w := httptest.NewRecorder()

		popularHandler.Action(w, req)

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		if item.ResponseBody != "" {
			resp := w.Result()

			body, err := io.ReadAll(resp.Body)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

			err = resp.Body.Close()
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

			var collectionResponse models2.Collection

			err = json.Unmarshal(body, &collectionResponse)
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Marshal must be success"))

			collection, err := cs.GetInCinema(context.TODO())
			require.Nil(t, err, pkg.TestErrorMessage(caseNum, "GetInCinema must be success"))

			require.Equal(t, collection, collectionResponse.Films, pkg.TestErrorMessage(caseNum, "Wrong body"))
		}
	}
}
