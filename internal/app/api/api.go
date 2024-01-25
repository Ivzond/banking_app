package api

import (
	"fintech_app/internal/pkg/helpers"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func readBody(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.HandleErr(err)
	}

	return body
}

func StartApi() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/transaction", transaction).Methods("POST")
	router.HandleFunc("/account", createAccount).Methods("POST")
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
