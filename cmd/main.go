package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"guacamole_api/internal/controller"
	"guacamole_api/internal/router"
	"log"
)

func main() {
	app := echo.New()

	app.Use(middleware.Secure())
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	c := controller.Controller{
		Url: "10.101.0.96:8080",
	}

	r := router.Router{
		App:        app,
		Controller: c,
	}

	r.RouteSetups()

	serverError := app.Start(":8080")
	if serverError != nil {
		log.Println("Server error")
	}
}
