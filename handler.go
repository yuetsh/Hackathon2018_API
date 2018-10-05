package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

type asset struct {
	name  string
	cover string
	gif   string
}

func (h *Handler) listMemes(r *http.Request) interface{} {
	return [2]asset{
		{
			name:  "真香",
			cover: "https://i.imgur.com/JpD5jcp.png",
			gif:   "https://i.imgur.com/TFPQMJm.gif",
		},
		{
			name:  "为所欲为",
			cover: "https://i.imgur.com/wwaBHEM.png",
			gif:   "https://i.imgur.com/vTTHmY7.gif",
		},
	}
}

func (h *Handler) createMeme(r *http.Request) interface{} {
	m := new(Meme)
	if err := json.NewDecoder(r.Body).Decode(m); err != nil {
		return err
	}
	if err := m.New(); err != nil {
		return err
	}
	data, err := UploadGif(m.paths.output.gif)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return data
}
