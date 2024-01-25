package api

import (
	"encoding/json"
	"fintech_app/internal/pkg/helpers"
	"fintech_app/internal/pkg/transactions"
	"fintech_app/internal/pkg/useraccounts"
	"fintech_app/internal/pkg/users"
	"github.com/gorilla/mux"
	"net/http"
)

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

	register := users.Register(formattedBody.Name, formattedBody.Username, formattedBody.Email, formattedBody.Password)
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

func createAccount(w http.ResponseWriter, r *http.Request) {
	// Read body
	body := readBody(r)

	// Formatting request body
	var formattedBody CreateAccountRequest
	err := json.Unmarshal(body, &formattedBody)
	if err != nil {
		helpers.HandleErr(err)
		authApiResponse(map[string]interface{}{}, w)
		return
	}

	// Call the function to create a new account
	response := useraccounts.CreateAccount(formattedBody.UserID, formattedBody.Type, formattedBody.Name)

	// Return the response
	authApiResponse(response, w)
}
