package validator

import (
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
	"gopkg.in/thedevsaddam/govalidator.v1"
)

type CourseValidator struct {
	Context *gin.Context
}

func (validator CourseValidator) ValidateStore() map[string]interface{} {

	rules := govalidator.MapData{
		"title":     []string{"required"},
		"studentId": []string{"required"},
	}

	messages := govalidator.MapData{
		"title":     []string{"required:Kurs adı alanı boş bırakılamaz!"},
		"studentId": []string{"required:Öğrenci alanı boş bırakılamaz!"},
	}

	opts := govalidator.Options{
		Request:  validator.Context.Request,
		Rules:    rules,
		Messages: messages,
	}

	v := govalidator.New(opts)

	e := v.Validate()

	if !reflect.DeepEqual(e, url.Values{}) {
		return map[string]interface{}{"err": e}
	} else {
		return nil
	}

}
