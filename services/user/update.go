package user

import "github.com/hichuyamichu/go-webserver-reference/store/models"

func (UserService) Update(user *models.User) {
	userRepo.UpdateOne(user)
}
