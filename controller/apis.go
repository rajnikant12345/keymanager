package controller

import (
	"github.com/labstack/echo"
)


func CreateKeyApi(c echo.Context) error {
	return CreateKey(c)
}


func DropKeyApi(c echo.Context) error {
	return DropKey(c)
}


func LoginApi(c echo.Context) error {
	return Login(c)
}


func CreateUser(c echo.Context) error {
	return CreateUserApi(c)
}

func DeleteUser(c echo.Context) error {
	return DropUserApi(c)
}

func ListUsers(c echo.Context) error {
	return ListUsersApi(c)
}

func UpdateUser(c echo.Context) error {
	return UpdateUserApi(c)
}


