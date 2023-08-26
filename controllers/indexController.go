package controllers

import (
	"github.com/labstack/echo/v4"
)

// IndexStatus godoc
//
// @Router / [get]
// @Tags Root
// @Summary Índice del servicio
// @Description Este método devuelve un archivo html que contiene documentación e información relevante sobre el servicio API
// @Accept json
// @Produce html
// @Success 200 {object} string
func IndexController(c echo.Context) error { return c.File("index.html") }
