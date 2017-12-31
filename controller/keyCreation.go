package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
)

func CreateKeyApi(c echo.Context) error {

	var m  = make(map[string]string)

	e,_ := ValidateAdmin(c)

	if e != nil {
		return e
	}

	inp := json.NewDecoder(c.Request().Body)

	e = inp.Decode(&m)

	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}


	if _,ok := m["keytype"];!ok {
		return c.JSON(http.StatusOK,&m)
	}

	return c.JSON(http.StatusOK,&m)
}