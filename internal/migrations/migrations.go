package migrations

import (
	"fintech_app/internal/pkg/database"
	"fintech_app/internal/pkg/helpers"
	"fintech_app/internal/pkg/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {
	users := &[3]interfaces.User{
		{
			Username: "Olga",
			Email:    "olga@gmail.com",
		},
		{
			Username: "Pavel",
			Email:    "pavel@yandex.ru",
		},
		{
			Username: "Ivan",
			Email:    "ivan@gmail.com",
		},
	}
	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{
			Username: users[i].Username,
			Email:    users[i].Email,
			Password: generatePassword,
		}
		database.DB.Create(&user)

		account := &interfaces.Account{
			Type:    "Credit account",
			Name:    users[i].Username + "'s " + "account",
			Balance: uint(1000 * (i + 1)),
			UserID:  user.ID,
		}
		database.DB.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, Transactions)

	createAccounts()
}
