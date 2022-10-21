package httpwrapper

import (
	"net/http"
)

// NoContent is function designed to give a response that
// has no request body, only the code and/or headers matter
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
