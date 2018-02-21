package jsonhelper

import (
	"net/http"
)

// BadRequest function return status code 400 as the response body
func BadRequest(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusBadRequest, msg)
}

// Unauthorized function return status code 401 as the response body
func Unauthorized(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusUnauthorized, msg)
}

// Forbidden function return status code 403 as the response body
func Forbidden(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusForbidden, msg)
}

// NotFound function return status code 404 as the response body
func NotFound(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusNotFound, msg)
}

// MethodNotAllowed function return status code 405 as the response body
func MethodNotAllowed(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusMethodNotAllowed, msg)
}
