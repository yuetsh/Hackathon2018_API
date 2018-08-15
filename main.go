package main

import (
	"database/sql"
	"github.com/yuetsh/Hackathon2018/api"
	"fmt"
	"os"
)

func init() {
	var err error
	connection := "user="+os.Getenv("POSTGRES_USER")+"password="+os.Getenv("POSTGRES_PASSWORD")+" sslmode=disable"
	api.Db, err = sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

func main() {
	api.StartWeb()

	post1 := api.Post{Id: 1, Name: "xu", Content: "yue"}
	post1.Create()

	posts, err := api.Posts(10)
	if err != nil {
		return
	}
	fmt.Print(posts)
}
