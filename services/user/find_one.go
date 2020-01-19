package user

import "github.com/hichuyamichu/go-webserver-reference/store/models"

func (UserService) FindOne(id int) *models.User {
	return userRepo.FindOneByID(id)
}
