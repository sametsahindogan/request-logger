package main

import (
	"request-logger/database"
	"request-logger/server"
)

func main() {
	database.Initialize()
	server.Initialize()
}
