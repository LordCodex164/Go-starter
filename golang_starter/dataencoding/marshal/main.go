package main

import (
	"encoding/json"
	"log"
)

type BankAccount struct {
	Name string 		`json:"name" xml:"name"`
	Balance float64		`json:"balance" xml:"balance"`
}


func main() {
	baAccount := BankAccount{
		Name: "Daniel Adeniran",
		Balance: 12.45,
	}

	//byte format
	dataByte, err := json.MarshalIndent(baAccount, "", "");

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(dataByte))
}