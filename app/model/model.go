package model

import (
	"pro/restapi/app/service"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init() {
	db = service.GetDB()
}
