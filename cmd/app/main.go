package main

import (
	"fintech_app/internal/app/api"
	"fintech_app/internal/pkg/database"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}
