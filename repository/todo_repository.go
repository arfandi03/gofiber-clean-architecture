package repository

import (
	"context"
	"golang-todo-app/entity"
)

type TodoRepository interface {
	Insert(ctx context.Context, todo entity.Todo) entity.Todo
	Update(ctx context.Context, todo entity.Todo) entity.Todo
	Delete(ctx context.Context, todo entity.Todo)
	FindById(ctx context.Context, id string) entity.Todo
	FindAll(ctx context.Context, username string) []entity.Todo
}
