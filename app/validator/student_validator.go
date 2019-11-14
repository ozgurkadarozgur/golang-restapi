package validator

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/gin-gonic/gin"
	"gopkg.in/thedevsaddam/govalidator.v1"
)

type StudentValidator struct {
	Context *gin.Context
}

func (validator StudentValidator) ValidateStore() map[string]interface{} {

	rules := govalidator.MapData{
		"first_name": []string{"required"},
		"last_name":  []string{"required"},
	}

	messages := govalidator.MapData{
		"first_name": []string{"required:Ad alanı boş bırakılamaz!"},
		"last_name":  []string{"required:Soyad alanı boş bırakılamaz!"},
	}

	opts := govalidator.Options{
		Request:  validator.Context.Request,
		Rules:    rules,
		Messages: messages,
	}

	v := govalidator.New(opts)

	e := v.Validate()

	fmt.Println("v_result", e)

	if !reflect.DeepEqual(e, url.Values{}) {
		return map[string]interface{}{"err": e}
	} else {
		return nil
	}

}
