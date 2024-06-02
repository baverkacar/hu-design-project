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

type DevicesRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewDevicesRepository(mongoURL, database, collection string) (*DevicesRepository, error) {
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

	return &DevicesRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *DevicesRepository) GetAllDevices(ctx context.Context) (*[]mongo_model.Device, error) {
	var devices []mongo_model.Device

	coll := r.client.Database(r.database).Collection(r.collection)

	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Failed to retrieve devices: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var device mongo_model.Device
		if err := cursor.Decode(&device); err != nil {
			log.Printf("Failed to decode device: %v", err)
			return nil, err
		}
		devices = append(devices, device)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error during fetching devices: %v", err)
		return nil, err
	}

	return &devices, nil
}
