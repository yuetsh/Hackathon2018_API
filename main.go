package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

var h Handler

func serverMux(e *echo.Echo) {
	e.GET("/ping", h.ping)
}

func init() {
	if _, err := os.Stat("./dist"); os.IsNotExist(err) {
		os.Mkdir("./dist", 0700)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	serverMux(e)
	e.Logger.Fatal(e.Start(":3010"))
}
