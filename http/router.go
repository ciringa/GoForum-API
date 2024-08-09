package sourceHttp

import (
	"github.com/ciringa/GoForum-API/http/controlers"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r*gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.POST("/room", controlers.CreateRoom)
	r.POST("/message",controlers.PostMessage)
}