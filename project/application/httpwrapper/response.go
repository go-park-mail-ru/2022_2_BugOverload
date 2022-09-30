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

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		DefHandlerError(w, err)
		return
	}
}
