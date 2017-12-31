package model



import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"io"
	"crypto/rand"
	"keymanager/configuration"
)

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
	k.DB = k.DB.Debug().CreateTable(&KeyProperties{}).AddForeignKey("owner_name", "owner_details(name)", "RESTRICT", "RESTRICT")
	if k.DB.Error != nil {
		return k.DB.Error
	}
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

	if k.DB.Error != nil {
		return k.DB.Error
	}

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

	if k.DB.Error != nil {
		return k.DB.Error, nil
	}

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

	if k.DB.Error != nil {
		return k.DB.Error, nil
	}

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

	if k.DB.Error != nil {
		return k.DB.Error, nil
	}


	return nil, kp.PublicKey;

}

func (k *KeyModel) Delete(key string) (error) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}

	k.DB = k.DB.Where("key_name = ?",key).Delete(KeyProperties{})

	if k.DB.Error != nil {
		return k.DB.Error
	}

	return nil;
}



func (k *KeyModel) CheckIfPresent(key string) (error, *KeyProperties) {
	defer k.Close()
	_,err := k.Connect()

	if err != nil {
		k.DB = nil
		return err, nil
	}

	kp := &KeyProperties{}

	k.DB = k.DB.Find(kp, KeyProperties{KeyName: key})

	if k.DB.Error != nil {
		return k.DB.Error, nil
	}


	return nil, kp;
}



func (k *KeyModel) Update(key *KeyProperties ) error {
	defer k.Close()

	err,kp := k.CheckIfPresent(key.KeyName)

	if err != nil {
		return err
	}

	//fmt.Println(*kp)

	if key.OwnerName != "" {
		kp.OwnerName = key.OwnerName
	}
	if key.KeyName  != "" {
		kp.KeyName = key.KeyName
	}
	if key.CustomAttributes != "" {
		kp.CustomAttributes = key.CustomAttributes
	}

	kp.Deletable = key.Deletable

	kp.Exportable = key.Exportable

	_,err = k.Connect()

	if err != nil {
		k.DB = nil
		return err
	}

	k.DB = k.DB.Save(kp)

	if k.DB.Error != nil {
		return k.DB.Error
	}

	return nil
}


