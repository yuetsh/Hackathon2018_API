package main

import (
	"testing"
	"fmt"
)

func TestNew(t *testing.T) {

	subs := []string{
		"咋样",
		"一个人加班",
		"写 Hackathon",
		"去群里叫战友",
		"各个回复",
		"没有，下一个",
		"好苦逼啊",
		"真的惨",
		"不写了 不写了",
	}

	if hash, err := New2("sorry", subs); err != nil {
		t.Error(err)
		fmt.Print(hash)
	}
}
