package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var h Handler

func init() {
	if _, err := os.Stat("./dist"); os.IsNotExist(err) {
		os.Mkdir("./dist", 0700)
	}

	conn := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected!")
}

func main() {
	addr := ":3010"
	logger := log.New(os.Stdout, "", log.LstdFlags)

	mux := http.NewServeMux()

	mux.Handle("/zhenxiang/v1/meme", Adapt(h.createMeme, UseMethod(http.MethodPost), Logging(logger), API(true)))
	mux.Handle("/zhenxiang/v1/memes", Adapt(h.listMemes, UseMethod(http.MethodGet), Logging(logger), API(true)))

	log.Fatal(http.ListenAndServe(addr, mux))
}
