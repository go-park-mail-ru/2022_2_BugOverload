package content_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"go-park-mail-ru/2022_2_BugOverload/project/application/database"
	"go-park-mail-ru/2022_2_BugOverload/project/application/handlers/content"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
)

// TestCase is structure for API testing
type TestCase struct {
	Method       string
	ContentType  string
	ResponseBody string
	URL          string
	StatusCode   int
}

func TestFilmsHandlerPopular(t *testing.T) {
	currentTestCase := TestCase{
		URL:        "http://localhost:8088/v1/popular_films",
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
	}

	fs := database.NewFilmStorage()
	fs.FillFilmStoragePartOne()
	fs.FillFilmStoragePartTwo()

	filmsHandler := content.NewHandlerFilms(fs)

	req := httptest.NewRequest(currentTestCase.Method, currentTestCase.URL, nil)
	w := httptest.NewRecorder()

	filmsHandler.GetPopularFilms(w, req)

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

	var responseCollection structs.FilmCollection
	err = json.Unmarshal(body, &responseCollection)

	if err != nil {
		t.Error("Popular films test: wrong response body, unmarshal error")
	}

	if responseCollection.Title != "Popular films" {
		t.Errorf("Wrong Title: got [%s], expected [%s]", responseCollection.Title, "Popular films")
	}

	for _, film := range responseCollection.Films {
		filmFromStorage, _ := fs.GetFilm(film.ID)
		if !cmp.Equal(film, filmFromStorage, cmpopts.IgnoreFields(structs.Film{}, "Rating")) {
			t.Errorf("Wrong Film: got [%v], expected [%v]", film, filmFromStorage)
		}
	}
}

func TestFilmsHandlerInCinema(t *testing.T) {
	currentTestCase := TestCase{
		URL:        "http://localhost:8088/v1/in_cinema",
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
	}

	fs := database.NewFilmStorage()
	fs.FillFilmStoragePartOne()
	fs.FillFilmStoragePartTwo()

	filmsHandler := content.NewHandlerFilms(fs)

	req := httptest.NewRequest(currentTestCase.Method, currentTestCase.URL, nil)
	w := httptest.NewRecorder()

	filmsHandler.GetFilmsInCinema(w, req)

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

	var responseCollection structs.FilmCollection
	err = json.Unmarshal(body, &responseCollection)

	if err != nil {
		t.Error("In cinema test: wrong response body, unmarshal error")
	}

	if responseCollection.Title != "In cinema" {
		t.Errorf("Wrong Title: got [%s], expected [%s]", responseCollection.Title, "In cinema")
	}

	for _, film := range responseCollection.Films {
		filmFromStorage, _ := fs.GetFilm(film.ID)
		if !cmp.Equal(film, filmFromStorage, cmpopts.IgnoreFields(structs.Film{}, "Rating")) {
			t.Errorf("Wrong Film: got [%v], expected [%v]", film, filmFromStorage)
		}
	}
}

func TestFilmsHandlerRecomended(t *testing.T) {
	currentTestCase := TestCase{
		URL:        "http://localhost:8088/v1/recommendation_film",
		Method:     http.MethodGet,
		StatusCode: http.StatusOK,
	}

	fs := database.NewFilmStorage()
	fs.FillFilmStoragePartOne()
	fs.FillFilmStoragePartTwo()

	filmsHandler := content.NewHandlerFilms(fs)

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

	var responseFilm structs.Film

	err = json.Unmarshal(body, &responseFilm)

	if err != nil {
		t.Errorf("[%d] wrong response body, unmarshal error", 2)
	}

	filmFromStorage, _ := fs.GetFilm(responseFilm.ID)

	if !cmp.Equal(responseFilm, filmFromStorage, cmpopts.IgnoreFields(structs.Film{}, "Rating")) {
		t.Errorf("[%d] wrong Film: got [%v], expected [%v]", 2, responseFilm, filmFromStorage)
	}
}

func TestFilmsHandlerEmptyStorage(t *testing.T) {
	cases := []TestCase{
		// unsuccess request popular films
		TestCase{
			URL:          "http://localhost:8088/v1/popular_films",
			Method:       http.MethodGet,
			StatusCode:   http.StatusNotFound,
			ResponseBody: "no such films\n",
		},
		// unsuccess request in cinema
		TestCase{
			URL:          "http://localhost:8088/v1/in_cinema",
			Method:       http.MethodGet,
			StatusCode:   http.StatusNotFound,
			ResponseBody: "no such films\n",
		},
		// unsuccess request recommendation film
		TestCase{
			URL:          "http://localhost:8088/v1/recommendation_film",
			Method:       http.MethodGet,
			StatusCode:   http.StatusNotFound,
			ResponseBody: "such film doesn't exist\n",
		},
	}

	fs := database.NewFilmStorage()

	filmsHandler := content.NewHandlerFilms(fs)

	for caseNum, item := range cases {
		req := httptest.NewRequest(item.Method, item.URL, nil)
		w := httptest.NewRecorder()

		switch item.URL {
		case "http://localhost:8088/v1/popular_films":
			filmsHandler.GetPopularFilms(w, req)
		case "http://localhost:8088/v1/in_cinema":
			filmsHandler.GetFilmsInCinema(w, req)
		case "http://localhost:8088/v1/recommendation_film":
			filmsHandler.GetRecommendedFilm(w, req)
		}

		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got [%d], expected [%d]", caseNum, w.Code, item.StatusCode)
		}

		resp := w.Result()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("[%d] err: [%s], expected: nil", caseNum, err)
		}
		err = resp.Body.Close()
		if err != nil {
			t.Errorf("[%d] err: [%s], expected: nil", caseNum, err)
		}

		bodyStr := string(body)
		if bodyStr != item.ResponseBody {
			t.Errorf("[%d] wrong Response: got [%+v], expected [%+v]", caseNum, bodyStr, item.ResponseBody)
		}
	}
}
