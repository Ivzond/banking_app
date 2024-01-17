package useraccounts

import (
	"fintech_app/helpers"
	"fintech_app/interfaces"
	"github.com/jinzhu/gorm"
)

func updateAccount(id int, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			helpers.HandleErr(err)
		}
	}(db)
}
