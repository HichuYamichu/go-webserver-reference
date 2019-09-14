package user

import (
	"github.com/hichuyamichu/go-webserver-reference/controllers/base"
)

type UserController struct {
	base.Controller
	test string
}

func NewUserController() *UserController {
	return &UserController{}
}
