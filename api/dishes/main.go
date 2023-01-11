package dishes

import (
	"example/service-hiwjung-project/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	dish := route.Group("/dish")
	dish.POST("/", func(c *gin.Context) {
		var dish responses.DishRequestCreate
		if err := c.BindJSON(&dish); err != nil {
			HandleBadRequest(c, err)
		}

		result, err := CreateDish(dish)
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
