package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Users - model for Users collection
	Users struct {
		ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Name      string        `json:"name" bson:"name"`
		FirstName string        `json:"firstname" bson:"firstname"`
		Birth     string        `json:"birth" bson:"birth"`
		Email     string        `json:"email" bson:"email"`
		CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
		UpdatedAt time.Time     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	}
)
