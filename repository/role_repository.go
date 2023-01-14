package repository

import (
	"context"
	"golang-todo-app/entity"
)

type RoleRepository interface {
	FindByRole(ctx context.Context, role string) entity.Role
}
