package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/enum"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"github.com/google/uuid"
)

type todoServiceImpl struct {
	repository.TodoRepository
}

func NewTodoServiceImpl(todoRepository *repository.TodoRepository) TodoService {
	return &todoServiceImpl{TodoRepository: *todoRepository}
}

func (service *todoServiceImpl) Create(ctx context.Context, todoModel model.CreateTodo, username string) model.TodoModel {
	validation.Validate(todoModel)

	todo := entity.Todo{
		Id:          uuid.New(),
		Title:       todoModel.Title,
		Description: todoModel.Description,
		Status:      enum.OPEN,
		Username:    username,
	}
	todo = service.TodoRepository.Insert(ctx, todo)

	return model.TodoModel{
		Id:          todo.Id.String(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
		Username:    todo.Username,
	}
}
func (service *todoServiceImpl) UpdateStatus(ctx context.Context, todoModel model.UpdateTodoStatus, id string, username string) model.TodoModel {
	validation.Validate(todoModel)

	todo := service.TodoRepository.FindById(ctx, id)
	if todo.Username != username {
		panic(exception.NotFoundError{
			Message: "Todo Not Found",
		})
	}

	todo.Status = enum.TodoStatus(todoModel.Status)

	service.TodoRepository.Update(ctx, todo)
	return model.TodoModel{
		Id:          todo.Id.String(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
		Username:    todo.Username,
	}
}

// func (service *todoServiceImpl) Update(ctx context.Context, todoModel model.TodoCreateOrUpdateModel, id int) model.TodoCreateOrUpdateModel {
// 	todo := entity.Todo{
// 		Id:       uint8(id),
// 		Name:     todoModel.Name,
// 		Price:    todoModel.Price,
// 		Quantity: todoModel.Quantity,
// 	}
// 	service.TodoRepository.Update(ctx, todo)
// 	return todoModel
// }

func (service *todoServiceImpl) Delete(ctx context.Context, id string, username string) {
	todo := service.TodoRepository.FindById(ctx, id)
	if todo.Username != username {
		panic(exception.NotFoundError{
			Message: "Todo Not Found",
		})
	}
	service.TodoRepository.Delete(ctx, todo)
}

// func (service *todoServiceImpl) FindById(ctx context.Context, id int) model.TodoModel {
// 	todo, err := service.TodoRepository.FindById(ctx, id)
// 	exception.PanicLogging(err)
// 	return model.TodoModel{
// 		Id:       todo.Id,
// 		Name:     todo.Name,
// 		Price:    todo.Price,
// 		Quantity: todo.Quantity,
// 	}
// }

func (service *todoServiceImpl) FindAll(ctx context.Context, username string) (responses []model.TodoModel) {
	todos := service.TodoRepository.FindAll(ctx, username)

	if len(todos) == 0 {
		return []model.TodoModel{}
	}

	for _, todo := range todos {
		responses = append(responses, model.TodoModel{
			Id:          todo.Id.String(),
			Title:       todo.Title,
			Description: todo.Description,
			Status:      string(todo.Status),
		})
	}
	return responses
}
