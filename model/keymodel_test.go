package model

import (
	"testing"
	"time"
)

func TestKeyModel_Connect(t *testing.T) {

}



func TestKeyModel_DropTable(t *testing.T) {
	k := new(KeyModel)
	k.DropTable()
	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}

func TestKeyModel_CreateTable(t *testing.T) {

	k := new(KeyModel)

	k.CreateTable()

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}

}



func TestKeyModel_Insert(t *testing.T) {
	k := new(KeyModel)
	k1 := KeyProperties{}
	k1.Algorithm = "aes"
	k1.Deletable = true
	k1.Exportable = false
	k1.KeyName = "rajni"
	k1.KeyUniqueId = k.CreateAUUID(16)
	k1.LifeCycle.Activationdate = time.Now()
	k1.LifeCycle.CreationDate = time.Now()
	k1.LifeCycle.State = "active"
	k1.KeyType = "Symmetric"
	k1.OwnerName = "rajni"
	k1.PrivateKey = k.CreateAUUID(16)
	k1.PublicKey = k.CreateAUUID(16)
	k1.Size = 100
	k1.VersionNumber = 1
	k.Insert(&k1)

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}

}

func TestKeyModel_SelectKeys(t *testing.T) {
	k := new(KeyModel)

	_, a := k.SelectAll()

	for _,u := range a {
		t.Log(t.Name(),u.KeyName , u.Size , u.OwnerName, u.Algorithm)
	}

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}

func TestKeyModel_GetPublicBytes(t *testing.T) {
	k := new(KeyModel)

	_,a := k.GetPublicBytes("rajni","rajni")

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}

	t.Log(t.Name(),string(a))


}

func TestKeyModel_GetPrivateBytes(t *testing.T) {
	k := new(KeyModel)

	_,a := k.GetPrivateBytes("rajni","rajni")

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}

	t.Log(t.Name(),string(a))


}