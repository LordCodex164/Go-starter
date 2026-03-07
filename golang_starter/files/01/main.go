package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	filePath := "files/example.txt"
	content  := []byte("hello world")
	err := os.WriteFile(filePath, content, 0644)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done Writing!!")

	body, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

	filePath2 := "files/example2.txt"

	file, err := os.Create(filePath2)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.WriteString("testing content\n testing content 2\n")

	if err != nil {
		log.Fatal(err)
	}

	printContent(filePath2)

	newFile, err := os.OpenFile(filePath2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	_, err = newFile.WriteString("testing content 3")
	
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done Appending!!")

}

func printContent (filePath string) {
	newFile, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	scanner := bufio.NewScanner(newFile)

	lineNum := 0

	for scanner.Scan() {
		lineNum++
		log.Printf("Line %d: %s\n", lineNum, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}