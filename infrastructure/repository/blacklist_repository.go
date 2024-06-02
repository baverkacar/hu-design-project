package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hu-design-project/application/model/mongo_model"
	"log"
	"time"
)

type BlacklistRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewBlacklistRepository(mongoURL, database, collection string) (*BlacklistRepository, error) {
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

	return &BlacklistRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *BlacklistRepository) AddBlacklist(ctx context.Context, blacklist *mongo_model.List) (*mongo_model.List, error) {
	coll := r.client.Database(r.database).Collection(r.collection)
	blacklist.ID = primitive.NewObjectID() // Yeni bir ObjectID oluştur

	result, err := coll.InsertOne(ctx, blacklist)
	if err != nil {
		log.Printf("Failed to insert blacklist entry: %v", err)
		return nil, err
	}

	blacklist.ID = result.InsertedID.(primitive.ObjectID)
	return blacklist, nil
}

func (r *BlacklistRepository) GetBlacklistById(ctx context.Context, id primitive.ObjectID) (*mongo_model.List, error) {
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

func (r *BlacklistRepository) DeleteBlacklistById(ctx context.Context, id primitive.ObjectID) error {
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

func (r *BlacklistRepository) GetAllBlacklists(ctx context.Context) (*[]mongo_model.List, error) {
	var blacklists []mongo_model.List
	coll := r.client.Database(r.database).Collection(r.collection)

	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var blacklist mongo_model.List
		if err := cursor.Decode(&blacklist); err != nil {
			return nil, err
		}
		blacklists = append(blacklists, blacklist)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &blacklists, nil
}
