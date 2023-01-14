package controller

import (
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{AuthService: *authService}
}

func (controller AuthController) Route(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
}

func (controller AuthController) Register(c *fiber.Ctx) error {
	var request model.RegisterUser
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.AuthService.CreateUser(c.Context(), request)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}
func (controller AuthController) Login(c *fiber.Ctx) error {
	var request model.LoginUser
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.AuthService.Login(c.Context(), request)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}
