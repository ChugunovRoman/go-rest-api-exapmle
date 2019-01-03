package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Card - model for card collection
	Card struct {
		ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty" valid:"-"`
		Type      string        `json:"type" bson:"type" valid:"in(credit|discont),required"`
		Code      string        `json:"code" bson:"code" valid:"-"`
		Active    bool          `json:"active" bson:"active" valid:"-"`
		CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty" valid:"-"`
		UpdatedAt time.Time     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" valid:"-"`
	}
)
