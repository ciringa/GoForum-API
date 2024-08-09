package controlers

import (
	"fmt"
	"github.com/ciringa/GoForum-API/lib"
	"github.com/ciringa/GoForum-API/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
func CreateRoom(ctx *gin.Context) {
	//get data from req body
	var body struct{
		Title string
	}
	//aply request body to 'body' var
	ctx.Bind(&body)
	if(body.Title == ""){
		panic("cant get body")
	}
	Room := models.Rooms{Id:uuid.NewString(),Title:body.Title}
	result := lib.DB.Create(&Room)
	fmt.Print(result)
	ctx.JSON(201, gin.H{
		"message":"Successfully created a room!",
		"content":Room,
	})
}