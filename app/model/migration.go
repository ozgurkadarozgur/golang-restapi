package model

import (
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.CreateTable(&Course{})
}
