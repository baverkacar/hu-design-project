package mongo_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type List struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AlertID   primitive.ObjectID `bson:"alertId"`
	IPAddress string             `bson:"ipAddress"`
	Date      time.Time          `bson:"date"`
}
