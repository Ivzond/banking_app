package api

import (
	"encoding/json"
	"fintech_app/internal/pkg/helpers"
	"net/http"
)

func authApiResponse(call map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	if call["message"] == "OK" {
		statusCode = http.StatusOK
	} else if call["message"] == "User not found" {
		statusCode = http.StatusNotFound
	} else {
		statusCode = http.StatusBadRequest
	}
	w.WriteHeader(statusCode)

	resp := call
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		helpers.HandleErr(err)
	}
}

func getUserApiResponse(call map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	if call["message"] == "OK" {
		statusCode = http.StatusOK
	} else if call["message"] == "Not valid token" {
		statusCode = http.StatusUnauthorized
	} else {
		statusCode = http.StatusNotFound
	}
	w.WriteHeader(statusCode)

	resp := call
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		helpers.HandleErr(err)
	}
}

func getTransactionsApiResponse(call map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	if call["message"] == "OK" {
		statusCode = http.StatusOK
	} else if call["message"] == "Not valid token" {
		statusCode = http.StatusUnauthorized
	}
	w.WriteHeader(statusCode)

	resp := call
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		helpers.HandleErr(err)
	}
}

func transactionApiResponse(call map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	if call["message"] == "OK" {
		statusCode = http.StatusOK
	} else if call["message"] == "Account not found" {
		statusCode = http.StatusNotFound
	} else if call["message"] == "Your are not the owner of the account" || call["message"] == "Not valid token" {
		statusCode = http.StatusUnauthorized
	} else {
		statusCode = http.StatusBadRequest
	}
	w.WriteHeader(statusCode)

	resp := call
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		helpers.HandleErr(err)
	}
}
