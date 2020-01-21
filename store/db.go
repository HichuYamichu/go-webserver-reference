package store

import (
	"fmt"
	"os"

	"github.com/hichuyamichu/go-webserver-reference/store/models"
	userRepo "github.com/hichuyamichu/go-webserver-reference/store/user_repo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() {
	connStr := getConnectionStr()
	fmt.Println(connStr)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&models.User{})
	userRepo.InjectDB(db)
}

func getConnectionStr() string {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "go-web-ref"
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = "changeme"
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	return connStr
}
