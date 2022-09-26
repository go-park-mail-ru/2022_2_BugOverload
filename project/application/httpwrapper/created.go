package httpwrapper

import (
	"encoding/json"
	"net/http"
)

// Created is function for generating a successful request
func Created(w http.ResponseWriter, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	//  Отдаем ответ
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(out)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}
}
