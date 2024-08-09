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
	
	r.GET("/messages",controlers.ReturnMessages)
	r.GET("/rooms",controlers.ReturnRooms)
	r.GET("/room/:Id",controlers.ReturnSpecificRoom)
	r.GET("/message/:Id",controlers.ReturnSpecificMessage)
	r.GET("/messages/:RoomId",controlers.ReturneMessageListByRoomId)
}