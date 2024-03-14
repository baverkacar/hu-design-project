package handler_impl

import (
	"context"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type GetUserByIDHandler struct {
	repository repository.UserRepository
}

func NewGetUserByIDHandler(
	repository repository.UserRepository) *GetUserByIDHandler {
	return &GetUserByIDHandler{
		repository: repository,
	}
}

func (handler *GetUserByIDHandler) Handle(ctx context.Context, id string) (*mongo_model.User, error) {
	user, err := handler.repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
