package jsonhelper

import (
	"net/http"
)

// MovedPermanently function return status code 301 as the response body
func MovedPermanently(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusMovedPermanently, msg)
}

// Found function return status code 302 as the response body
func Found(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusFound, msg)
}
