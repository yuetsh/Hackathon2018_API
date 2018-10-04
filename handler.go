package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler struct{}

func (h *Handler) listMemes(r *http.Request) interface{} {
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		return err
	}
	var memes []string
	for _, f := range files {
		if f.Name() == ".DS_Store" {
			continue
		}
		memes = append(memes, f.Name())
	}
	return memes
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
