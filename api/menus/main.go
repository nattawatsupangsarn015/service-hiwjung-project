package menus

import (
	"example/service-hiwjung-project/responses"
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
			HandleBadRequest(c, err)
		}

		result, err := CreateMenu(menu)
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
