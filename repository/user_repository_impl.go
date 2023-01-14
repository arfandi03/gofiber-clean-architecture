package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: DB}
}

func (repository userRepositoryImpl) Create(ctx context.Context, user entity.User) entity.User {
	err := repository.DB.WithContext(ctx).Omit("Todos").Create(&user).Error
	exception.PanicLogging(err)
	return user
}

func (userRepository userRepositoryImpl) Authentication(ctx context.Context, username string) entity.User {
	var user entity.User
	result := userRepository.DB.WithContext(ctx).
		Preload("Roles").
		Preload("Permissions").
		Where("tb_user.username = ? and tb_user.is_active = ?", username, true).
		Find(&user)
	if result.RowsAffected == 0 {
		panic(exception.UnauthorizedError{
			Message: "Incorrect username and password",
		})
	}
	return user
}

func (repository userRepositoryImpl) FindById(ctx context.Context, username string) entity.User {
	var user entity.User
	result := repository.DB.WithContext(ctx).Preload("Roles").Preload("Permissions").Where("username = ?", username).First(&user)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "User Not Found",
		})
	}
	return user
}
