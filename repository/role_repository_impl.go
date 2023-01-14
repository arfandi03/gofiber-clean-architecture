package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	*gorm.DB
}

func NewRoleRepositoryImpl(DB *gorm.DB) RoleRepository {
	return &roleRepositoryImpl{DB: DB}
}

func (repository roleRepositoryImpl) FindByRole(ctx context.Context, rolename string) entity.Role {
	var role entity.Role
	result := repository.DB.WithContext(ctx).Preload("Permissions").Where("name = ?", rolename).First(&role)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Role Not Found",
		})
	}
	return role
}
