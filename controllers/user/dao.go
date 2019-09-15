package user

import (
	"strconv"

	"github.com/hichuyamichu/go-webserver-reference/models"
	"github.com/jinzhu/gorm"
)

type dao struct {
	db *gorm.DB
}

func (d *dao) getUser(id int) models.User {
	var user models.User
	d.db.Find(&user, id)
	return user
}

func (d *dao) getAllUsers() []models.User {
	var users []models.User
	d.db.Find(&users)
	return users
}

func (d *dao) createUser(user *models.User) {
	d.db.Create(user)
}

func (d *dao) updateUser(userID *int, user *models.User) {
	user.ID = *userID
	d.db.Update(user)
}

func (d *dao) deleteUser(userID *string) {
	var user models.User
	user.ID, _ = strconv.Atoi(*userID)
	d.db.Delete(user)
}
