package httpwrapper

import (
	"encoding/json"
	"net/http"
)

// ResponseOK is function for generating response
func ResponseOK(w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//  Отдаем ответ
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
