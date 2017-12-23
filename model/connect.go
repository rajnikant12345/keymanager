package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

type Model interface {
	Connect() (*gorm.DB, error)
}

