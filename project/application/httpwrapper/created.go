package httpwrapper

import (
	"encoding/json"
	"net/http"
)

func Created(w http.ResponseWriter, someStruct interface{}) {
	out, err := json.Marshal(someStruct)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	//  Отдаем ответ
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	w.Write(out)
}
