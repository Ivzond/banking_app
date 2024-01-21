package main

import (
	"fintech_app/api"
	"fintech_app/database"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}
