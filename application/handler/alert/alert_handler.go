package alert

import (
	"context"
	"hu-design-project/application/model/mongo_model"
)

type Handler struct {
	GetAlerts       GetAlertsHandler
	AddWhitelist    AddWhitelistHandler
	AddBlacklist    AddBlacklistHandler
	GetWhitelist    GetWhitelistHandler
	GetBlacklist    GetBlacklistHandler
	DeleteWhitelist DeleteWhitelistHandler
	DeleteBlacklist DeleteBlacklistHandler
}

type GetAlertsHandler interface {
	Handle(context context.Context) (*[]mongo_model.Alert, error)
}

type AddWhitelistHandler interface {
	Handle(context context.Context, id string) (*mongo_model.List, error)
}

type AddBlacklistHandler interface {
	Handle(context context.Context, id string) (*mongo_model.List, error)
}

type GetWhitelistHandler interface {
	Handle(context context.Context, id string) (*mongo_model.List, error)
}

type GetBlacklistHandler interface {
	Handle(context context.Context, id string) (*mongo_model.List, error)
}

type DeleteWhitelistHandler interface {
	Handle(ctx context.Context, id string) error
}

type DeleteBlacklistHandler interface {
	Handle(ctx context.Context, id string) error
}
