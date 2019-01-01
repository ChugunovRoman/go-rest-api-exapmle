package conrtollers

import (
	mgo "gopkg.in/mgo.v2"
)

type (
	// Conrtollers -
	Conrtollers struct {
		DB *mgo.Session
	}
)
