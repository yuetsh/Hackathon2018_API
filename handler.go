package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

type Asset struct {
	Name   string `json:"name"`
	NameEN string `json:"name_en"`
	Cover  string `json:"cover"`
	Gif    string `json:"gif"`
}

func (h *Handler) listMemes(r *http.Request) interface{} {
	return [2]Asset{
		{
			Name:   "真香",
			NameEN: "Zhen Xiang",
			Cover:  "https://i.imgur.com/JpD5jcp.png",
			Gif:    "https://i.imgur.com/vTTHmY7.gif",
		},
		{
			Name:   "为所欲为",
			NameEN: "Rich",
			Cover:  "https://i.imgur.com/wwaBHEM.png",
			Gif:    "https://i.imgur.com/vTTHmY7.gif",
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
