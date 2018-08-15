package main

import (
	"fmt"
	"database/sql"
	"github.com/yuetsh/Hackathon2018/api"
	"github.com/yuetsh/Hackathon2018/file"
)

func init() {
	var err error
	api.Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func main() {
	post := api.Post{Id: 1, Name: "Hello", Content: "World"}
	file.Store(post, "post1")
	var post1 api.Post
	file.Load(&post1, "post1")
	fmt.Println(post1)
	allPost := []api.Post{
		{Id: 1, Name: "Hello", Content: "World"},
		{Id: 1, Name: "Hello", Content: "World"},
		{Id: 1, Name: "Hello", Content: "World"},
		{Id: 1, Name: "Hello", Content: "World"},
	}
	
	file.CreateCSV(allPost, "post1")
	file.ReadCSV("post1")
	
	api.StartWeb()

	post2 := api.Post{Id: 1, Name: "xu", Content: "yue"}
	post2.Create()

	posts, err := api.Posts(10)
	if err != nil {
		return
	}
	fmt.Print(posts)
}
