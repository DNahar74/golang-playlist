package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Netflix represents a Netflix movie model, just showing whether the movie has been watched or not
type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
