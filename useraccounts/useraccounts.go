package useraccounts

import (
	"fintech_app/database"
	"fintech_app/helpers"
	"fintech_app/interfaces"
	"fintech_app/transactions"
	"fmt"
)

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	database.DB.Where("id = ?", id).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = account.Balance
	return responseAcc
}

func getAccount(id uint) *interfaces.Account {
	account := &interfaces.Account{}
	if database.DB.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)
	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.ID != userId {
			return map[string]interface{}{"message": "Your are not the owner of the account"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Not enough money on the account"}
		}
		updatedAccount := updateAccount(from, int(fromAccount.Balance)-amount)
		updateAccount(to, int(toAccount.Balance)+amount)

		transactions.CreateTransaction(from, to, amount)

		var response = map[string]interface{}{"message": "OK"}
		response["data"] = updatedAccount
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
