package repository

import (
	"context"
	"hu-design-project/application/model/mongo_model"
)

type RequestRepository interface {
	GetAllRequests(ctx context.Context) (*[]mongo_model.Request, error)
}
