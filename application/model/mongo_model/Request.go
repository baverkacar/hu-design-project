package mongo_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Request struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SourceIP   string             `bson:"sourceIp"`
	DestIP     string             `bson:"destIp"`
	Protocol   string             `bson:"protocol"`
	Time       string             `bson:"time"`
	Date       time.Time          `bson:"date"`
	Permission string             `bson:"permission"`
}
