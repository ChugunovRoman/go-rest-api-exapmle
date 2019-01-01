package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	mgo "gopkg.in/mgo.v2"

	C "./Conrtollers"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Logger())

	db, err := mgo.Dial("localhost" + ":" + "27017")
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

	h := &C.Conrtollers{DB: db}

	routePrefix := "/api/v1"

	e.GET(routePrefix+"/card", h.GetCard)
	e.POST(routePrefix+"/card", h.CreateCard)
	e.DELETE(routePrefix+"/card/:id", h.DeleteCard)
	e.PUT(routePrefix+"/card/:id", h.UpdateCard)

	e.Logger.Fatal(e.Start(":5000"))
}
