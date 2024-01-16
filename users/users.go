package users

import (
	"errors"
	"fintech_app/helpers"
	"fintech_app/interfaces"
	"github.com/dgrijalva/jwt-go"
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
	helpers.HandlerErr(err)

	return token
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount) map[string]interface{} {
	// Setup response
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	// Prepare response
	var token = prepareToken(user)
	var response = map[string]interface{}{"message": "Authentication is fine"}
	response["jwt"] = token
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

		defer db.Close()

		var response = prepareResponse(user, accounts)

		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

func Register(username string, email string, pass string) map[string]interface{} {

}
