package user

import userRepo "github.com/hichuyamichu/go-webserver-reference/store/user_repo"

func Delete(id int) {
	userRepo.DeleteOneByID(id)
}
