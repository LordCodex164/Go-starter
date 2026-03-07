package main

import (
	"encoding/base64"
	"log"
	"os"
)

func main() {
	
	fileBytes, err := os.ReadFile("dataencoding/base64/example2.txt"); 

	if err != nil {
		log.Fatal(err)
	}

	encoding := base64.StdEncoding.EncodeToString(fileBytes);

	payload := map[string]string{
		"filename": "dataencoding/base64/example2.txt",
		"file": encoding,
	}

	log.Printf("filename: %s with its encoded value %v", payload["filename"], payload["file"])

	decodedValue, err := base64.StdEncoding.DecodeString(payload["file"])

	if err != nil {
		log.Fatal(err)
	}

	if(string(decodedValue)) != string(fileBytes) {
		log.Println("content in the file does not match the decoded value")
	}

	log.Println(string(decodedValue))


}