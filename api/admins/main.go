package admins

import (
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	admin := route.Group("/admin")
	admin.POST("/", func(c *gin.Context) {
		var admin responses.AdminRequestCreate
		if err := c.BindJSON(&admin); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}

		result, err := CreateAdmin(admin)
		utils.HandleResponse(c, result, err, 201, 500)
		return
	})

	admin.POST("/login", func(c *gin.Context) {
		var admin responses.AdminRequestCreate
		if err := c.BindJSON(&admin); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}

		result, err, errStatus := LoginAdmin(admin)
		utils.HandleResponse(c, result, err, 200, errStatus)
		return
	})
}
