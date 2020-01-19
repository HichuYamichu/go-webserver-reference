package user

import "github.com/hichuyamichu/go-webserver-reference/store/models"

func (UserService) Create(user *models.User) {
	userRepo.Save(user)
}
