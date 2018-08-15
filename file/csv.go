package file

import (
	"os"
	"encoding/csv"
	"strconv"
	"fmt"
	"github.com/yuetsh/Hackathon2018/api"
)

func CreateCSV(data []api.Post, filename string) {
	file, _ := os.Create(filename + ".csv")
	defer file.Close()
	csv.NewWriter(file)
	writer := csv.NewWriter(file)
	for _, post := range data {
		line := []string{strconv.Itoa(post.Id), post.Name, post.Content}
		writer.Write(line)
	}
	writer.Flush()
}

func ReadCSV(filename string) {
	file, _ := os.Open(filename + ".csv")
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	record, _ := reader.ReadAll()
	var posts []api.Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := api.Post{int(id), item[1], item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts)
}
