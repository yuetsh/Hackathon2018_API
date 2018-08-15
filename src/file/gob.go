package file

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

func Store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	encoder.Encode(data)
	ioutil.WriteFile(filename, buffer.Bytes(), 0600)
}

func Load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	decoder.Decode(data)
}