package httpwrapper

import (
	"encoding/json"
	"net/http"
)

// Success is function for generating response success
func Success(w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		DefHandlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		DefHandlerError(w, err)
		return
	}
}
