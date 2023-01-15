package api

import (
	"example/service-hiwjung-project/api/admins"
	"example/service-hiwjung-project/api/chatbot"
	"example/service-hiwjung-project/api/dishes"
	"example/service-hiwjung-project/api/line"
	"example/service-hiwjung-project/api/menus"
	"example/service-hiwjung-project/utils"

	"example/service-hiwjung-project/api/ingredients"

	"github.com/gin-gonic/gin"
)

func Router(NODE_ENV string) *gin.Engine {
	if NODE_ENV != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/healthcheck", healthcheck)

	admins.Routes(router)
	line.Routes(router)

	router.Use(utils.JwtMiddleware())
	chatbot.Routes(router)
	menus.Routes(router)
	ingredients.Routes(router)
	dishes.Routes(router)
	return router
}
