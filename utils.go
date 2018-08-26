package main

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func NewMd5(subs []string) string {
	cipher := md5.New()
	text := []byte(strings.Join(subs, ","))
	cipher.Write(text)
	return hex.EncodeToString(cipher.Sum(nil))
}
