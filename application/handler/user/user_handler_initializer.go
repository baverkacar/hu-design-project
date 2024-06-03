package user

import (
	"hu-design-project/application/handler/user/handler_impl"
	"hu-design-project/application/repository"
)

func InitializeUserHandler(
	repository repository.UserRepository,
) *Handler {
	userHandler := Handler{}
	userHandler.CreateUser = handler_impl.NewCreateUserHandler(repository)
	userHandler.GetUser = handler_impl.NewGetUserByIDHandler(repository)
	userHandler.ActivateUser = handler_impl.NewActivateUserHandler(repository)
	userHandler.ChangePassword = handler_impl.NewChangePasswordHandler(repository)
	userHandler.ResetPassword = handler_impl.NewResetPasswordHandler(repository)
	return &userHandler
}
