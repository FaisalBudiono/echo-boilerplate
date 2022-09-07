package bootstraps

import (
	"echo-boilerplate/bootstraps/routing"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func createRoute() http.Handler {
	router := echo.New()
	router.Logger.SetPrefix(getAppName())
	router.Logger.SetLevel(log.DEBUG)

	router.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewString()
		},
	}))
	router.Use(middleware.Logger())

	routing.HealthCheckRoute{}.BindRoute(router.Group(""))

	return router
}
