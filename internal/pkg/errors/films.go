package errors

import (
	"fmt"
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	ErrFilmNotFound  = stdErrors.New("no such film")
	ErrFilmsNotFound = stdErrors.New("no such films")
)

type errClassifierFilms struct {
	table map[error]int
}

func NewErrClassifierFilms() errClassifier {
	res := make(map[error]int)

	res[ErrFilmNotFound] = http.StatusNotFound
	res[ErrFilmsNotFound] = http.StatusNotFound

	return &errClassifierFilms{
		table: res,
	}
}

func (ec *errClassifierFilms) GetCode(err error) int {
	code, exist := ec.table[err]
	if !exist {
		return http.StatusInternalServerError
	}

	return code
}

var ErrCsfFilms = NewErrClassifierFilms()

type FilmsError struct {
	Reason string
	Code   int
}

func (e FilmsError) Error() string {
	return fmt.Sprintf("Films: [%s]", e.Reason)
}

func NewErrFilms(err error) FilmsError {
	return FilmsError{
		Reason: err.Error(),
		Code:   ErrCsfFilms.GetCode(err),
	}
}
