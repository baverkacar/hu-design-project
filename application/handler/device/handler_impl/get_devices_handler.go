package handler_impl

import (
	"context"
	"hu-design-project/application/model/mongo_model"
	"hu-design-project/application/repository"
)

type DevicesHandler struct {
	repository repository.DevicesRepository
}

func NewDevicesHandler(repository repository.DevicesRepository) *DevicesHandler {
	return &DevicesHandler{
		repository: repository,
	}
}

func (h *DevicesHandler) Handle(ctx context.Context) (*[]mongo_model.Device, error) {
	devices, err := h.repository.GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
