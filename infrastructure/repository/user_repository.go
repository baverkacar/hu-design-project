package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hu-design-project/application/model/mongo_model"
	"time"
)

type UserRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewUserRepository(mongoURL, database, collection string) (*UserRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (repo *UserRepository) CheckEmail(ctx context.Context, email string) (bool, error) {
	collection := repo.client.Database(repo.database).Collection(repo.collection)
	var result interface{}
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *mongo_model.User) (*mongo_model.User, error) {
	collection := repo.client.Database(repo.database).Collection(repo.collection)
	inserted, err := collection.InsertOne(ctx, user)

	objID, ok := inserted.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Info("[UserRepository] error occurred on inserting user")
		return nil, err
	}

	newUser, err := repo.GetUserById(ctx, objID.Hex())
	if err != nil {
		log.Info("[UserRepository] error occurred on getting user with id")
		return nil, err
	}
	return newUser, nil
}

func (repo *UserRepository) GetUserById(ctx context.Context, id string) (*mongo_model.User, error) {
	var user mongo_model.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Info("[UserRepository] converting hex to object id")
		return nil, err
	}

	coll := repo.client.Database(repo.database).Collection(repo.collection)
	err = coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		log.Info("[UserRepository] error occurred on getting user with id")
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*mongo_model.User, error) {
	var user mongo_model.User
	coll := repo.client.Database(repo.database).Collection(repo.collection)

	err := coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Info("[UserRepository] No user found with email:", email)
			return nil, fmt.Errorf("user not found")
		} else {
			log.Error("[UserRepository] Error occurred on getting user with email:", email)
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) ActivateUser(ctx context.Context, user *mongo_model.User) error {
	coll := repo.client.Database(repo.database).Collection(repo.collection)

	update := bson.M{
		"$set": bson.M{
			"isActive":   true,
			"lastUpdate": time.Now(),
		},
	}
	log.Info("[UserController] Activating user")
	_, err := coll.UpdateOne(ctx, bson.M{"_id": user.UserID}, update)
	if err != nil {
		return err
	}

	return nil
}
