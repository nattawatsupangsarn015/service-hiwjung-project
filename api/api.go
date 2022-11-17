package api

import (
	"example/service-hiwjung-project/api/home"
	"example/service-hiwjung-project/api/products"

	"github.com/gin-gonic/gin"
)

func Router(NODE_ENV string) *gin.Engine {
	if NODE_ENV != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/healthcheck", healthcheck)

	home.Routes(router)
	products.Routes(router)
	return router
}
