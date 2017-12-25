package model

import (
	"testing"
)

func TestOwnerModel_Connect(t *testing.T) {

}



func TestOwnerModel_DropTable(t *testing.T) {
	k := new(OwnerModel)
	k.DropTable()
	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}

func TestOwnerModel_CreateTable(t *testing.T) {

	k := new(OwnerModel)

	k.CreateTable()

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}

}



func TestOwnerModel_Insert(t *testing.T) {
	k := new(OwnerModel)
	k1 := OwnerDetails{}

	k1.Name = "rajni"
	k1.Admin = true
	k1.Crypto = true

	k.Insert(&k1)

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}

}

func TestOwnerModel_SelectKeys(t *testing.T) {
	k := new(OwnerModel)

	_, a := k.SelectAll()

	for _,u := range a {
		t.Log(t.Name(),u )
	}

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}


func TestOwnerModel_CheckIfPresent(t *testing.T) {
	k := new(OwnerModel)

	k.CheckIfPresent("rajni")


	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}