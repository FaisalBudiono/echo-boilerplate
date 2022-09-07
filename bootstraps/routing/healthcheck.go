package routing

import (
	"echo-boilerplate/app/healthcheck"

	"github.com/labstack/echo/v4"
)

type HealthCheckRoute struct{}

func (h HealthCheckRoute) BindRoute(g *echo.Group) {
	controller := healthcheck.ControllerContext{}

	g.GET("/", controller.Index)
}
