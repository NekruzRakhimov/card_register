package main

import (
	"card_register/db"
	"card_register/routes"
	"card_register/utils"
)

func main() {
	utils.ReadSettings()
	db.StartDbConnection()

	routes.RunAllRoutes()
}
