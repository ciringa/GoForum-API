package controlers

import (
	"fmt"
	"github.com/ciringa/GoForum-API/lib"
	"github.com/ciringa/GoForum-API/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostMessage(ctx *gin.Context) {
	//get data from req body
	var body struct {
		Message string
		RoomId string
	}
	//aply request body to 'body' var
	ctx.Bind(&body)
	if body.Message == "" {
		panic("cant get body")
	}
	Message := models.Message{Id: uuid.NewString(),Message: body.Message, RoomId: body.RoomId,Reaction_count: 0}
	result := lib.DB.Create(&Message)
	fmt.Print(result)
	ctx.JSON(201, gin.H{
		"message": "Successfully created a Message!",
		"content": Message,
	})
}