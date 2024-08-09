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

func ReturnMessages(ctx*gin.Context){
	var messages[] models.Message
	
	lib.DB.Find(&messages)

	ctx.JSON(200,gin.H{
		"message":"Successfully returned Message list!",
		"content":messages,
	})
}

func ReturnSpecificMessage(ctx*gin.Context){
	//get id 
	Id:= ctx.Param("Id")
	//store var
	var message models.Message
	lib.DB.First(&message, Id)

	ctx.JSON(200,gin.H{
		"message":"successfully returned specified message",
		"content":message,
	})
}

func ReturneMessageListByRoomId(ctx*gin.Context){
	//get roomId
	roomId := ctx.Param("RoomId")
	var messages[] models.Message
	var querylist[] models.Message
	lib.DB.Find(&messages)
	for i := 0; i < len(messages); i++ {
		if(messages[i].RoomId==roomId){
			querylist = append(querylist, messages[i])
		}
	}

	ctx.JSON(200, gin.H{
		"Message":"successfully returned message list from specified room",
		"Content":querylist,
	})
}