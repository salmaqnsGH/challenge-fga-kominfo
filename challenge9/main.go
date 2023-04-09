package main

import (
	database "latihan-jwt/database"
	"latihan-jwt/router"
)

func main() {
	database.StartDB()

	db := database.GetDB()
	router.New(db).Run(":3000")
}
