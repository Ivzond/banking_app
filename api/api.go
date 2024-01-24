package api

import (
	"encoding/json"
	"fintech_app/helpers"
	"fintech_app/transactions"
	"fintech_app/useraccounts"
	"fintech_app/users"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type TransactionBody struct {
	UserID uint
	From   uint
	To     uint
	Amount int
}

func readBody(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.HandleErr(err)
	}

	return body
}

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

func login(w http.ResponseWriter, r *http.Request) {
	// Read body
	body := readBody(r)

	// Formatting request body
	var formattedBody Login
	err := json.Unmarshal(body, &formattedBody)
	if err != nil {
		helpers.HandleErr(err)
		authApiResponse(map[string]interface{}{}, w)
		return
	}

	login := users.Login(formattedBody.Username, formattedBody.Password)
	authApiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	// Read body
	body := readBody(r)

	// Formatting request body
	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	if err != nil {
		helpers.HandleErr(err)
		authApiResponse(map[string]interface{}{}, w)
		return
	}

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	authApiResponse(register, w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	getUserApiResponse(user, w)
}

func transaction(w http.ResponseWriter, r *http.Request) {
	// Read body
	body := readBody(r)
	auth := r.Header.Get("Authorization")

	// Formatting request body
	var formattedBody TransactionBody
	err := json.Unmarshal(body, &formattedBody)
	if err != nil {
		helpers.HandleErr(err)
		transactionApiResponse(map[string]interface{}{}, w)
		return
	}
	transaction := useraccounts.Transaction(
		formattedBody.UserID,
		formattedBody.From,
		formattedBody.To,
		formattedBody.Amount,
		auth,
	)
	transactionApiResponse(transaction, w)
}

func getMyTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userID"]
	auth := r.Header.Get("Authorization")

	userTransactions := transactions.GetMyTransactions(userId, auth)
	getTransactionsApiResponse(userTransactions, w)
}

func StartApi() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	router.HandleFunc("/transaction/{userID}", getMyTransactions).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")

	// Use PanicHandler from helpers
	handler := helpers.PanicHandler(router)

	// Enable CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Start the server with CORS handling
	fmt.Println("App is working on port :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(handler)))
}
