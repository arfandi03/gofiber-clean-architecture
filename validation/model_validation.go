package validation

import (
	"encoding/json"
	"golang-todo-app/exception"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Validate(modelValidate interface{}) {
	validate := validator.New()
	// add custom validation
	validate.RegisterValidation("regexp", Regexp)
	err := validate.Struct(modelValidate)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": msgForFieldTag(err),
			})
		}

		jsonMessage, errJson := json.Marshal(messages)
		exception.PanicLogging(errJson)

		panic(exception.ValidationError{
			Message: string(jsonMessage),
		})
	}
}

func Regexp(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(fl.Param())
	return re.MatchString(fl.Field().String())
}

func msgForFieldTag(err validator.FieldError) string {
	if err.Field() == "Password" && err.Tag() == "regexp" {
		return "This field must begin with char"
	}
	return "this field is " + err.Tag() + " " + err.Param()
}

/*
ref:
- https://medium.com/tunaiku-tech/go-validator-v10-c7a4f1be37df
*/
