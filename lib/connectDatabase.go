package lib
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)
var DB *gorm.DB 
var err error
func DatabaseLoad(){
	dsn := "host=localhost user=postgres password=123456789 dbname=wsrs port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("shomeshit happens",DB,err)
	}
}