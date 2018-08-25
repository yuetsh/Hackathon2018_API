package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct{}

func (h *Handler) ping(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the secret garden.")
}

func (h *Handler) createMeme(c echo.Context) (err error) {
	t := c.QueryParam("type")
	if t != "gif" || t != "mp4" {
		t = "gif"
	}
	m := new(Meme)
	if err = c.Bind(m); err != nil {
		return err
	}
	if err = m.New(); err != nil {
		return err
	}
	return c.File("./dist/" + m.Name + "/" + m.hash + "." + t)
}
