package model



import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"io"
	"crypto/rand"
)

type KeyProperties struct {
	KeyUniqueId []byte	`gorm:"not null;unique;primary_key"`
	KeyName string		`gorm:"not null;unique"`
	OwnerName string	`gorm:"not null"`
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

type KeyLifeCycle struct {
	State string			`gorm:"not null"`
	CreationDate time.Time		`gorm:"not null"`
	Activationdate time.Time	`gorm:"not null"`
	DeactivationDate time.Time
	CompromiseDate time.Time
}


type KeyModel struct {
	DB *gorm.DB
}

func (k *KeyModel)CreateAUUID (size int) ([]byte ) {
	rawBytes := make([]byte,size)
	io.ReadFull(rand.Reader,rawBytes)
	return rawBytes

}


func (k *KeyModel) Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@tcp(0.0.0.0:3456)/Keymanager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	k.DB = db
	return db, nil
}


func (k *KeyModel) Close()  {
	if k.DB != nil {
		k.DB.Close()
	}

}

func (k *KeyModel) DropTable() {
	k.Connect()
	defer k.Close()
	k.DB = k.DB.DropTable(&KeyProperties{})

}

func (k *KeyModel) CreateTable() error {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}
	k.DB = k.DB.CreateTable(&KeyProperties{})
	return nil
}

func (k *KeyModel) Insert(v interface{}) error {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}


	s := v.(*KeyProperties)
	k.DB = k.DB.Create(s)
	return nil

}

func (k *KeyModel) SelectAll() ( error , []KeyProperties) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err, nil
	}

	s := make([]KeyProperties,0)

	k.DB = k.DB.Find(&s)

	return nil, s
}

func (k *KeyModel) GetPrivateBytes(keyName string , owner string) (error, []byte) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err, nil
	}

	kp := &KeyProperties{}

	k.DB = k.DB.Find(kp, KeyProperties{KeyName: keyName , OwnerName:owner})

	return nil, kp.PrivateKey;

}



func (k *KeyModel) GetPublicBytes(keyName string , owner string) (error, []byte) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err, nil
	}

	kp := &KeyProperties{}

	k.DB = k.DB.Find(kp, KeyProperties{KeyName: keyName , OwnerName:owner})

	return nil, kp.PublicKey;

}

