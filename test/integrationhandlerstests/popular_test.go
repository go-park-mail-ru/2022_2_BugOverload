package integrationhandlerstests_test

import (
	"context"
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/popularfilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/contextparams"
	"go-park-mail-ru/2022_2_BugOverload/test/integrationhandlerstests"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/repository/memory"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
)

func TestPopularHandler(t *testing.T) {
	cases := []integrationhandlerstests.TestCase{
		// Success
		integrationhandlerstests.TestCase{
			Method:     http.MethodGet,
			StatusCode: http.StatusOK,
		},
	}

	pathInCinema := "../testdata/incinema.json"
	pathPopular := "../testdata/popular.json"

	//  init
	collectionMutex := &sync.Mutex{}
	cs := memoryCollection.NewCollectionRepo(collectionMutex, pathPopular, pathInCinema)

	collectionService := serviceCollection.NewCollectionService(cs, contextparams.ContextTimeout)
	popularHandler := popularfilmshandler.NewHandler(collectionService)

	url := "http://localhost:8088/v1/popular_films"

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		w := httptest.NewRecorder()

		popularHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))

		body, err := io.ReadAll(resp.Body)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Body.Close must be success"))

		collectionResponse := models.NewFilmCollection("", []models.Film{})

		err = json.Unmarshal(body, collectionResponse)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Marshal must be success"))

		collection, err := cs.GetPopular(context.TODO())
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "GetPopular must be success"))

		require.Equal(t, collection, collectionResponse.Films, utils.TestErrorMessage(caseNum, "Wrong body"))
	}
}
