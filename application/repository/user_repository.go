package repository

import (
	"context"
	"hu-design-project/application/model/mongo_model"
)

type UserRepository interface {
	CheckEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *mongo_model.User) (*mongo_model.User, error)
	GetUserById(ctx context.Context, id string) (*mongo_model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*mongo_model.User, error)
	ActivateUser(ctx context.Context, user *mongo_model.User) error
	UpdatePasswordByEmail(ctx context.Context, email, newPassword string) error
	ChangePassword(ctx context.Context, userID string, oldPassword, newPassword string) error
}
