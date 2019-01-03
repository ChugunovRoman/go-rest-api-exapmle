package main

import (
	"os"

	"github.com/labstack/gommon/log"
	mgo "gopkg.in/mgo.v2"
)

// GetDb - get database object
func GetDb() *mgo.Session {
	println("MONGO URI: " + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))

	db, err := mgo.Dial(os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))

	if err != nil {
		println("Mongodb connect error ", err)
		log.Fatal(err)
	}

	if err = db.Copy().DB("CardsDB").C("Cards").EnsureIndex(mgo.Index{
		Key:    []string{"code"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	if err = db.Copy().DB("CardsDB").C("Users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	return db
}
