package handler_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type GetWhitelistHandler struct {
	repository repository.WhitelistRepository
}

func NewGetWhitelistHandler(repository repository.WhitelistRepository) *GetWhitelistHandler {
	return &GetWhitelistHandler{
		repository: repository,
	}
}

func (h *GetWhitelistHandler) Handle(ctx context.Context, id string) (*mongo_model.List, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	whitelistEntry, err := h.repository.GetWhitelistById(ctx, objId)
	if err != nil {
		return nil, err
	}
	return whitelistEntry, nil
}
