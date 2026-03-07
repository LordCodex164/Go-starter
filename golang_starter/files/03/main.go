package main

import (
	"log"
	"os"
)

func main() {
	//to create a single directory
	if err := os.Mkdir("files/example", 0755); err != nil {
		log.Fatal(err)
	}
	log.Println("directory created")

	//to create a nested directory
	if err := os.MkdirAll("files/example/nested", 0755); err != nil {
		log.Fatal(err)
	}
	log.Println("nested directory created")

	
}