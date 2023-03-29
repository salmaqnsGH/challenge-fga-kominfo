package main

import (
	"belajar-gin/database"
	router "belajar-gin/routers"

	_ "github.com/lib/pq"
)

const (
	PORT = ":8080"
)

func main() {
	db := database.StartDB()

	router.StartServer(db).Run(PORT)
}
