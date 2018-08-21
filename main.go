package main

import (
	"database/sql"
	"fmt"
	"os"
)

func init() {
	var err error
	connection := "user=" + os.Getenv("POSTGRES_USER") + "password=" + os.Getenv("POSTGRES_PASSWORD") + " sslmode=disable"
	Db, err = sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

func main() {
	StartServer()

	posts, err := Posts(10)
	if err != nil {
		return
	}
	fmt.Print(posts)
}
