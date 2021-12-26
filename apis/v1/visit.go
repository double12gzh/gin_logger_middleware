package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	ctrlVisit "main/controller/Visit"
	"main/model"
)

func VisitorEmailValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(model.Visitor)

	if len(user.Email) == 0 {
		sl.ReportError(user.Email, "Email", "email", "email", "")
	}

}

var visitorNameValidation validator.Func = func(fl validator.FieldLevel) bool {
	name := fl.Field().Interface().(string)
	if len(name) == 0 {
		return false
	}

	return true
}

func SetupVisit(router *gin.RouterGroup) {
	auth := router.Group("/visit")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//https://github.com/gin-gonic/examples/tree/master/struct-lvl-validations
		v.RegisterStructValidation(VisitorEmailValidation, model.Visitor{})
		_ = v.RegisterValidation("visitorNameValidation", visitorNameValidation)
	}

	auth.GET("", ctrlVisit.Visit)
	auth.POST("/user", ctrlVisit.DO)
}
