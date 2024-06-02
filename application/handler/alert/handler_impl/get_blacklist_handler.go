package handler_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type GetBlacklistHandler struct {
	repository repository.BlacklistRepository
}

func NewGetBlacklistHandler(repository repository.BlacklistRepository) *GetBlacklistHandler {
	return &GetBlacklistHandler{
		repository: repository,
	}
}

func (h *GetBlacklistHandler) Handle(ctx context.Context, id string) (*mongo_model.List, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	blacklistEntry, err := h.repository.GetBlacklistById(ctx, objId)
	if err != nil {
		return nil, err
	}
	return blacklistEntry, nil
}
