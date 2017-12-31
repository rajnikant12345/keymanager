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
	"strings"
	"keymanager/kmerrors"
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


func ValidateAdmin(c echo.Context) (error, string) {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Name

	m := &model.OwnerModel{}

	e,o := m.CheckIfPresent(name)

	if e != nil {
		return echo.ErrUnauthorized,""
	}

	if !o.Admin {
		return echo.ErrUnauthorized,""
	}

	return nil,name

}

func DropUserApi(c echo.Context) error {
	e,use := ValidateAdmin(c)

	if e != nil {
		return e
	}

	m := &model.OwnerModel{}

	inp := json.NewDecoder(c.Request().Body)

	l := model.OwnerDetails{}

	e = inp.Decode(&l)

	if e != nil {
		return c.JSON(http.StatusPartialContent,&kmerrors.ErrorStruct{Value:e.Error(),Action:"PLease enter valid JSON."})
	}

	if strings.ToUpper(l.Name) == "ADMIN" {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"Cannot Delete Admin",Action:"Delete Valid User."})
	}

	if strings.ToUpper(use) == strings.ToUpper(l.Name) {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"User Cannot Delete Itself",Action:"Delete Valid User."})
	}


	e = m.Delete(l.Name)

	if e != nil {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:e.Error(),Action:"See Error Message."})
	}

	return c.JSON(http.StatusOK,&kmerrors.ErrorStruct{Value:"Success",Action:"User"+ l.Name +"Dropped."})



}

func CreateUserApi(c echo.Context) error {

	e,_ := ValidateAdmin(c)

	if e != nil {
		return e
	}

	m := &model.OwnerModel{}

	inp := json.NewDecoder(c.Request().Body)

	l := model.OwnerDetails{}

	e = inp.Decode(&l)

	if e != nil {
		return c.JSON(http.StatusPartialContent,&kmerrors.ErrorStruct{Value:e.Error(),Action:"PLease enter valid JSON."})
	}

	l.Password = fmt.Sprintf("%X",sha256.Sum256([]byte(l.Password)))

	e = m.Insert(&l)

	if e != nil {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:e.Error(),Action:"See Error Message."})
	}


	return c.JSON(http.StatusOK,&kmerrors.ErrorStruct{Value:"Success",Action:"User"+ l.Name +"Created."})

}

func Login(c echo.Context) error {

	inp := json.NewDecoder(c.Request().Body)

	l := LoginStruct{}

	e := inp.Decode(&l)

	if e != nil {
		return echo.ErrUnauthorized
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

func ListUsersApi(c echo.Context) error {

	e,_ := ValidateAdmin(c)

	if e != nil {
		return e
	}

	m := &model.OwnerModel{}

	e,d := m.SelectAll()

	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK,&d)


}

func UpdateUserApi(c echo.Context) error {
	e,use := ValidateAdmin(c)

	if e != nil {
		return e
	}

	m := &model.OwnerModel{}

	inp := json.NewDecoder(c.Request().Body)

	l := model.OwnerDetails{}

	e = inp.Decode(&l)

	if e != nil {
		return c.String(http.StatusBadRequest, e.Error())
	}

	if strings.ToUpper(l.Name) == "ADMIN" {
		return c.String(http.StatusNotAcceptable, "Cannot update admin")
	}

	if strings.ToUpper(use) == strings.ToUpper(l.Name) {
		return c.String(http.StatusNotAcceptable, "Cannot update itself")
	}

	l.Password = fmt.Sprintf("%X",sha256.Sum256([]byte(l.Password)))

	e = m.Update(&l)

	if e != nil {
		return c.String(http.StatusNotAcceptable, e.Error())
	}

	return c.JSON(http.StatusOK,echo.Map{"Status":"Success"})
}
