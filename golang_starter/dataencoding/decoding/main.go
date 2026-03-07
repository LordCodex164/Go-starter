package main

import (
	"encoding/json"
	"log"
	"strings"
	//"os"
)

type Profile struct {
	EmailAddress string `json:"email"`
	Password     string  `json:"password"`
}

var payload = `
 {
  "email":"test@gmail.com",
  "password":"3212po"
}
`

//decoding and encoding is more flexible because one can read from connection and file disk 

func main() {
	
	var p Profile

	if err := json.NewDecoder(strings.NewReader(payload)).Decode(&p); err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", p)

	log.Println("done")
}
