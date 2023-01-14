package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type todoRepositoryImpl struct {
	*gorm.DB
}

func NewTodoRepositoryImpl(DB *gorm.DB) TodoRepository {
	return &todoRepositoryImpl{DB: DB}
}

func (repository *todoRepositoryImpl) Insert(ctx context.Context, todo entity.Todo) entity.Todo {
	err := repository.DB.WithContext(ctx).Create(&todo).Error
	exception.PanicLogging(err)
	return todo
}

func (repository *todoRepositoryImpl) Update(ctx context.Context, todo entity.Todo) entity.Todo {
	err := repository.DB.WithContext(ctx).Where("id = ?", todo.Id).Updates(&todo).Error
	exception.PanicLogging(err)
	return todo
}

func (repository *todoRepositoryImpl) Delete(ctx context.Context, todo entity.Todo) {
	repository.DB.WithContext(ctx).Where("id = ?", todo.Id).Delete(&todo)
}

func (repository *todoRepositoryImpl) FindById(ctx context.Context, id string) entity.Todo {
	var todo entity.Todo
	result := repository.DB.WithContext(ctx).Where("id = ?", id).First(&todo)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Todo Not Found",
		})
	}
	return todo
}

func (repository *todoRepositoryImpl) FindAll(ctx context.Context, username string) []entity.Todo {
	var todos []entity.Todo
	repository.DB.WithContext(ctx).Where("username = ?", username).Find(&todos)
	return todos
}
