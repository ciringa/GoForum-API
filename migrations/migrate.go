package main

//use this to setup the tables in our database
import (
	"github.com/ciringa/GoForum-API/lib"
	"github.com/ciringa/GoForum-API/models"
)
func init(){
	lib.LoadEnv()
	lib.DatabaseLoad()
}
func main(){

	lib.DB.AutoMigrate(&models.Message{},&models.Rooms{})
	
}