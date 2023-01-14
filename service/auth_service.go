package service

import (
	"context"
	"golang-todo-app/model"
)

type AuthService interface {
	CreateUser(ctx context.Context, userModel model.RegisterUser) model.UserModel
	Login(ctx context.Context, userModel model.LoginUser) model.AuthToken
}
