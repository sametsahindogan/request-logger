package route

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"request-logger/controllers/request"
	errResponse "request-logger/helpers/response/error"
)

func NewRouter() *gin.Engine {

	routes := bootstrapGinFramework()

	requestController := new(request.RequestController)

	routes.POST("/", basicAuthentication, requestController.Store)
	routes.GET("/list", basicAuthentication, requestController.GetByDomain)

	return routes
}

func bootstrapGinFramework() *gin.Engine {
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

	return routes
}

func basicAuthentication(c *gin.Context) {
	_ = godotenv.Load()

	apiKey, apiSecret, hasAuth := c.Request.BasicAuth()

	if hasAuth && apiKey == os.Getenv("API_KEY") && apiSecret == os.Getenv("API_SECRET") {
		c.Next()
	} else {
		c.Abort()
		c.JSON(http.StatusInternalServerError, errResponse.NewErrorResponse(5, "Error", "Authentication failure.", []string{}))
		return
	}
}
