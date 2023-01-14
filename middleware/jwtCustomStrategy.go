package middleware

import (
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JwtCustomStrategy(userRepository repository.UserRepository) func(*fiber.Ctx) error {
	jwtSecret := os.Getenv("JWT_SECRET_TOKEN")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)

			user := userRepository.FindById(c.Context(), claims["username"].(string))

			c.Locals("user", user)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				// panic(exception.ValidationError{
				// 	Message: "Missing or malformed JWT",
				// })
				return c.
					Status(fiber.StatusBadRequest).
					JSON(model.GeneralResponse{
						Code:    400,
						Message: "Bad Request",
						Data:    "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(model.GeneralResponse{
						Code:    401,
						Message: "Unauthorized",
						Data:    "Invalid or expired JWT",
					})
			}
		},
	})
}
