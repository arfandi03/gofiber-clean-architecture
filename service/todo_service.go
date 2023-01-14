package service

import (
	"context"
	"golang-todo-app/model"
)

type TodoService interface {
	Create(ctx context.Context, todoModel model.CreateTodo, username string) model.TodoModel
	UpdateStatus(ctx context.Context, todoModel model.UpdateTodoStatus, id string, username string) model.TodoModel
	// Update(ctx context.Context, productModel model.TodoCreateOrUpdateModel, id int) model.TodoCreateOrUpdateModel
	Delete(ctx context.Context, id string, username string)
	// FindById(ctx context.Context, id int) model.TodoModel
	FindAll(ctx context.Context, username string) []model.TodoModel
}
