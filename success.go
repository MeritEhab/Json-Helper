package jsonhelper

import (
	"net/http"
)

// Ok function return status code 200 as the response body
func Ok(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusOK, msg)
}

// Created function return status code 201 as the response body
func Created(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusCreated, msg)
}

// Accepted function return status code 202 as the response body
func Accepted(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusAccepted, msg)
}
