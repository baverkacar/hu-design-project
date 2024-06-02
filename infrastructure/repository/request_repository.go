package repository

import (
	"context"
	"hu-design-project/application/model/mongo_model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RequestRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewRequestRepository(mongoURL, database, collection string) (*RequestRepository, error) {
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

	return &RequestRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *RequestRepository) GetAllRequests(ctx context.Context) (*[]mongo_model.Request, error) {
	var requests []mongo_model.Request

	coll := r.client.Database(r.database).Collection(r.collection)

	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Failed to retrieve requests: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var request mongo_model.Request
		if err := cursor.Decode(&request); err != nil {
			log.Printf("Failed to decode request: %v", err)
			return nil, err
		}
		requests = append(requests, request)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error during fetching requests: %v", err)
		return nil, err
	}

	return &requests, nil
}
