package main

import (
	"bytes"
	"encoding/json"
	"log"
	//"os"
)

type Profile struct {
	EmailAddress string `json:"email"`
	Password     string  `json:"password"`
}

func main() {
	p := Profile{
		EmailAddress: "test@gmail.com",
		Password: "3212po",
	}

	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(p); err != nil {
		log.Fatal(err)
	}

	log.Println(">>", buf.String())

	log.Println("done")
}