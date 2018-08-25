package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var h Handler

func serverMux(e *echo.Echo) {
	e.GET("/ping", h.ping)
}

func init() {

}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	serverMux(e)
	e.Logger.Fatal(e.Start(":3010"))
}
