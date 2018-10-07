package main

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDB() {
	conn := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(os.Getenv("ENV") + " database is connected!")
}
