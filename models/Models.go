package models
import "gorm.io/gorm"
type Message struct {
	gorm.Model
	Id string
	Message string
	RoomId string
	Room Rooms 
	Reaction_count int
}
type Rooms struct {
	Id string
	Title string
}