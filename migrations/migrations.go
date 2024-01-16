package migrations

import (
	"fintech_app/helpers"
	"fintech_app/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createAccounts() {
	db := helpers.ConnectDB()

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
		db.Create(&user)

		account := &interfaces.Account{
			Type:    "Credit account",
			Name:    users[i].Username + "'s " + "account",
			Balance: uint(1000 * (i + 1)),
			UserID:  user.ID,
		}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}

	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}
