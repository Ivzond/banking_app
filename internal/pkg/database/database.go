package database

import (
	"fintech_app/internal/pkg/helpers"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bankapp password=12345678")
	if err != nil {
		helpers.HandleErr(err)
	}
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
