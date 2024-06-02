package handler_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/repository"
)

type DeleteWhitelistHandler struct {
	repository repository.WhitelistRepository
}

func NewDeleteWhitelistHandler(repository repository.WhitelistRepository) *DeleteWhitelistHandler {
	return &DeleteWhitelistHandler{
		repository: repository,
	}
}

func (h *DeleteWhitelistHandler) Handle(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err // ID format hatası
	}

	err = h.repository.DeleteWhitelistById(ctx, objId)
	if err != nil {
		return err // Silme işlemi sırasında hata
	}
	return nil
}
