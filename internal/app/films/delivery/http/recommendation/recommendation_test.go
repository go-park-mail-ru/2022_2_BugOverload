package recommendation_test

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go-park-mail-ru/2022_2_BugOverload/internal/app/auth/repository/memory"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-park-mail-ru/2022_2_BugOverload/internal/app/films/delivery/http/recommendation"
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

func TestFilmsHandlerRecommended(t *testing.T) {
	currentTestCase := TestCase{
		URL:        "http://localhost:8088/v1/recommendation_film",
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
	}

	fs := memory.NewFilmStorage()

	filmsHandler := recommendation.NewHandlerRecommendationFilm(fs)

	req := httptest.NewRequest(currentTestCase.Method, currentTestCase.URL, nil)
	w := httptest.NewRecorder()

	filmsHandler.GetRecommendedFilm(w, req)

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

	var responseFilm models.Film

	err = json.Unmarshal(body, &responseFilm)

	if err != nil {
		t.Errorf("[%d] wrong response body, unmarshal error", 2)
	}

	filmFromStorage, _ := fs.GetFilm(responseFilm.ID)

	if !cmp.Equal(responseFilm, filmFromStorage, cmpopts.IgnoreFields(models.Film{}, "Rating", "PosterVer")) {
		t.Errorf("[%d] wrong Film: got [%v], expected [%v]", 2, responseFilm, filmFromStorage)
	}
}
