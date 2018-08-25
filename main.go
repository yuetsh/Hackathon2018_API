package main

import (
	"os"
	"net/http"
	"log"
)

var h Handler

func init() {
	if _, err := os.Stat("./dist"); os.IsNotExist(err) {
		os.Mkdir("./dist", 0700)
	}
}

func main() {
	http.HandleFunc("/ping", h.ping)
	http.HandleFunc("/meme", h.createMeme)
	log.Fatal(http.ListenAndServe(":3010", nil))
}
