package store

import "github.com/hichuyamichu/go-webserver-reference/store/models"

type UserRepository struct{}

func (UserRepository) FindOneByID(id int) *models.User {
	user := &models.User{ID: id}
	db.First(user)
	return user
}

func (UserRepository) FindWithSkipAndTake(offset int, limit int) []*models.User {
	users := make([]*models.User, limit)
	db.Offset(offset).Limit(limit).Find(&users)
	return users
}

func (UserRepository) Save(user *models.User) {
	db.Create(user)
}

func (UserRepository) DeleteOneByID(id int) {
	user := &models.User{ID: id}
	db.Delete(user)
}

func (UserRepository) UpdateOne(user *models.User) {
	db.Save(user)
}
