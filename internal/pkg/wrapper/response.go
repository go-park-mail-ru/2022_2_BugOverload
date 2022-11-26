package wrapper

import (
	"context"
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"net/http"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

// Response is a function for giving any response with a JSON body
func Response(ctx context.Context, w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		DefaultHandlerHTTPError(ctx, w, errors.ErrJSONUnexpectedEnd)
		return
	}

	w.Header().Set("Content-Type", constparams.ContentTypeJSON)

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		DefaultHandlerHTTPError(ctx, w, err)
		return
	}
}

// ResponseImage is a function for giving any response with a body - image
func ResponseImage(ctx context.Context, w http.ResponseWriter, statusCode int, image []byte) {
	w.Header().Set("Content-Type", constparams.ContentTypeWEBP)

	w.WriteHeader(statusCode)

	_, err := w.Write(image)
	if err != nil {
		DefaultHandlerHTTPError(ctx, w, err)
		return
	}
}
