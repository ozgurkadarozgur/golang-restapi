package handler

import (
	"net/http"
	"pro/restapi/app/model"
	"pro/restapi/app/validator"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CourseHandler struct {
	Db *gorm.DB
}

func (handler CourseHandler) Index(c *gin.Context) {
	var course model.Course
	courses := course.All()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   courses,
	})
}

func (handler CourseHandler) Store(c *gin.Context, course model.Course) {

	v := validator.CourseValidator{
		Context: c,
	}

	err := v.ValidateStore()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	} else {
		course.Store()
	}

}
