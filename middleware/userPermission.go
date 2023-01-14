package middleware

import (
	"golang-todo-app/entity"
	"golang-todo-app/model"

	"github.com/gofiber/fiber/v2"
)

func UserPemission(permission string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(entity.User)

		if user.HasPermission(permission) {
			return c.Next()
		}

		return c.
			Status(fiber.StatusUnauthorized).
			JSON(model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Dont have permission",
			})
	}
}
