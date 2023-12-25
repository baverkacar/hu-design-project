package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hu-design-project/application/repository"
	"time"
)

type userRepository struct {
	client *mongo.Client
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (repository *userRepository) ConnectToMongo(mongoURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repository *userRepository) Get() string {
	return "sagas"
}
