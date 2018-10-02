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

	mux.Handle("/meme", Adapt(h.createMeme, Logging(logger), UseMethod(http.MethodPost)))
	mux.Handle("/memes", Adapt(h.listMemes, Logging(logger), UseMethod(http.MethodGet), API(true)))

	log.Fatal(http.ListenAndServe(addr, mux))
}
