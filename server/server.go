package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"request-logger/route"
)

func Initialize() {

	setErrorLogger()

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	server := route.NewRouter()

	fmt.Println("GIN Server started...")
	err = server.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

	if err != nil {
		log.Fatal(err)
	}

}

func setErrorLogger() {
	errorLogFile, _ := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if errorLogFile == nil {
		errorLogFile, _ = os.Create("error.log")
	}

	log.SetOutput(errorLogFile)
}
