package route

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"request-logger/controllers/request"
)

func NewRouter() *gin.Engine {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	file, _ := os.OpenFile("gin.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if file == nil {
		file, _ = os.Create("gin.log")
	}

	gin.DefaultWriter = io.MultiWriter(file)

	gin.SetMode(os.Getenv("GIN_MODE"))

	routes := gin.New()

	routes.Use(gin.Logger())

	routes.Use(gin.Recovery())

	requestController := new(request.RequestController)

	routes.POST("/", requestController.Store)
	routes.GET("/list", requestController.GetByDomain)

	return routes
}
