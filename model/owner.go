package model

import (
	"github.com/jinzhu/gorm"
	"keymanager/configuration"
)

type OwnerDetails struct {
	Name string   `gorm:"not null;unique;primary_key"`
	Admin bool     `gorm:"not null"`
	Crypto bool    `gorm:"not null"`
	Password string `gorm:"not null"`
}


type OwnerModel struct {
	DB *gorm.DB
}


func (k *OwnerModel) Connect() (*gorm.DB, error) {

	dbusr := configuration.ConfMap["DBUSR"]
	dbpassword := configuration.ConfMap["DBPASSWORD"]
	host := configuration.ConfMap["DBHOST"]
	port := configuration.ConfMap["DBPORT"]
	database := configuration.ConfMap["DBNAME"]

	connsting := dbusr + ":" + dbpassword +"@tcp("+host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", connsting)
	if err != nil {
		return nil, err
	}
	k.DB = db
	return db, nil
}


func (k *OwnerModel) Close()  {
	if k.DB != nil {
		k.DB.Close()
	}

}


func (k *OwnerModel) DropTable() error {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}
	k.DB = k.DB.DropTable(&OwnerDetails{})

	if k.DB.Error != nil {
		return k.DB.Error
	}

	return  nil

}

func (k *OwnerModel) CreateTable() error {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}
	k.DB = k.DB.CreateTable(&OwnerDetails{})

	if k.DB.Error != nil {
		return k.DB.Error
	}
	return nil
}

func (k *OwnerModel) Insert(v interface{}) error {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}
	s := v.(*OwnerDetails)

	k.DB = k.DB.Create(s)

	if k.DB.Error != nil {
		return k.DB.Error
	}

	return nil

}

func (k *OwnerModel) SelectAll() ( error , []OwnerDetails) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err, nil
	}

	s := make([]OwnerDetails,0)

	k.DB = k.DB.Select("name, admin, crypto").Find(&s)

	if k.DB.Error != nil {
		return k.DB.Error, nil
	}

	return nil, s
}


func (k *OwnerModel) Delete(owner string) (error) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}

	k.DB = k.DB.Where("name = ?",owner).Delete(OwnerDetails{})

	if k.DB.Error != nil {
		return k.DB.Error
	}

	return nil;
}


func (k *OwnerModel) CheckIfPresent(owner string) (error, *OwnerDetails) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err, nil
	}

	kp := &OwnerDetails{}

	k.DB = k.DB.Find(kp, OwnerDetails{Name: owner})

	if k.DB.Error != nil {
		return k.DB.Error, nil
	}


	return nil, kp;
}



func (k *OwnerModel) Verify(owner string , password string) bool {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return false
	}

	kp := &OwnerDetails{}



	k.DB = k.DB.Find(kp, OwnerDetails{Name: owner, Password:password})

	if k.DB.Error != nil {
		return false
	}

	return true
}