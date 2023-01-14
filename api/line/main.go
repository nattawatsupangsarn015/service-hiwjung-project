package line

import (
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	line := route.Group("/line")

	line.POST("/main", func(c *gin.Context) {
		var line responses.LineMessage
		if err := c.BindJSON(&line); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}
		result, err := ReplyUser(line)
		utils.HandleResponse(c, result, err, 200, 500)
		return
	})
}
