package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Handler struct{}

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to the secret garden.")
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createMeme(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		m := new(Meme)
		if err := json.NewDecoder(r.Body).Decode(m); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		} else {
			if err := m.New(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				name := m.hash + "." + m.Format
				if data, err := ioutil.ReadFile("./dist/" + m.Name + "/" + name); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					w.Header().Set("Content-Length", fmt.Sprint(len(data)))
					w.Header().Set("Content-Disposition", "attachment; filename="+name)
					contentType := "image/gif"
					if m.Format == "mp4" {
						contentType = "video/mp4"
					}
					w.Header().Set("Content-Type", contentType)
					w.Write(data)
				}
			}
		}
	}
}
