package service

import (
	"context"
	"golang-todo-app/configuration"
	"golang-todo-app/entity"
	"golang-todo-app/enum"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	repository.UserRepository
	repository.RoleRepository
}

func NewAuthServiceImpl(userRepository *repository.UserRepository, roleRepository *repository.RoleRepository) AuthService {
	return &authServiceImpl{UserRepository: *userRepository, RoleRepository: *roleRepository}
}

func (service authServiceImpl) CreateUser(ctx context.Context, authModel model.RegisterUser) model.UserModel {
	validation.Validate(authModel)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authModel.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	role := service.RoleRepository.FindByRole(ctx, string(enum.USER))
	permissons := role.Permissions
	user := entity.User{
		Username: authModel.Username,
		Password: string(hashedPassword),
		IsActive: true,
		Roles: []entity.Role{
			role,
		},
		Permissions: permissons,
	}

	user = service.UserRepository.Create(ctx, user)

	return model.UserModel{
		Username:    user.Username,
		Roles:       user.GetRoles(),
		Permissions: user.GetPermissions(),
	}
}

func (service authServiceImpl) Login(ctx context.Context, authModel model.LoginUser) model.AuthToken {
	validation.Validate(authModel)

	user := service.UserRepository.Authentication(ctx, authModel.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authModel.Password)); err != nil {
		panic(exception.UnauthorizedError{
			Message: "incorrect username and password",
		})
	}

	jwtBody := model.UserModel{
		Username:    user.Username,
		Roles:       user.GetRoles(),
		Permissions: user.GetPermissions(),
	}

	tokenJwtResult := configuration.GenerateToken(jwtBody)

	return model.AuthToken{
		Token:     tokenJwtResult,
		UserModel: jwtBody,
	}
}

/*
ref:
- https://stackoverflow.com/questions/23259586/bcrypt-password-hashing-in-golang-compatible-with-node-js
*/
