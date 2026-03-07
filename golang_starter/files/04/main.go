package main

import (
	"embed"
	"fmt"
	"log"
)

//embedding static assets into your go program

var name = "Joseph"

//go:embed public
var embedData string

var p embed.FS

func main() {

	data, err := p.ReadFile("public/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println(embedData)
}