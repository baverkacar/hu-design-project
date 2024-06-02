package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/model/mongo_model"
)

type BlacklistRepository interface {
	AddBlacklist(ctx context.Context, whitelist *mongo_model.List) (*mongo_model.List, error)
	GetBlacklistById(ctx context.Context, id primitive.ObjectID) (*mongo_model.List, error)
	DeleteBlacklistById(ctx context.Context, id primitive.ObjectID) error
	GetAllBlacklists(ctx context.Context) (*[]mongo_model.List, error)
}
