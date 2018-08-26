package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	m := Meme{
		Name: ZhenXiang,
		Subs: []string{
			"上下",
			"左右",
			"转圈",
			"牛逼111",
		},
	}

	if err := m.New(); err != nil {
		t.Error(err)
		fmt.Print(m.hash)
	}
}
