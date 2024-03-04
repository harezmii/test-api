package router

import (
	"github.com/labstack/echo/v4"
	"guacamole_api/internal/controller"
)

type Router struct {
	App        *echo.Echo
	Controller controller.Controller
}

func (r Router) RouteSetups() {
	r.TestRoute()
	r.GetTokens()
	r.DeleteTokens()
	r.GetConnections()
}
