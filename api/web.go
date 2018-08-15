package api

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func jsonExample(w http.ResponseWriter, r *http.Request)  {
	post := &Post{Id: 1, Name: "s", Content: "qw"}
	data, _ := json.Marshal(post)
	w.Write(data)
}

func StartWeb() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/hello/", jsonExample)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
