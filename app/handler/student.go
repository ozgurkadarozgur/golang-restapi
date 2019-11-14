package handler

import (
	"net/http"
	"pro/restapi/app/model"
	"pro/restapi/app/resource"
	"pro/restapi/app/validator"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	//Db *gorm.DB
}

func (handler StudentHandler) Index(c *gin.Context) {

	var s model.Student

	students := s.All()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   resource.StudentResource{}.Collection(students),
	})

}

func (handler StudentHandler) Show(c *gin.Context, id string) {
	var s model.Student

	student := s.Show(id)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   resource.StudentResource{}.One(student),
	})

}

func (handler StudentHandler) Store(c *gin.Context, student model.Student) {

	v := validator.StudentValidator{
		Context: c,
	}

	err := v.ValidateStore()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	} else {
		s := student.Store()

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   s,
		})
	}

}
