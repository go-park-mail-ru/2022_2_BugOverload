package myhttp

import (
	"encoding/json"
	"go-park-mail-ru/2022_2_BugOverload/project/application/structs"
	"net/http"
)

func Success(w http.ResponseWriter, u structs.User) {
	out, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	//  Отдаем ответ
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	w.Write(out)
}
