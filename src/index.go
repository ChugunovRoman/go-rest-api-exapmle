package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	C "./Conrtollers"
)

func main() {
	e := echo.New()

	e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Logger())

	h := &C.Conrtollers{DB: GetDb()}

	routePrefix := ""
	if os.Getenv("APP_ROUTE_PREFIX") != "" {
		routePrefix = os.Getenv("APP_ROUTE_PREFIX")
	} else {
		routePrefix = "/api/v1"
	}

	e.GET(routePrefix+"/card", h.GetCard)
	e.POST(routePrefix+"/card", h.CreateCard)
	e.DELETE(routePrefix+"/card/:id", h.DeleteCard)
	e.PUT(routePrefix+"/card/:id", h.UpdateCard)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
