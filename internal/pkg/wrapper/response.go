package wrapper

import (
	"context"
	"net/http"

	"github.com/mailru/easyjson"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/constparams"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

func getEasyJSON(someStruct interface{}) ([]byte, error) {
	someStructUpdate, ok := someStruct.(easyjson.Marshaler)
	if !ok {
		return []byte{}, errors.ErrGetEasyJSON
	}

	out, err := easyjson.Marshal(someStructUpdate)
	if !ok {
		return []byte{}, errors.ErrJSONUnexpectedEnd
	}

	return out, err
}

// Response is a function for giving any response with a JSON body
func Response(ctx context.Context, w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := getEasyJSON(someStruct)
	if err != nil {
		DefaultHandlerHTTPError(ctx, w, err)
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
