package users

import (
	"errors"
	"fintech_app/helpers"
	"fintech_app/interfaces"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func prepareToken(user *interfaces.User) string {
	// Sign token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(30 * time.Minute).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)

	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	if err != nil {
		helpers.HandleErr(err)
	}

	return token
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount, withToken bool) map[string]interface{} {
	// Setup response
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	// Prepare response
	var response = map[string]interface{}{"message": "OK"}
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser
	return response
}

func Login(username string, pass string) map[string]interface{} {
	// Validation to login
	valid := helpers.Validation(
		[]interfaces.Validation{
			{
				Value: username,
				Valid: "username",
			},
			{
				Value: pass,
				Valid: "password",
			},
		},
	)
	if valid {
		// Connect db
		db := helpers.ConnectDB()
		user := &interfaces.User{}

		// Check if the user exists
		if db.Where("username = ?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}
		// Verify password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if errors.Is(passErr, bcrypt.ErrMismatchedHashAndPassword) && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}
		// Found accounts for the user
		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {
				helpers.HandleErr(err)
			}
		}(db)

		var response = prepareResponse(user, accounts, true)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

func Register(username string, email string, pass string) map[string]interface{} {
	// Validation to register
	valid := helpers.Validation(
		[]interfaces.Validation{
			{
				Value: username,
				Valid: "username",
			},
			{
				Value: email,
				Valid: "email",
			},
			{
				Value: pass,
				Valid: "password",
			},
		},
	)
	if valid {
		db := helpers.ConnectDB()
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{
			Username: username,
			Email:    email,
			Password: generatedPassword,
		}
		db.Create(&user)

		account := interfaces.Account{
			Type:    "Daily account",
			Name:    username + "'s " + "account",
			Balance: 0,
			UserID:  user.ID,
		}
		db.Create(&account)
		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {
				helpers.HandleErr(err)
			}
		}(db)

		accounts := []interfaces.ResponseAccount{}
		respAccount := interfaces.ResponseAccount{
			ID:      account.ID,
			Name:    account.Name,
			Balance: account.Balance,
		}
		accounts = append(accounts, respAccount)
		var response = prepareResponse(user, accounts, true)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)
	// Find and return user
	if isValid {
		db := helpers.ConnectDB()
		user := &interfaces.User{}
		if db.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}
		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		defer func(db *gorm.DB) {
			err := db.Close()
			if err != nil {
				helpers.HandleErr(err)
			}
		}(db)

		var response = prepareResponse(user, accounts, false)
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
