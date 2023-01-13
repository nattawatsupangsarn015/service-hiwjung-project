package menus

import (
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	menus := route.Group("/menus")
	menu := route.Group("/menu")
	menus.GET("/", func(c *gin.Context) {
		result := GetAllMenus()
		c.JSON(http.StatusOK, result)
	})

	menu.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, id)
	})

	menu.POST("/", func(c *gin.Context) {
		var menu responses.MenuRequestCreate

		if err := c.BindJSON(&menu); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}

		result, err := CreateMenu(menu)
		utils.HandleResponse(c, result, err, 201, 500)
		return
	})
}
