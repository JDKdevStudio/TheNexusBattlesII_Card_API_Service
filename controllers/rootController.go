package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RootController(c echo.Context) error { return c.String(http.StatusOK, "Service is up!") }
