package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	"keymanager/keys/aes"
)

type CreateKeyStruct struct {

	Keyname string
	KeySize int
	KeyOwner string
	KeyType string
	Algorithm string
	Deletable bool
	Exportable bool

}

func CreateSymmetricKey(c *CreateKeyStruct) error {
	_,e := aes.CreateAESKey(c.Keyname, c.KeySize)
	if e != nil {
		return e
	}
	return nil
}


func CreateASymmetricKey(c *CreateKeyStruct) {

}



func CreateKey(c echo.Context) error {

	input := json.NewDecoder(c.Request().Body)

	keyinput := CreateKeyStruct{}

	input.Decode(&keyinput)

	switch keyinput.KeyType {
	case "S":
		e := CreateSymmetricKey(&keyinput)
		if e != nil {
			return c.String(http.StatusNotAcceptable, e.Error())
		}
	case "A":
		CreateASymmetricKey(&keyinput)
	default:
		return c.String(http.StatusNotAcceptable, "invalid key type")
	}
	return c.String(http.StatusOK, "Success")
}