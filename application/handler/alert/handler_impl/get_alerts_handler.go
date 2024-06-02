package handler_impl

import (
	"context"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type GetAlertsHandler struct {
	repository repository.AlertsRepository
}

func NewGetAlertsHandler(repository repository.AlertsRepository) *GetAlertsHandler {
	return &GetAlertsHandler{
		repository: repository,
	}
}

func (h *GetAlertsHandler) Handle(ctx context.Context) (*[]mongo_model.Alert, error) {
	alerts, err := h.repository.GetAllAlerts(ctx)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}
