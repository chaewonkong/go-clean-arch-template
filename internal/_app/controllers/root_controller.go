package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (controller Controller) GetHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
