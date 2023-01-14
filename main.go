package main

import (
	"golang-todo-app/configuration"
	"golang-todo-app/controller"
	"golang-todo-app/exception"
	"golang-todo-app/middleware"
	"golang-todo-app/repository"
	"golang-todo-app/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database := configuration.NewDatabase()

	//repository
	todoRepository := repository.NewTodoRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)
	roleRepository := repository.NewRoleRepositoryImpl(database)

	//service
	todoService := service.NewTodoServiceImpl(&todoRepository)
	authService := service.NewAuthServiceImpl(&userRepository, &roleRepository)

	//controller
	todoController := controller.NewTodoController(&todoService)
	authController := controller.NewAuthController(&authService)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfig())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New(configuration.NewLoggerConfig()))

	//routing
	api := app.Group("/api/v1")
	authController.Route(api)
	api.Use(middleware.JwtCustomStrategy(userRepository))
	todoController.Route(api)

	err := app.Listen(os.Getenv("SERVER_PORT"))
	exception.PanicLogging(err)

	configuration.NewLogger().Info("log info")
}

/*
ref:
- https://github.com/RizkiMufrizal/gofiber-clean-architecture
- https://project-awesome.org/gofiber/awesome-fiber
*/
