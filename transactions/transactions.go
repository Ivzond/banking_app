package transactions

import (
	"fintech_app/helpers"
	"fintech_app/interfaces"
	"github.com/jinzhu/gorm"
)

func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()

	transaction := &interfaces.Transaction{
		From:   From,
		To:     To,
		Amount: Amount,
	}
	db.Create(transaction)

	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			helpers.HandleErr(err)
		}
	}(db)
}
