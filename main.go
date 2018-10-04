package main

import (
	"log"
	"net/http"
	"os"
)

var h Handler

func init() {
	if _, err := os.Stat("./dist"); os.IsNotExist(err) {
		os.Mkdir("./dist", 0700)
	}
}

func main() {
	addr := ":3010"
	logger := log.New(os.Stdout, "", log.LstdFlags)

	mux := http.NewServeMux()

	mux.Handle("/zhenxiang/v1/meme", Adapt(h.createMeme, UseMethod(http.MethodPost), Logging(logger), API(true)))
	mux.Handle("/zhenxiang/v1/memes", Adapt(h.listMemes, UseMethod(http.MethodGet), Logging(logger), API(true)))

	log.Fatal(http.ListenAndServe(addr, mux))
}
