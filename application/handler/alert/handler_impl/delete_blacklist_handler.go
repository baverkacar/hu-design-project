package handler_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/repository"
)

type DeleteBlacklistHandler struct {
	repository repository.BlacklistRepository
}

func NewDeleteBlacklistHandler(repository repository.BlacklistRepository) *DeleteBlacklistHandler {
	return &DeleteBlacklistHandler{
		repository: repository,
	}
}

func (h *DeleteBlacklistHandler) Handle(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err // ID format hatası
	}

	err = h.repository.DeleteBlacklistById(ctx, objId)
	if err != nil {
		return err // Silme işlemi sırasında hata
	}
	return nil
}
