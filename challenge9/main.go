package main

import (
	database "latihan-jwt/database"
	"latihan-jwt/router"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}
