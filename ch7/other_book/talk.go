package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "Glurb! Glurb!"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

func main() {
	var t talker
	t = martian{}
	fmt.Println(t.talk())


	t = laser(3)
	fmt.Println(t.talk())

	
}
