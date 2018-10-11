package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct{}

type Asset struct {
	Name         string   `json:"name"`
	Cover        string   `json:"cover"`
	Gif          string   `json:"gif"`
	Placeholders []string `json:"placeholders"`
}

func (h *Handler) listMemes(r *http.Request) interface{} {
	return [2]Asset{
		{
			Name:  "zhenxiang",
			Cover: "https://i.imgur.com/JpD5jcp.png",
			Gif:   "https://i.imgur.com/vTTHmY7.gif",
			Placeholders: []string{
				"我王境泽就是饿死",
				"死外面 从这里跳下去",
				"也不会吃你们一点东西",
				"真香",
			},
		},
		{
			Name:  "sorry",
			Cover: "https://i.imgur.com/wwaBHEM.png",
			Gif:   "https://i.imgur.com/7eRIgA5.gif",
			Placeholders: []string{
				"好啊",
				"别说我是一等良民",
				"就算你们真的想诬告我",
				"我有的是钱请律师打官司",
				"我想我根本不用坐牢",
				"你别以为有钱了不起啊",
				"Sorry 有钱真的可以为所欲为",
				"不过 我想你不会明白这种感觉",
				"不会 不会",
			},
		},
		//{
		//	Name:  "dagong",
		//	Cover: "",
		//	Gif:   "",
		//	Placeholders: []string{
		//
		//	}
		//}
	}
}

type UploadData struct {
	Id     string `json:"id"`
	Link   string `json:"link"`
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
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
	res := new(UploadData)
	res.Height = data.Height
	res.Width = data.Width
	res.Id = data.Storename
	res.Link = data.Url
	res.Name = data.Filename
	res.Type = "image/gif"
	return res
}
