package httpwrapper

import (
	"encoding/json"
	"net/http"
)

// Response is function for generating response
func Response(w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		DefHandlerError(w, err)
		return
	}

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		DefHandlerError(w, err)
		return
	}
}
