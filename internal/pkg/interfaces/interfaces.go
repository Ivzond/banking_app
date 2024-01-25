package interfaces

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

type Transaction struct {
	gorm.Model
	From   uint
	To     uint
	Amount int
}

type ResponseTransaction struct {
	ID     uint
	From   uint
	To     uint
	Amount int
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance uint
}

type ResponseUser struct {
	ID       uint
	Name     string
	Username string
	Email    string
	Accounts []ResponseAccount
}

type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}
