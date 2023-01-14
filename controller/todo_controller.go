package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"

	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	service.TodoService
}

func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{TodoService: *todoService}
}

func (controller TodoController) Route(api fiber.Router) {
	todos := api.Group("/todos")
	todos.Post("/", controller.Create)
	// todos.Put("/:id", controller.Update)
	todos.Patch("/:id", controller.UpdateStatus)
	todos.Delete("/:id", controller.Delete)
	// todos.Get("/:id", controller.FindById)
	todos.Get("/", controller.FindAll)
}

func (controller TodoController) Create(c *fiber.Ctx) error {
	var request model.CreateTodo
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("user").(entity.User)

	response := controller.TodoService.Create(c.Context(), request, user.Username)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller TodoController) UpdateStatus(c *fiber.Ctx) error {
	var request model.UpdateTodoStatus
	id := c.Params("id")

	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("user").(entity.User)

	response := controller.TodoService.UpdateStatus(c.Context(), request, id, user.Username)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// func (controller TodoController) Update(c *fiber.Ctx) error {
// 	var request model.TodoCreateOrUpdateModel
// 	id, err := strconv.Atoi(c.Params("id"))
// 	err = c.BodyParser(&request)
// 	exception.PanicLogging(err)

// 	response := controller.TodoService.Update(c.Context(), request, id)
// 	return c.JSON(model.GeneralResponse{
// 		Code:    200,
// 		Message: "Success",
// 		Data:    response,
// 	})
// }

func (controller TodoController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	user := c.Locals("user").(entity.User)

	controller.TodoService.Delete(c.Context(), id, user.Username)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

// func (controller TodoController) FindById(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	exception.PanicLogging(err)

// 	result := controller.TodoService.FindById(c.Context(), id)
// 	return c.JSON(model.GeneralResponse{
// 		Code:    200,
// 		Message: "Success",
// 		Data:    result,
// 	})
// }

func (controller TodoController) FindAll(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.User)
	result := controller.TodoService.FindAll(c.Context(), user.Username)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
