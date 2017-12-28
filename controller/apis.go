package controller

import (
	"github.com/labstack/echo"
)


func CreateKeyApi(c echo.Context) error {
	return CreateKey(c)
}


func LoginApi(c echo.Context) error {
	return nil
}


func CreateUser(c echo.Context) error {
	return nil
}

