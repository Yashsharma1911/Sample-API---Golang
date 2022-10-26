package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// model for sending back response message
type SendMessage struct {
	Message   string   `json:"message,omitempty"`
	DataAdded *Netflix `json: "data_added`
}

// model for netflix
type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
