package integrationhandlerstests_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/incinemafilmshandler"
	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/repository/memory"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/app/collection/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
	"go-park-mail-ru/2022_2_BugOverload/test/integrationhandlerstests"
)

func TestInCinemaHandler(t *testing.T) {
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

	collectionService := serviceCollection.NewCollectionService(cs, params.ContextTimeout)
	inCinemaHandler := incinemafilmshandler.NewHandler(collectionService)

	url := "http://localhost:8088/v1/in_cinema"

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		w := httptest.NewRecorder()

		inCinemaHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))

		body, err := io.ReadAll(resp.Body)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Body.Close must be success"))

		collectionResponse := models.NewFilmCollection("", []models.Film{})

		err = json.Unmarshal(body, collectionResponse)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Marshal must be success"))

		collection, err := cs.GetInCinema(context.TODO())
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "GetInCinema must be success"))

		require.Equal(t, collection, collectionResponse.Films, utils.TestErrorMessage(caseNum, "Wrong body"))
	}
}
