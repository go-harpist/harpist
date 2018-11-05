package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Permissions string
}
