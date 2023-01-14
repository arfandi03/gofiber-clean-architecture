package repository

import (
	"context"
	"golang-todo-app/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) entity.User
	Authentication(ctx context.Context, username string) entity.User
	FindById(ctx context.Context, username string) entity.User
}
