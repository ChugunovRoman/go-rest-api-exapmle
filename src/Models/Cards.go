package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Card - model for card collection
	Card struct {
		ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Type      string        `json:"type" bson:"type"`
		Code      string        `json:"code" bson:"code"`
		Active    bool          `json:"active" bson:"active"`
		CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
		UpdatedAt time.Time     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	}
)
