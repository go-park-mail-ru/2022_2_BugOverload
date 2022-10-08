package incinemafilmshandler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/collection/delivery/http/handlers/incinemafilmshandler"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/models"
)

// TestCase is structure for API testing
type TestCase struct {
	Method       string
	ContentType  string
	ResponseBody string
	URL          string
	StatusCode   int
}

func TestFilmsHandlerInCinema(t *testing.T) {
	currentTestCase := TestCase{
		URL:        "http://localhost:8088/v1/in_cinema",
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
	}

	fs := memory.NewFilmStorage()

	filmsHandler := incinemafilmshandler.NewCollectionInCinemaHandler(fs)

	req := httptest.NewRequest(currentTestCase.Method, currentTestCase.URL, nil)
	w := httptest.NewRecorder()

	filmsHandler.Action(w, req)

	if w.Code != currentTestCase.StatusCode {
		t.Errorf("Wrong StatusCode: got [%d], expected [%d]", w.Code, currentTestCase.StatusCode)
	}

	resp := w.Result()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err)
	}
	err = resp.Body.Close()
	if err != nil {
		t.Errorf("Err: [%s], expected: nil", err)
	}

	var responseCollection models.FilmCollection
	err = json.Unmarshal(body, &responseCollection)

	if err != nil {
		t.Error("In cinema test: wrong response body, unmarshal error")
	}

	if responseCollection.Title != "Сейчас в кино" {
		t.Errorf("Wrong Title: got [%s], expected [%s]", responseCollection.Title, "Сейчас в кино")
	}

	for _, film := range responseCollection.Films {
		filmFromStorage, _ := fs.GetFilm(film.ID)
		if !cmp.Equal(film, filmFromStorage, cmpopts.IgnoreFields(models.Film{}, "Rating")) {
			t.Errorf("Wrong Film: got [%v], expected [%v]", film, filmFromStorage)
		}
	}
}
