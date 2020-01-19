package store

import (
	"github.com/hichuyamichu/go-webserver-reference/store/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	conn, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=go-web-ref password=changeme sslmode=disable")
	if err != nil {
		panic(err)
	}
	conn.LogMode(true)
	conn.AutoMigrate(&models.User{})
	db = conn
}
