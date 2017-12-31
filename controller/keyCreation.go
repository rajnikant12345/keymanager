package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"encoding/json"
	"keymanager/kmerrors"
	"errors"
)

/*
type KeyProperties struct {
	Id uint			`gorm:"primary_key"`
	KeyName string		`gorm:"not null;unique"`
	OwnerName string	`gorm:"not null;unique"`
	KeyType string		`gorm:"not null"`
	Algorithm string	`gorm:"not null"`
	Size int		`gorm:"not null"`
	VersionNumber int	`gorm:"not null"`
	PublicKey []byte	`gorm:"size:4096"`
	PrivateKey []byte	`gorm:"size:4096"`
	Deletable bool		`gorm:"not null"`
	Exportable bool		`gorm:"not null"`
	LifeCycle KeyLifeCycle
	CustomAttributes string `gorm:"size:2048"`
}
 */




func validateSymmetricKey(m map[string]string) error {

	var v Helper

	if m["algorithm"] == "aes" {
		v = new(AesHelper)
	}else {
		return errors.New("Unsupported algo.")
	}

	e := v.ValidateKeySize(m)
	if e != nil {
		return e
	}
	e = v.ValidateDeletable(m)
	if e != nil {
		return e
	}
	e = v.ValidateExportable(m)
	if e != nil {
		return e
	}

	e = v.CreateKey(m)
	if e != nil {
		return e
	}

	return nil
}


func validateASymmetricKey(m map[string]string) error {
	return nil
}



func validatedata(m map[string]string) error{

	if  m["keytype"] != "symmetric" &&  m["keytype"] != "asymmetric" {
		return errors.New("keytype must be symmetric or asymmetric")
	}

	if m["keytype"] == "symmetric" {
		return validateSymmetricKey(m)
	}else {
		return validateASymmetricKey(m)
	}


	return nil
}




func CreateKeyApi(c echo.Context) error {

	var m  = make(map[string]string)

	e,_ := ValidateAdmin(c)

	if e != nil {
		return e
	}

	inp := json.NewDecoder(c.Request().Body)

	e = inp.Decode(&m)

	if e != nil {
		return c.JSON(http.StatusPartialContent,&kmerrors.ErrorStruct{Value:e.Error(),Action:"PLease enter valid JSON."})
	}


	if _,ok := m["keyname"];!ok {

		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"keyname missing",Action:"PLease enter valid keyname."})
	}

	if _,ok := m["keytype"];!ok {

		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"keytype missing",Action:"PLease enter valid keytype."})
	}


	if _,ok := m["keysize"];!ok {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"keysize missing",Action:"PLease enter valid keysize."})
	}


	if _,ok := m["algorithm"];!ok {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"algorithm missing",Action:"PLease enter valid algorithm."})
	}


	if _,ok := m["owner"];!ok {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"owner missing",Action:"PLease enter valid owner."})
	}


	if _,ok := m["deletable"];!ok {

		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"deletable attribute missing",
			Action:"PLease enter valid deletable attribute."})
	}


	if _,ok := m["exportable"];!ok {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:"exportable attribute missing",
			Action:"PLease enter valid exportable attribute."})
	}

	e = validatedata(m)

	if e != nil {
		return c.JSON(http.StatusBadRequest,&kmerrors.ErrorStruct{Value:e.Error(),
			Action:"PLease check error."})
	}

	return c.JSON(http.StatusOK,&m)
}