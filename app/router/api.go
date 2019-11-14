package router

import (
	"net/http"
	"pro/restapi/app/handler"
	"pro/restapi/app/model"
	"pro/restapi/app/service"

	"github.com/gin-gonic/gin"
)

func Init() {

	studentHandler := handler.StudentHandler{}

	courseHandler := handler.CourseHandler{}

	r := service.GetRouter()

	r.GET("/students", func(c *gin.Context) {
		studentHandler.Index(c)
	})

	r.GET("/students/:id", func(c *gin.Context) {
		id := c.Param("id")
		studentHandler.Show(c, id)
	})

	r.POST("/students", func(c *gin.Context) {

		var s model.Student

		if err := c.ShouldBind(&s); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		studentHandler.Store(c, s)

	})

	r.GET("/courses", func(c *gin.Context) {
		courseHandler.Index(c)
	})

	r.POST("/courses", func(c *gin.Context) {
		var course model.Course

		if err := c.ShouldBind(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		courseHandler.Store(c, course)
	})

	r.Run(":8082")
}
