package handler

import (
	"context"
	"hu-design-project/application/model"
	"hu-design-project/application/model/dto"
	"hu-design-project/application/model/mongo_model"
)

type UserHandler struct {
	CreateUser   CreateUserHandler
	GetUser      GetUserByIdHandler
	ActivateUser ActivateUserHandler
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
