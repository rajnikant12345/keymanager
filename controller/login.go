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
type jwtCustomClaims struct {
	Name  string `json:"name"`
	jwt.StandardClaims
}

type LoginStruct struct {
	Login    string
	Password string
}


func CreateUserApi(c echo.Context) error {
	return nil;
}




func Login(c echo.Context) error {

	inp := json.NewDecoder(c.Request().Body)

	l := LoginStruct{}

	e := inp.Decode(&l)

	if e != nil {
		return c.String(http.StatusNotFound, e.Error())
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
	claims := &jwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 60).Unix(),
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
