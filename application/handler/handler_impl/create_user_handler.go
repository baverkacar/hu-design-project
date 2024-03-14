package handler_impl

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"hu-design-project/application/model"
	"hu-design-project/application/model/dto"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
	"hu-design-project/infrastructure/util"
)

type CreateUserHandler struct {
	repository repository.UserRepository
}

func NewCreateUserHandler(
	repository repository.UserRepository,
) *CreateUserHandler {
	return &CreateUserHandler{
		repository: repository,
	}
}

func (handler *CreateUserHandler) Handle(ctx context.Context, userCreateModel *model.UserCreateModel) (*dto.UserDTO, error) {
	log.Info("[CreateUserHandler] creating user")
	exists, err := handler.repository.CheckEmail(ctx, userCreateModel.Email)
	if err != nil {
		log.Info("[CreateUserHandler] error occurred on email check")
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("email already in use: %s", userCreateModel.Email)
	}

	user := mongo_model.NewUser(*userCreateModel)
	user, err = handler.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	code, err := util.SendMail(user.Email)
	if err != nil {
		log.Info("[CreateUserHandler] error occurred on sending email")
		return nil, err
	}

	return user.ToDTO(code), nil
}
