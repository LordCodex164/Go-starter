package main

import (
	"encoding/json"
	"log"
)

type BankAccount struct {
	Name string 		`json:"name" xml:"name"`
	Balance float64		`json:"balance" xml:"balance"`
	BankCode int64      `json:"-" xml:"-"`
}

var payload = `{
	"name": "Daniel Adeniran",
	"balance": 12.45,
	"bankcode": 476
}`

func main() {
	
	var b BankAccount

	//byte format
	err := json.Unmarshal([]byte(payload), &b);

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v:", b)
}