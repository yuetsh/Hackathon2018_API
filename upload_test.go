package main

import (
	"fmt"
	"testing"
)

func TestUpload(t *testing.T) {
	b, err := UploadGif("./dist/zhenxiang/d2b5956f44e93a640665d457ca5a52e4.gif")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b.Link)
	}
}
