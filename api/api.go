package api

import (
	"encoding/json"
	"fintech_app/helpers"
	"fintech_app/users"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandlerErr(err)

	// Formatting request body
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandlerErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	// Check if all is fine and prepare response
	if login["message"] == "Authentication is fine" {
		resp := login
		err := json.NewEncoder(w).Encode(resp)
		helpers.HandlerErr(err)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		err := json.NewEncoder(w).Encode(resp)
		helpers.HandlerErr(err)
	}
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
