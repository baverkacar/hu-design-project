package device

import (
	"context"
	"hu-design-project/application/model/mongo_model"
)

type Handler struct {
	GetDevices GetDevicesHandler
}

type GetDevicesHandler interface {
	Handle(context context.Context) (*[]mongo_model.Device, error)
}
