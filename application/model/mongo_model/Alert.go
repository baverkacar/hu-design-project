package mongo_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Alert struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	AlertRisk  string             `bson:"alertRisk"`
	AlertType  string             `bson:"alertType"`
	IPAddress  string             `bson:"ipAddress"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	LastUpdate time.Time          `bson:"lastUpdate" json:"lastUpdate"`
}
