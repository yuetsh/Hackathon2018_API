package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
)

type UploadData struct {
	Id     string `json:"id"`
	Link   string `json:"link"`
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
}

type UploadRes struct {
	Data    UploadData `json:"data"`
	Success bool       `json:"success"`
	Status  int        `json:"status"`
}

func UploadGif(path string) (*UploadData, error) {
	url := "https://api.imgur.com/3/image"
	accessToken := "02c650a94bbbd7dc6011ffb4fe759d4c03323197"
	album := "gHa5oze"

	reg := regexp.MustCompile(`[0-9a-z]{32}\.gif`)
	filename := reg.FindString(path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	writer.WriteField("album", album)
	writer.WriteField("name", filename)
	writer.WriteField("type", "gif")
	writer.WriteField("title", "hackathon2018")
	writer.WriteField("description", "zhenxiang app - hackathon 2018 for xianghuanji")

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	uploadRes := new(UploadRes)
	json.Unmarshal(data, uploadRes)
	if uploadRes.Success {
		return &uploadRes.Data, nil
	} else {
		return nil, errors.New(string(uploadRes.Status))
	}
}
