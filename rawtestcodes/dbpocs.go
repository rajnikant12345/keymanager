package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"time"
)

type KeyProperties struct {
	KeyUniqueId string	`gorm:"not null;unique;primary_key"`
	KeyName string		`gorm:"not null;unique"`
	OwnerName string	`gorm:"not null"`
	KeyType string		`gorm:"not null"`
	Algorithm string	`gorm:"not null"`
	Size int		`gorm:"not null"`
	VersionNumber int	`gorm:"not null"`
	PublicKey string	`gorm:"not null"`
	PrivateKey string	`gorm:"not null"`
	Deletable bool		`gorm:"not null"`
	Exportable bool		`gorm:"not null"`
	LifeCycle KeyLifeCycle
	CustomAttributes string `gorm:"size:2048"`
}

type KeyLifeCycle struct {
	State string			`gorm:"not null"`
	CreationDate time.Time		`gorm:"not null"`
	Activationdate time.Time	`gorm:"not null"`
	DeactivationDate time.Time
	CompromiseDate time.Time
}



func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(0.0.0.0:3456)/Keymanager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}

	db.DropTable("key_properties")

	db.CreateTable(&KeyProperties{})




	defer db.Close()
}
