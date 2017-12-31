package controller

import (
	"errors"
	"keymanager/model"
	"strconv"
	"keymanager/keys/aes"
	"time"
)

type AesHelper struct {

}


func (a *AesHelper)ValidateKeySize(i interface{}) error {

	m := i.(map[string]string)

	switch m["keysize"] {
	case "128":
	case "192":
	case "256":
	default:
		return errors.New("Invalid keysize must be 128, 192,256")
	}

	return nil
}

func (a *AesHelper)ValidateDeletable(i interface{}) error {

	m := i.(map[string]string)

	switch m["deletable"] {
	case "true":
	case "false":
	default:
		return errors.New("deletable must be true, false")
	}

	return nil

}


func (a *AesHelper)ValidateExportable(i interface{}) error {
	m := i.(map[string]string)

	switch m["exportable"] {
	case "true":
	case "false":
	default:
		return errors.New("exportable must be true, false")
	}

	return nil
}


func (p *AesHelper) CreateKey(i interface{}) error{
	m := i.(map[string]string)

	a :=  model.KeyProperties{}

	a.CustomAttributes = m["custom"]
	a.KeyName = m["keyname"]
	a.Algorithm = m["algorithm"]
	a.Deletable,_ = strconv.ParseBool(m["deletable"])
	a.Exportable,_ = strconv.ParseBool(m["exportable"])
	a.KeyType = m["keytype"]
	a.OwnerName = m["owner"]
	tmp,_ := strconv.ParseInt(m["keysize"],10,32)
	a.Size = int(tmp)
	t,e := aes.CreateAESKey(a.KeyName,a.Size/8)

	if e != nil {
		return e
	}
	a.PublicKey = t.RawBytes

	a.PrivateKey = t.RawBytes
	a.LifeCycle.Activationdate = time.Now()
	a.LifeCycle.CreationDate = time.Now()
	a.VersionNumber = 1

	mod := &model.KeyModel{}

	e = mod.Insert(&a)

	if e != nil {
		return e
	}

	return nil

}

func (a *AesHelper) ValidateOwner(i interface{}) error {
	return nil
}


