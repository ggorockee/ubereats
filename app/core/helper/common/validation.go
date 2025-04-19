package common

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(input any) error {

	validate := validator.New()

	// validate.RegisterValidation("follow_status", func(fl validator.FieldLevel) bool {
	// 	// fl.Field().Interface().(string)
	// 	return fl.Field().String() == "pending" ||
	// 		fl.Field().String() == "approved" ||
	// 		fl.Field().String() == "blocked"
	// })

	validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "client" ||
			fl.Field().String() == "owner" ||
			fl.Field().String() == "delivery" ||
			fl.Field().String() == "any" ||
			fl.Field().String() == "admin" ||
			fl.Field().String() == ""

	})

	err := validate.Struct(input)
	if err != nil {
		return err // 컨텍스트가 없으면 오류만 반환
	}
	return nil

}
