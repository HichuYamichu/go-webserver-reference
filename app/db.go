package app

import (
	"github.com/hichuyamichu/go-webserver-reference/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go password=changeme sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}
