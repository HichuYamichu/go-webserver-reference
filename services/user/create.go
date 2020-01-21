package user

import (
	"github.com/hichuyamichu/go-webserver-reference/store/models"
	userRepo "github.com/hichuyamichu/go-webserver-reference/store/user_repo"
)

func Create(user *models.User) {
	userRepo.Save(user)
}
