package transactions

import (
	"fintech_app/internal/pkg/database"
	"fintech_app/internal/pkg/helpers"
	"fintech_app/internal/pkg/interfaces"
	"github.com/jinzhu/gorm"
)

// CreateTransactionWithinTransaction creates a new transaction within a transaction
func CreateTransactionWithinTransaction(tx *gorm.DB, From uint, To uint, Amount int) {
	transaction := &interfaces.Transaction{
		From:   From,
		To:     To,
		Amount: Amount,
	}
	tx.Create(transaction)
}

func GetTransactionsByAccount(id uint) []interfaces.ResponseTransaction {
	var transactions []interfaces.ResponseTransaction
	database.DB.Table("transactions").
		Select("id, transactions.from, transactions.to, amount").
		Where(interfaces.Transaction{From: id}).
		Or(interfaces.Transaction{To: id}).
		Scan(&transactions)
	return transactions
}

func GetMyTransactions(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		var accounts []interfaces.ResponseAccount
		database.DB.Table("accounts").
			Select("id, name, balance").
			Where("user_id = ? ", id).
			Scan(&accounts)

		var transactions []interfaces.ResponseTransaction
		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionsByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		var response = map[string]interface{}{"message": "OK"}
		response["data"] = transactions
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
