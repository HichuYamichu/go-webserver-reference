package user_repo

import (
	"github.com/hichuyamichu/go-webserver-reference/store/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InjectDB(dbInstance *gorm.DB) {
	db = dbInstance
}

func FindOneByID(id int) *models.User {
	user := &models.User{ID: id}
	db.First(user)
	return user
}

func FindWithSkipAndTake(offset int, limit int) []*models.User {
	users := make([]*models.User, limit)
	db.Offset(offset).Limit(limit).Find(&users)
	return users
}

func Save(user *models.User) {
	db.Create(user)
}

func DeleteOneByID(id int) {
	user := &models.User{ID: id}
	db.Delete(user)
}

func UpdateOne(user *models.User) {
	db.Save(user)
}
