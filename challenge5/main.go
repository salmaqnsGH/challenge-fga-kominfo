package main

import (
	router "belajar-gin/routers"
)

const PORT = ":4000"

func main() {
	router.StartServer().Run(PORT)
}
