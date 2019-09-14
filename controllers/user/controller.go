package user

import (
	"github.com/hichuyamichu/go-webserver-reference/controllers/base"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	base.Controller
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) *UserController {
	return &UserController{DB: DB}
}
