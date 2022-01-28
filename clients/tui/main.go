package main

import (
	"fmt"
	"log"

	"github.com/Xuanwo/go-locale"
)

func main() {
	tag, err := locale.Detect()
	if err != nil {
		log.Fatal(err)
	}
	// Have fun with language.Tag!

	tags, err := locale.DetectAll()
	if err != nil {
		log.Fatal(err)
	}
	// Get all available tags
	fmt.Println(tag)
	fmt.Println(tags)
}
