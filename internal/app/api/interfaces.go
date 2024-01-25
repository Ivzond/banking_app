package api

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TransactionBody struct {
	UserID uint `json:"userId"`
	From   uint `json:"from"`
	To     uint `json:"to"`
	Amount int  `json:"amount"`
}

type CreateAccountRequest struct {
	UserID uint   `json:"userID"`
	Type   string `json:"type"`
	Name   string `json:"name"`
}
