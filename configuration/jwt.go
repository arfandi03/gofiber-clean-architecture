package configuration

import (
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(model model.UserModel) string {
	jwtSecret := os.Getenv("JWT_SECRET_TOKEN")
	jwtExpired, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"username":    model.Username,
		"roles":       model.Roles,
		"permissions": model.Permissions,
		"exp":         time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
