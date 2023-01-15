package chatbot

import (
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	chatbot := route.Group("/chatbot")

	chatbot.POST("/", func(c *gin.Context) {
		var chatbot responses.RequestCreateChatbot
		if err := c.BindJSON(&chatbot); err != nil {
			utils.HandleBadRequest(c, err)
			return
		}
		result, err := CreateChatbot(chatbot)
		utils.HandleResponse(c, result, err, 200, 500)
		return
	})
}
