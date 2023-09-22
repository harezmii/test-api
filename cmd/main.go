package main

import (
	"github.com/labstack/echo/v4"
	"guacamole_api/internal/controller"
	"guacamole_api/internal/router"
	"log"
)

func main() {
	app := echo.New()

	c := controller.Controller{
		Url: "url_adres",
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
