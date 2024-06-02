package mongo_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Request struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RequestType string             `bson:"requestType"`
	Details     string             `bson:"details"`
	CreatedAt   time.Time          `bson:"createdAt"`
}
