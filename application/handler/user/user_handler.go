package user

import (
	"context"
	"hu-design-project/application/model"
	"hu-design-project/application/model/dto"
	"hu-design-project/application/model/mongo_model"
)

type Handler struct {
	CreateUser     CreateUserHandler
	GetUser        GetUserByIdHandler
	ActivateUser   ActivateUserHandler
	ChangePassword ChangePasswordHandler
	ResetPassword  ResetPasswordHandler
}
type CreateUserHandler interface {
	Handle(context context.Context, song *model.UserCreateModel) (*dto.UserDTO, error)
}

type GetUserByIdHandler interface {
	Handle(ctx context.Context, id string) (*mongo_model.User, error)
}

type ActivateUserHandler interface {
	Handle(ctx context.Context, email string) error
}

type ChangePasswordHandler interface {
	Handle(ctx context.Context, passwordModel model.ChangePasswordModel) error
}

type ResetPasswordHandler interface {
	Handle(ctx context.Context, email string) error
}
