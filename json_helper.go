package jsonhelper

import (
	"encoding/json"
	"net/http"
)

// Write function return the passed message as json with a status code
func Write(w http.ResponseWriter, status int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(msg)

	if err != nil {
		Write(w, status, err.Error())
	}
}
