package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	errResponse "request-logger/helpers/response/error"
	scsResponse "request-logger/helpers/response/success"
	"request-logger/helpers/validation"
	requestRepository "request-logger/repositories/request"
	requestTypes "request-logger/requests/request"
	"strconv"
	"time"
)

type RequestController struct{}

func (h RequestController) Store(c *gin.Context) {

	var requestBody requestTypes.StoreProcessRequestValidation

	if err := c.BindJSON(&requestBody); err != nil {
		log.Fatal("Error when binding parameters as json")
	}

	request := requestTypes.StoreProcessRequestValidation{
		UserId:    requestBody.UserId,
		IpAddress: requestBody.IpAddress,
		Uri:       requestBody.Uri,
		Domain:    requestBody.Domain,
	}

	validate := validator.New()

	if errs := validate.Struct(request); errs != nil {
		c.JSON(http.StatusInternalServerError, errResponse.NewErrorResponse(10, "Error", "Validation Error", validation.Descriptive(errs)))
		return
	}

	repository := requestRepository.RequestRepository{}
	err := repository.Create(&requestBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse.NewErrorResponse(5, "Error", err.Error(), []string{}))
		return
	}

	c.JSON(http.StatusOK, scsResponse.Response(make(map[string]string), make(map[string]string)))

	return
}

func (h RequestController) GetByDomain(c *gin.Context) {

	request := requestTypes.GetByDomainRequestValidation{}

	request.Domain = c.DefaultQuery("domain", "")

	date := c.DefaultQuery("date", "")

	if date != "" {

		// year-month-day
		formattedTime, err := time.Parse("2006-01-02", date)

		if err != nil {
			c.JSON(http.StatusInternalServerError, errResponse.NewErrorResponse(5, "Error", err.Error(), []string{}))
			return
		}

		request.Created = formattedTime
	}

	request.UserId = c.DefaultQuery("user", "")
	request.PerPage, _ = strconv.ParseInt(c.DefaultQuery("per-page", "10"), 0, 64)
	request.Page, _ = strconv.ParseInt(c.DefaultQuery("page", "1"), 0, 64)
	request.Sort = c.DefaultQuery("sort", "DESC")

	validate := validator.New()

	if errs := validate.Struct(request); errs != nil {
		c.JSON(http.StatusInternalServerError, errResponse.NewErrorResponse(10, "Error", "Validation Error", validation.Descriptive(errs)))
		return
	}

	repository := requestRepository.RequestRepository{}

	results, information, err := repository.GetByDomainName(&request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse.NewErrorResponse(5, "Error", err.Error(), []string{}))
		return
	}

	c.JSON(http.StatusOK, scsResponse.Response(results, information))

	return
}
