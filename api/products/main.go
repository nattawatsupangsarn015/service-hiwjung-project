package products

import (
	"example/service-hiwjung-project/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	products := route.Group("/products")
	products.GET("/", func(c *gin.Context) {
		result := GetAllProducts()
		c.JSON(http.StatusOK, result)
	})

	products.POST("/", func(c *gin.Context) {
		product := responses.ProductRequestCreate{}
		HandleBadRequest(c, product)
		result, err := CreateProducts(product)
		HandleResponse(c, result, err, 201, 500)
	})
}

func HandleBadRequest(c *gin.Context, structure interface{}) {
	if err := c.BindJSON(&structure); err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
}

func HandleResponse(c *gin.Context, response interface{}, err error, status int, statusError int) {
	if err != nil {
		c.JSON(statusError, err)
	} else {
		c.JSON(status, response)
	}
}
