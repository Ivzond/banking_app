package database

import (
	"fintech_app/internal/pkg/helpers"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=bankapp password=12345678 sslmode=disable")
	if err != nil {
		helpers.HandleErr(err)
	}
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)

	DB = database
}
