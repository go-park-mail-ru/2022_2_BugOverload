package integrationhandlerstests_test

import (
	"context"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/recommendationfilmhandler"
	memoryFilms "go-park-mail-ru/2022_2_BugOverload/internal/app/films/repository/memory"
	serviceFilms "go-park-mail-ru/2022_2_BugOverload/internal/app/films/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	memoryCookie "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	serviceAuth "go-park-mail-ru/2022_2_BugOverload/internal/app/auth/service"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
	memoryUser "go-park-mail-ru/2022_2_BugOverload/internal/app/user/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/utils/params"
	"go-park-mail-ru/2022_2_BugOverload/test/integrationhandlerstests"
)

func TestRecommendationHandler(t *testing.T) {
	cases := []integrationhandlerstests.TestCase{
		// Success
		integrationhandlerstests.TestCase{
			Method:     http.MethodGet,
			Cookie:     "GeneratedData",
			StatusCode: http.StatusOK,
		},
	}

	url := "http://localhost:8088/v1/auth"

	// Base
	userMutex := &sync.Mutex{}
	authMutex := &sync.Mutex{}

	us := memoryUser.NewUserRepo(userMutex)
	cs := memoryCookie.NewCookieRepo(authMutex)

	testUser := &models.User{
		Nickname: "Andeo",
		Email:    "YasaPupkinEzji@top.world",
		Password: "Widget Adapter",
		Avatar:   "URL",
	}

	_, err := us.CreateUser(context.TODO(), testUser)
	require.Nil(t, err, utils.TestErrorMessage(-1, "Err create user for test"))

	var cookie string
	cookie, err = cs.CreateSession(context.TODO(), testUser)
	require.Nil(t, err, utils.TestErrorMessage(-1, "Err create session-cookie for test"))

	cases[0].Cookie = strings.Split(cookie, ";")[0]

	authService := serviceAuth.NewAuthService(cs, params.ContextTimeout)

	// Films
	pathPreview := "../testdata/preview.json"

	filmsMutex := &sync.Mutex{}

	fs := memoryFilms.NewFilmRepo(filmsMutex, pathPreview)

	filmsService := serviceFilms.NewFilmService(fs, params.ContextTimeout)

	recommendationHandler := recommendationfilmhandler.NewHandler(filmsService, authService)

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, url, nil)
		if item.Cookie != "" {
			req.Header.Set("Cookie", item.Cookie)
		}

		w := httptest.NewRecorder()

		recommendationHandler.Action(w, req)

		resp := w.Result()

		require.Equal(t, item.StatusCode, w.Code, utils.TestErrorMessage(caseNum, "Wrong StatusCode"))

		var body []byte
		body, err = io.ReadAll(resp.Body)
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "io.ReadAll must be success"))

		err = resp.Body.Close()
		require.Nil(t, err, utils.TestErrorMessage(caseNum, "Body.Close must be success"))

		require.True(t, string(body) != "", utils.TestErrorMessage(caseNum, "Wrong body"))
	}
}
