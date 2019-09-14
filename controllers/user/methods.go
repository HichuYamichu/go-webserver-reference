package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hichuyamichu/go-webserver-reference/models"
)

// GetUsers : Returns JSON with all users
func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) (int, error) {
	var user models.User
	uc.DB.First(&user)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

// InsertUser : Inserts new user to database
func (uc *UserController) InsertUser(w http.ResponseWriter, r *http.Request) (int, error) {
	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 400, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return 400, err
	}

	uc.DB.NewRecord(user)
	uc.DB.Create(&user)

	return 200, nil
}

// UpdateUser : Updates user info
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) (int, error) {
	var user models.User
	uc.DB.First(&user)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 400, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return 400, err
	}

	uc.DB.Save(&user)

	return 200, nil
}

// DeleteUser : Removes user from database
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) (int, error) {
	var user models.User
	uc.DB.First(&user)
	uc.DB.Delete(&user)
	return 200, nil
}
