package dishes

import (
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	dish := route.Group("/dish")
	dish.POST("/", func(c *gin.Context) {
		var dish responses.DishRequestCreate
		if err := c.BindJSON(&dish); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}

		result, err := CreateDish(dish)
		utils.HandleResponse(c, result, err, 201, 500)
		return
	})
}
