package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	Tags      []string           `bson:"tags"`
}
