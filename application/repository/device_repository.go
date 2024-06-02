package repository

import (
	"context"
	"hu-design-project/application/model/mongo_model"
)

type DevicesRepository interface {
	GetAllDevices(ctx context.Context) (*[]mongo_model.Device, error)
}
