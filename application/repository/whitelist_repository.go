package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hu-design-project/application/model/mongo_model"
)

type WhitelistRepository interface {
	AddWhitelist(ctx context.Context, whitelist *mongo_model.List) (*mongo_model.List, error)
	GetWhitelistById(ctx context.Context, id primitive.ObjectID) (*mongo_model.List, error)
	DeleteWhitelistById(ctx context.Context, id primitive.ObjectID) error
	GetAllWhitelists(ctx context.Context) (*[]mongo_model.List, error)
}
