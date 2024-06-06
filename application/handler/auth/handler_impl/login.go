package handler_impl

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"hu-design-project/application/model"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type LoginHandler struct {
	repository repository.UserRepository
}

func NewUserLoginHandler(
	repository repository.UserRepository) *LoginHandler {
	return &LoginHandler{
		repository: repository,
	}
}

func (h *LoginHandler) Handle(ctx context.Context, requestBody model.UserLoginModel) (*mongo_model.User, error) {
	log.Info("[AuthHandler] Login user", requestBody.Email)
	user, err := h.repository.GetUserByEmail(ctx, requestBody.Email)
	if err != nil {
		return nil, err
	}
	ispPasswordsEquals := user.Password == requestBody.Password
	if !ispPasswordsEquals {
		return nil, errors.New("invalid password")
	}

	if !user.IsActive {
		return nil, errors.New("account is not activated. Please activate your account")
	}

	return user, nil
}
