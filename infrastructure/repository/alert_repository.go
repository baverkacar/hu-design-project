package repository

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hu-design-project/application/model/mongo_model"
	"time"
)

type AlertsRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewAlertsRepository(mongoURL, database, collection string) (*AlertsRepository, error) {
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

	return &AlertsRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *AlertsRepository) GetAllAlerts(ctx context.Context) (*[]mongo_model.Alert, error) {
	var alerts []mongo_model.Alert

	coll := r.client.Database(r.database).Collection(r.collection)

	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Printf("Failed to retrieve alerts: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var alert mongo_model.Alert
		if err := cursor.Decode(&alert); err != nil {
			log.Printf("Failed to decode alert: %v", err)
			return nil, err
		}
		alerts = append(alerts, alert)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error during fetching alerts: %v", err)
		return nil, err
	}

	return &alerts, nil
}

func (r *AlertsRepository) GetAlertById(ctx context.Context, id primitive.ObjectID) (*mongo_model.Alert, error) {
	var alert mongo_model.Alert

	coll := r.client.Database(r.database).Collection(r.collection)

	// ID'ye göre MongoDB'de arama yap
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&alert)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No alert found with ID: %v", id)
			return nil, err
		}
		log.Printf("Failed to retrieve alert with ID: %v, error: %v", id, err)
		return nil, err
	}

	return &alert, nil
}
