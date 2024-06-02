package handler_impl

import (
	"context"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type RequestHandler struct {
	repository repository.RequestRepository
}

// NewRequestHandler, RequestHandler için bir constructor fonksiyonu oluşturur.
func NewGetRequestHandler(repository repository.RequestRepository) *RequestHandler {
	return &RequestHandler{
		repository: repository,
	}
}

// Handle, repository'den tüm requestleri alır.
func (h *RequestHandler) Handle(ctx context.Context) (*[]mongo_model.Request, error) {
	requests, err := h.repository.GetAllRequests(ctx)
	if err != nil {
		return nil, err
	}
	return requests, nil
}
