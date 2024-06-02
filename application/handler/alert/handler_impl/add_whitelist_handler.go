package handler_impl

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type AddWhitelistHandler struct {
	alertRepository     repository.AlertsRepository
	whitelistRepository repository.WhitelistRepository
}

func NewAddWhitelistHandler(alertsRepository repository.AlertsRepository, whitelistRepository repository.WhitelistRepository) *AddWhitelistHandler {
	return &AddWhitelistHandler{
		alertRepository:     alertsRepository,
		whitelistRepository: whitelistRepository,
	}
}

func (h *AddWhitelistHandler) Handle(ctx context.Context, alertId string) (*mongo_model.List, error) {
	objId, err := primitive.ObjectIDFromHex(alertId)
	if err != nil {
		return nil, errors.New("invalid alert ID")
	}

	alert, err := h.alertRepository.GetAlertById(ctx, objId)
	if err != nil {
		return nil, err
	}

	whitelistEntry := &mongo_model.List{
		AlertID:   alert.ID,
		IPAddress: alert.IPAddress,
		Date:      alert.CreatedAt,
	}

	whitelistEntry, err = h.whitelistRepository.AddWhitelist(ctx, whitelistEntry)
	if err != nil {
		return nil, err
	}

	return whitelistEntry, nil
}
