package httpwrapper

import (
	"encoding/json"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/OLD/application/errors"
)

// Response is function for generating response
func Response(w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		DefaultHandlerError(w, errors.NewErrValidation(errors.ErrCJSONUnexpectedEnd))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		DefaultHandlerError(w, err)
		return
	}
}
