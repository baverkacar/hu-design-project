package handler

import (
	"hu-design-project/application/handler/handler_impl"
	"hu-design-project/application/repository"
)

func InitializeUserHandler(
	repository repository.UserRepository,
) *UserHandler {
	userHandler := UserHandler{}
	userHandler.CreateUser = handler_impl.NewCreateUserHandler(repository)
	userHandler.GetUser = handler_impl.NewGetUserByIDHandler(repository)
	userHandler.ActivateUser = handler_impl.NewActivateUserHandler(repository)
	return &userHandler
}
