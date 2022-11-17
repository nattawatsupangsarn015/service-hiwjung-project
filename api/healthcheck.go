package api

import (
	"github.com/gin-gonic/gin"
)

func healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
