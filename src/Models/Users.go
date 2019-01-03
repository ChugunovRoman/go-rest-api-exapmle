package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Users - model for Users collection
	Users struct {
		ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty" valid:"-"`
		Name      string        `json:"name" bson:"name" valid:"-"`
		FirstName string        `json:"firstname" bson:"firstname" valid:"-"`
		Email     string        `json:"email" bson:"email" valid:"email"`
		Birth     time.Time     `json:"birth" bson:"birth" valid:"-"`
		CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty" valid:"-"`
		UpdatedAt time.Time     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" valid:"-"`
	}
)
