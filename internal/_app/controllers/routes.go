package controllers

import (
	"github.com/labstack/echo"
)

func BindRoutes(e *echo.Echo, c Controller) {
	e.GET("/", c.GetHealthCheck)
}
