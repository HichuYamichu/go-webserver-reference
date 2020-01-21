package user

import (
	"github.com/hichuyamichu/go-webserver-reference/store/models"
	userRepo "github.com/hichuyamichu/go-webserver-reference/store/user_repo"
)

func Find(skip int, take int) []*models.User {
	return userRepo.FindWithSkipAndTake(skip, take)
}
