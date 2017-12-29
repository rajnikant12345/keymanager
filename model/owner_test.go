package model

import (
	"testing"
	"os"
	"keymanager/configuration"
)

func TestOwnerModel_Connect(t *testing.T) {

	os.Setenv("DBNAME","Keymanager")
	os.Setenv("DBUSR","root")
	os.Setenv("DBPASSWORD","root")
	os.Setenv("DBHOST","0.0.0.0")
	os.Setenv("DBPORT","3456")


	configuration.InitializeConfiguration()

}



func TestOwnerModel_DropTable(t *testing.T) {

	//t.SkipNow()

	k := new(OwnerModel)
	e := k.DropTable()


	if e != nil {
		t.Log(t.Name(),e.Error())
		t.Fail()
	}


	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}

func TestOwnerModel_CreateTable(t *testing.T) {

	//t.SkipNow()
	k := new(OwnerModel)

	k.CreateTable()

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}

}



func TestOwnerModel_Insert(t *testing.T) {
	//t.SkipNow()
	k := new(OwnerModel)
	k1 := OwnerDetails{}

	k1.Name = "admin"
	k1.Admin = true
	k1.Crypto = true
	k1.Password = "312433C28349F63C4F387953FF337046E794BEA0F9B9EBFCB08E90046DED9C76"

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

	k.CheckIfPresent("admin")


	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}

func TestOwnerModel_Verify(t *testing.T) {
	k := new(OwnerModel)

	k.Verify("admin","312433C28349F63C4F387953FF337046E794BEA0F9B9EBFCB08E90046DED9C76")


	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}






func TestOwnerModel_Delete(t *testing.T) {

	t.SkipNow()

	k := new(OwnerModel)

	e := k.Delete("admin")

	if e != nil {
		t.Log(t.Name(),e.Error())
		t.Fail()
	}

	if k.DB.Error != nil {
		t.Log(t.Name(),k.DB.Error.Error())
		t.Fail()
	}
}