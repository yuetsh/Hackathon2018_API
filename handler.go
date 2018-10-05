package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

type Asset struct {
	Name         string   `json:"name"`
	NameEN       string   `json:"name_en"`
	Cover        string   `json:"cover"`
	Gif          string   `json:"gif"`
	SubLen       int      `json:"subs_length"`
	Placeholders []string `json:"Placeholders"`
}

func (h *Handler) listMemes(r *http.Request) interface{} {
	return [2]Asset{
		{
			Name:   "真香",
			NameEN: "Zhen Xiang",
			Cover:  "https://i.imgur.com/JpD5jcp.png",
			Gif:    "https://i.imgur.com/vTTHmY7.gif",
			SubLen: 4,
			Placeholders: []string{
				"我王境泽就是饿死",
				"死外面 从这里跳下去",
				"也不会吃你们一点东西",
				"真香",
			},
		},
		{
			Name:   "有钱可以为所欲为",
			NameEN: "Rich Men Can Do Anything He Wants",
			Cover:  "https://i.imgur.com/wwaBHEM.png",
			Gif:    "https://i.imgur.com/7eRIgA5.gif",
			SubLen: 9,
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
