package handlers

import (
	"fmt"
	"strconv"

	"github.com/hichuyamichu/go-webserver-reference/services/user"
	"github.com/hichuyamichu/go-webserver-reference/store/models"
	"github.com/labstack/echo"
)

var userService = user.UserService{}

func FindUser(c echo.Context) error {
	idParam := c.Param("id")
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	user := userService.FindOne(id)
	c.JSON(200, user)
	return nil
}

func FindUsers(c echo.Context) error {
	skipParam := c.QueryParam("skip")
	skip, err := strconv.Atoi(skipParam)
	if err != nil {
		return err
	}
	takeParam := c.QueryParam("take")
	take, err := strconv.Atoi(takeParam)
	if err != nil {
		return err
	}
	users := userService.Find(skip, take)
	c.JSON(200, users)
	return nil
}

func CreateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	userService.Create(user)
	c.JSON(200, user)
	return nil
}

func DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	userService.Delete(id)
	return nil
}

func UpdateUser(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	userService.Update(user)
	c.JSON(200, user)
	return nil
}
