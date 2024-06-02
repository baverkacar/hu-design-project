package mongo_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OSName    string             `bson:"osName"`
	IPAddress string             `bson:"ipAddress"`
}
