package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return "参数验证失败"
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		// 使用validator进行验证
		validate := validator.New()
		errs = ValidErrors{}
		if err := validate.Struct(v); err != nil {
			for _, fieldErr := range err.(validator.ValidationErrors) {
				errs = append(errs, &ValidError{
					Key:     fieldErr.Field(),
					Message: fieldErr.Error(),
				})
			}
		}
		return false, errs
	}

	return true, errs
}
