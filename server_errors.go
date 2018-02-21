package jsonhelper

import (
	"net/http"
)

// InternalServerError function return status code 500 as the response body
func InternalServerError(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusInternalServerError, msg)
}

// BadGateway function return status code 502 as the response body
func BadGateway(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusBadGateway, msg)
}
