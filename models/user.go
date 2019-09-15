package models

// User : represents a user
type User struct {
	ID   int `gorm:"AUTO_INCREMENT, primary_key"`
	Name string
	Age  int
}
