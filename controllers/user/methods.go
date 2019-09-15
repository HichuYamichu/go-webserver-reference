package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/hichuyamichu/go-webserver-reference/models"
)

func (c *Controller) sendResponce(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}

// GetUser : Returns JSON with user
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) (int, error) {
	userID := chi.URLParam(r, "userID")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return 500, err
	}

	user := c.dao.getUser(id)

	err = c.sendResponce(w, user)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

// GetAllUsers : Returns JSON with all users
func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) (int, error) {
	user := c.dao.getAllUsers()
	err := c.sendResponce(w, user)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

// InsertUser : Inserts new user to database
func (c *Controller) InsertUser(w http.ResponseWriter, r *http.Request) (int, error) {
	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 400, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return 400, err
	}

	c.dao.createUser(&user)
	return 200, nil
}

// UpdateUser : Updates user info
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) (int, error) {
	var user models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 400, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return 400, err
	}

	userID := chi.URLParam(r, "userID")
	id, err := strconv.Atoi(userID)
	c.dao.updateUser(&id, &user)

	return 200, nil
}

// DeleteUser : Removes user from database
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) (int, error) {
	userID := chi.URLParam(r, "userID")
	c.dao.deleteUser(&userID)
	return 200, nil
}
