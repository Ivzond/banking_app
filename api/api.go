package api

import (
	"encoding/json"
	"fintech_app/helpers"
	"fintech_app/useraccounts"
	"fintech_app/users"
	"fmt"
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

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "OK" {
		resp := call
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			helpers.HandleErr(err)
		}
	} else {
		resp := call
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			helpers.HandleErr(err)
		}
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
	}

	login := users.Login(formattedBody.Username, formattedBody.Password)

	// Check if all is fine and prepare response
	apiResponse(login, w)
}

func register(w http.ResponseWriter, r *http.Request) {
	// Read body
	body := readBody(r)

	// Formatting request body
	var formattedBody Register
	err := json.Unmarshal(body, &formattedBody)
	if err != nil {
		helpers.HandleErr(err)
	}

	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)

	// Check if all is fine and prepare response
	apiResponse(register, w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
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
	}

	transaction := useraccounts.Transaction(
		formattedBody.UserID,
		formattedBody.From,
		formattedBody.To,
		formattedBody.Amount,
		auth,
	)
	apiResponse(transaction, w)
}

func StartApi() {
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	fmt.Println("App is working on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
