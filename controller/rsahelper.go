package controller

import (
	"errors"
	"keymanager/model"
	"strconv"
	"time"
	"keymanager/keys/rsa"
)

type RsaHelper struct {

}


func (a *RsaHelper)ValidateKeySize(i interface{}) error {

	m := i.(map[string]string)

	switch m["keysize"] {
	case "1024":
	case "2048":
	case "3072":
	case "4096":
	default:
		return errors.New("Invalid keysize must be 1024, 2048,3072,4096")
	}

	return nil
}

func (a *RsaHelper)ValidateDeletable(i interface{}) error {

	m := i.(map[string]string)

	switch m["deletable"] {
	case "true":
	case "false":
	default:
		return errors.New("deletable must be true, false")
	}

	return nil

}


func (a *RsaHelper)ValidateExportable(i interface{}) error {
	m := i.(map[string]string)

	switch m["exportable"] {
	case "true":
	case "false":
	default:
		return errors.New("exportable must be true, false")
	}

	return nil
}


func (p *RsaHelper) CreateKey(i interface{}) error{
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
	t,e := rsa.CreateRSAKey(a.KeyName,a.Size)

	if e != nil {
		return e
	}
	a.PublicKey = t.Public

	a.PrivateKey = t.Private
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

func (a *RsaHelper) ValidateOwner(i interface{}) error {
	return nil
}


