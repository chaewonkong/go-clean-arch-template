package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func BindRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
}
