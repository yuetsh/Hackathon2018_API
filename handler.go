package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct{}

func (h *Handler) ping(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the secret garden.")
}
