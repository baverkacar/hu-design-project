package auth

import (
	"hu-design-project/application/handler/auth/handler_impl"
	"hu-design-project/application/repository"
)

func InitializeAuthHandler(
	repository repository.UserRepository,
) *Handler {
	authHandler := Handler{}
	authHandler.UserLogin = handler_impl.NewUserLoginHandler(repository)
	return &authHandler
}
