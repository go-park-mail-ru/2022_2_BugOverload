package httpwrapper

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// Response is a function for giving any response with a JSON body
func Response(w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		DefaultHandlerError(w, errors.NewErrValidation(errors.ErrCJSONUnexpectedEnd))
		return
	}

	w.Header().Set("Content-Type", pkg.ContentTypeJSON)

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		DefaultHandlerError(w, err)
		return
	}
}

// ResponseImage is a function for giving any response with a body - image
func ResponseImage(w http.ResponseWriter, statusCode int, image []byte) {
	w.Header().Set("Content-Type", pkg.ContentTypeImage)

	w.WriteHeader(statusCode)

	_, err := w.Write(image)
	if err != nil {
		DefaultHandlerError(w, err)
		return
	}
}
