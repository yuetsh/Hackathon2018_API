package api

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, name from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Name)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	stmt, err := Db.Prepare("insert into posts (content, name) values ($1, $2) returning id")
	if err != nil {
		return
	}
	err = stmt.QueryRow(post.Content, post.Name).Scan(&post.Id)
	return
}
