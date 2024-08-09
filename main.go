package main

import (
	"github.com/ciringa/GoForum-API/http"

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
	sourceHttp.LoadRoutes(r)
	r.Run()
}