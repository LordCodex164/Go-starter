package main

import (
	"errors"
	"fmt"
	"log"
)

type Contract struct {
	name string
	email string
	phone  string
}

var contactList []Contract
var contactIndexByName map[string]int

func init() {
	contactList = make([]Contract, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(payload Contract) ([]Contract, error) {
	if _, exists := contactIndexByName[payload.email]; exists {
			fmt.Println("name already exist")
			return nil, errors.New("name already exists")
	}
	contactList = append(contactList, Contract{
		name: payload.name,
		email: payload.email,
		phone: payload.phone,
	})
	contactIndexByName[payload.email] = len(contactList) - 1
	return contactList, nil
}

func main(){
	contract := Contract{
		name: "dan",
		email: "dan@gmail.com",
		phone: "+234",
	}

	contracts, err := addContact(contract)

	fmt.Println("contracts", contracts)

	contract = Contract{
		name: "dan2",
		email: "dan@gmail.com",
		phone: "+234",
	}

	contracts2, err := addContact(contract)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contracts", contracts2)
}