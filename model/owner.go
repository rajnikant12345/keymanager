package model

import "github.com/jinzhu/gorm"

type OwnerDetails struct {
	Name string
	Admin bool
	Crypto bool
}


type OwnerModel struct {
	DB *gorm.DB
}


func (k *OwnerModel) Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@tcp(0.0.0.0:3456)/Keymanager?charset=utf8&parseTime=True&loc=Local")
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


func (k *OwnerModel) DropTable() {
	k.Connect()
	defer k.Close()
	k.DB = k.DB.DropTable(&OwnerDetails{})

}

func (k *OwnerModel) CreateTable() error {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}
	k.DB = k.DB.CreateTable(&OwnerDetails{})
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

	k.DB = k.DB.Find(&s)

	return nil, s
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

	return nil, kp;
}