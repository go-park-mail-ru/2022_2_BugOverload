package tests_test

import (
	"context"
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/collection/delivery/handlers"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/params"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"go-park-mail-ru/2022_2_BugOverload/cmd/debug/tests"
	memoryCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/repository"
	serviceCollection "go-park-mail-ru/2022_2_BugOverload/internal/collection/service"
	models2 "go-park-mail-ru/2022_2_BugOverload/internal/models"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

func TestInCinemaHandler(t *testing.T) {
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
	cs := memoryCollection.NewCollectionCash(pathPopular, pathInCinema)

	collectionService := serviceCollection.NewCollectionService(cs, params.ContextTimeout)
	inCinemaHandler := handlers.NewInCinemaHandler(collectionService)

	url := "http://localhost:8088/v1/in_cinema"

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		w := httptest.NewRecorder()

		inCinemaHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, pkg.TestErrorMessage(caseNum, "Wrong StatusCode"))

		body, err := io.ReadAll(resp.Body)
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Body.Close must be success"))

		collectionResponse := models2.NewFilmCollection("", []models2.Film{})

		err = json.Unmarshal(body, collectionResponse)
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "Marshal must be success"))

		collection, err := cs.GetInCinema(context.TODO())
		require.Nil(t, err, pkg.TestErrorMessage(caseNum, "GetInCinema must be success"))

		require.Equal(t, collection, collectionResponse.Films, pkg.TestErrorMessage(caseNum, "Wrong body"))
	}
}
