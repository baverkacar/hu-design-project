package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/model/mongo_model"
)

type AlertsRepository interface {
	GetAllAlerts(ctx context.Context) (*[]mongo_model.Alert, error)
	GetAlertById(ctx context.Context, id primitive.ObjectID) (*mongo_model.Alert, error)
}
