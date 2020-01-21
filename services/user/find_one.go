package user

import (
	"github.com/hichuyamichu/go-webserver-reference/store/models"
	userRepo "github.com/hichuyamichu/go-webserver-reference/store/user_repo"
)

func FindOne(id int) *models.User {
	return userRepo.FindOneByID(id)
}
