package service

import (
	"pro/restapi/app/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Db     *gorm.DB
	Router *gin.Engine
}

var service Service

func GetDB() *gorm.DB {
	if service.Db == nil {
		db, err := database.Init()
		if err != nil {
			panic(err)
		}
		service.Db = db
	}
	return service.Db
}

func GetRouter() *gin.Engine {
	if service.Router == nil {
		service.Router = gin.Default()
	}
	return service.Router
}
