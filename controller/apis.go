package controller

import (
	"github.com/labstack/echo"
)



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


func CreateKey(c echo.Context) error {
	return CreateKeyApi(c)
}

func DeleteKey(c echo.Context) error {
	return DeleteKeyApi(c)
}