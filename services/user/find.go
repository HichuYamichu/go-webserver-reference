package user

import "github.com/hichuyamichu/go-webserver-reference/store/models"

func (UserService) Find(skip int, take int) []*models.User {
	return userRepo.FindWithSkipAndTake(skip, take)
}
