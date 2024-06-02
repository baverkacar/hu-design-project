package request

import (
	"context"
	"hu-design-project/application/model/mongo_model"
)

type Handler struct {
	GetRequests GetRequestHandler
}

type GetRequestHandler interface {
	Handle(context context.Context) (*[]mongo_model.Request, error)
}
