package controller

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"keymanager/model"
	"net/http"
	"time"
	"fmt"
)

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	jwt.StandardClaims
}

type LoginStruct struct {
	Login    string
	Password string
}


func CreateUserApi(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Name

	m := &model.OwnerModel{}

	e,o := m.CheckIfPresent(name)

	if e != nil {
		return echo.ErrUnauthorized
	}

	if !o.Admin {
		return echo.ErrUnauthorized
	}


	inp := json.NewDecoder(c.Request().Body)

	l := model.OwnerDetails{}

	e = inp.Decode(&l)

	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}

	l.Password = fmt.Sprintf("%X",sha256.Sum256([]byte(l.Password)))

	e = m.Insert(&l)

	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}


	return c.String(http.StatusOK, "Welcome "+l.Name+"!")

}




func Login(c echo.Context) error {

	inp := json.NewDecoder(c.Request().Body)

	l := LoginStruct{}

	e := inp.Decode(&l)

	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}

	username := l.Login
	password := l.Password

	sum := sha256.Sum256([]byte(password))

	s := fmt.Sprintf("%X",sum)

	m := &model.OwnerModel{}

	if !m.Verify(username, s) {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 300).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.

	//TODO: create root of trust for this
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

	return echo.ErrUnauthorized
}
