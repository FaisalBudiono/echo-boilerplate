package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ControllerContext struct{}

func (c ControllerContext) Index(ctx echo.Context) error {
	something := struct {
		Foo string `json:"foo"`
	}{
		Foo: "bar",
	}

	return ctx.JSON(http.StatusOK, something)
}
