package main

import (
	"net/http"
	"github.com/ciringa/GoForum-API/lib"
	"github.com/gin-gonic/gin"

)

//runs before main, so it can be used to initialize things
func init(){
	lib.LoadEnv()
	lib.DatabaseLoad()
}
func main(){

	r:= gin.Default()
	r.GET("/ping",func (c*gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message":"pong!",
		})
	})

	r.POST("/create",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"Description":"created",
		})
	})

	r.Run()

}