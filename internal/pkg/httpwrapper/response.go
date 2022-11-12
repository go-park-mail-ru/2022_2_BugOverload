package httpwrapper

import (
	"context"
	"encoding/json"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// Response is a function for giving any response with a JSON body
func Response(ctx context.Context, w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		DefaultHandlerError(ctx, w, errors.ErrJSONUnexpectedEnd)
		return
	}

	w.Header().Set("Content-Type", pkg.ContentTypeJSON)

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		DefaultHandlerError(ctx, w, err)
		return
	}
}

// ResponseImage is a function for giving any response with a body - image
func ResponseImage(ctx context.Context, w http.ResponseWriter, statusCode int, image []byte) {
	w.Header().Set("Content-Type", pkg.ContentTypeWEBP)

	w.WriteHeader(statusCode)

	_, err := w.Write(image)
	if err != nil {
		DefaultHandlerError(ctx, w, err)
		return
	}
}
