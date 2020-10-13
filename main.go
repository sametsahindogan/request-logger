package main

import (
	"request-logger/bootstrap/database"
	"request-logger/bootstrap/server"
)

func main() {
	database.Initialize()
	server.Initialize()
}