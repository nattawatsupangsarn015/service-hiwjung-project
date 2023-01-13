package ingredients

import (
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Routes(route *gin.Engine) {
	ingredients := route.Group("/ingredients")
	ingredient := route.Group("/ingredient")

	ingredients.GET("/", func(c *gin.Context) {
		result := GetAllIngredients()
		c.JSON(http.StatusOK, gin.H{"data": result})
		return
	})

	ingredient.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := GetIngredientById(id)
		c.JSON(http.StatusOK, gin.H{"data": result})
		return
	})

	ingredient.GET("/by-name/:name", func(c *gin.Context) {
		name := c.Param("name")
		result, err := GetIngredientByName(name)
		utils.HandleResponse(c, result, err, 200, 500)
		return
	})

	ingredient.POST("/", func(c *gin.Context) {
		var ingredient responses.IngredientRequestCreate
		if err := c.BindJSON(&ingredient); err != nil {
			utils.HandleBadRequest(c, err)
		}

		result, err := CreateIngredient(ingredient)
		utils.HandleResponse(c, result, err, 201, 500)
		return
	})

	ingredient.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var ingredient responses.IngredientRequestUpdate
		if err := c.BindJSON(&ingredient); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}

		result, err := UpdateIngredientById(id, ingredient)
		utils.HandleResponse(c, result, err, 201, 500)
		return
	})

	ingredient.PUT("/activate/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := ActivateIngredientById(id)
		utils.HandleResponse(c, bson.M{"message": "OK"}, err, 200, 500)
		return
	})
}
