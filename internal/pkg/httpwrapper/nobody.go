package httpwrapper

import (
	"net/http"
)

// NoBody is function designed to give a response that
// has no request body, only the code and/or headers matter
func NoBody(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
