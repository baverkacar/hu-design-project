package handler_impl

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type AddBlacklistHandler struct {
	alertRepository     repository.AlertsRepository
	blacklistRepository repository.BlacklistRepository
}

func NewAddBlacklistHandler(alertsRepository repository.AlertsRepository, blacklistRepository repository.BlacklistRepository) *AddBlacklistHandler {
	return &AddBlacklistHandler{
		alertRepository:     alertsRepository,
		blacklistRepository: blacklistRepository,
	}
}

func (h *AddBlacklistHandler) Handle(ctx context.Context, alertId string) (*mongo_model.List, error) {
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

	whitelistEntry, err = h.blacklistRepository.AddBlacklist(ctx, whitelistEntry)
	if err != nil {
		return nil, err
	}

	return whitelistEntry, nil
}
