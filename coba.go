package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/joho/godotenv/autoload"
)

type AuthToken struct {
	Token   string
	JwtBody JwtBody
}

type JwtBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=40,regexp=^.*[a-zA-Z].*$"`
	// Password string `json:"password" validate:"min=3,max=40,regexp=^(?:(?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[A-Z])(?=.*[a-z]).*$"`
}

func msgForFieldTag(err validator.FieldError) string {
	fmt.Println("err", err)
	if err.Field() == "Password" && err.Tag() == "regexp" {
		return "This field must begin with char"
	}
	return "this field is " + err.Tag() + " " + err.Param()
}

func Regexp(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(fl.Param())
	return re.MatchString(fl.Field().String())
}

func coba() {
	u := AuthToken{}
	// u.Username = "admin"
	// u.Password = "s"
	u.JwtBody = JwtBody{
		Username: "admin",
		Password: "s",
	}

	valid := validator.New()
	valid.RegisterValidation("regexp", Regexp)
	// valid.RegisterCustomTypeFunc(Regexp, String{})
	err := valid.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("field", err.Field())
			fmt.Println("message", msgForFieldTag(err))
		}
		// os.Exit(1)
	}

	fmt.Println("ok")
	fmt.Println(os.Getenv("SERVER_PORT"))
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("usemahu001"), bcrypt.DefaultCost)

	fmt.Println("password", string(hashedPassword))
	json, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("json", string(json))
}
