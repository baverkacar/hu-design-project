package handler_impl

import (
	"context"
	"github.com/labstack/gommon/log"
	"hu-design-project/application/repository"
)

type ActivateUserHandler struct {
	repository repository.UserRepository
}

func NewActivateUserHandler(
	repository repository.UserRepository) *ActivateUserHandler {
	return &ActivateUserHandler{
		repository: repository,
	}
}

func (handler *ActivateUserHandler) Handle(ctx context.Context, email string) error {
	user, err := handler.repository.GetUserByEmail(ctx, email)
	log.Info("[UserHandler] Activating user with email:", email)
	if err != nil {
		return err
	}
	err = handler.repository.ActivateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
