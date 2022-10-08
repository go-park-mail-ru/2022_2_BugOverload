package httpwrapper

import (
	"encoding/json"
	errors2 "go-park-mail-ru/2022_2_BugOverload/internal/app/utils/errors"
	"net/http"
)

// Response is function for generating response
func Response(w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		DefaultHandlerError(w, errors2.NewErrValidation(errors2.ErrCJSONUnexpectedEnd))
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
