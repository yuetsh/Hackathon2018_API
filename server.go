package main

import (
	"net/http"
	"log"
	"fmt"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!!!")
}

func StartServer() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
