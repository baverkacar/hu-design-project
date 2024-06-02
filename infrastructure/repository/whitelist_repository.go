package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hu-design-project/application/model/mongo_model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WhitelistRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewWhitelistRepository(mongoURL, database, collection string) (*WhitelistRepository, error) {
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

	return &WhitelistRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *WhitelistRepository) AddWhitelist(ctx context.Context, whitelist *mongo_model.List) (*mongo_model.List, error) {
	coll := r.client.Database(r.database).Collection(r.collection)
	whitelist.ID = primitive.NewObjectID() // Yeni bir ObjectID oluştur

	result, err := coll.InsertOne(ctx, whitelist)
	if err != nil {
		log.Printf("Failed to insert whitelist entry: %v", err)
		return nil, err
	}

	whitelist.ID = result.InsertedID.(primitive.ObjectID)
	return whitelist, nil
}

func (r *WhitelistRepository) GetWhitelistById(ctx context.Context, id primitive.ObjectID) (*mongo_model.List, error) {
	var whitelist mongo_model.List

	coll := r.client.Database(r.database).Collection(r.collection)

	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&whitelist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No whitelist entry found with ID: %v", id)
			return nil, err
		}
		log.Printf("Failed to retrieve whitelist entry with ID: %v, error: %v", id, err)
		return nil, err
	}

	return &whitelist, nil
}

func (r *WhitelistRepository) DeleteWhitelistById(ctx context.Context, id primitive.ObjectID) error {
	coll := r.client.Database(r.database).Collection(r.collection)
	result, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *WhitelistRepository) GetAllWhitelists(ctx context.Context) (*[]mongo_model.List, error) {
	var whitelists []mongo_model.List
	coll := r.client.Database(r.database).Collection(r.collection)

	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var whitelist mongo_model.List
		if err := cursor.Decode(&whitelist); err != nil {
			return nil, err
		}
		whitelists = append(whitelists, whitelist)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &whitelists, nil
}
