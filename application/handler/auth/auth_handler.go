package auth

import (
	"context"
	"hu-design-project/application/model"
	"hu-design-project/application/model/mongo_model"
)

type Handler struct {
	UserLogin UserLoginHandler
}

type UserLoginHandler interface {
	Handle(context context.Context, requestBody model.UserLoginModel) (*mongo_model.User, error)
}
