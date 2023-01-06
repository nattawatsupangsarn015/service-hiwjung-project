package ingredients

import (
	"example/service-hiwjung-project/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	ingredients := route.Group("/ingredients")
	ingredient := route.Group("/ingredient")
	ingredients.GET("/", func(c *gin.Context) {
		result := GetAllIngredients()
		c.JSON(http.StatusOK, gin.H{ "data": result })
	})

	ingredient.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := GetIngredientById(id);
		c.JSON(http.StatusOK,  gin.H{ "data": result })
	})

	ingredient.POST("/", func(c *gin.Context) {
		var ingredient responses.IngredientRequestCreate

		if err := c.BindJSON(&ingredient); err != nil {
			HandleBadRequest(c, err)
		}

		result, err := CreateIngredient(ingredient)
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
