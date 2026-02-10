package main

import "fmt"

func main () {
	emails := make(map[string]string)

	//to get the type of the map

	fmt.Printf(">>emails: %T\n", emails)

	// //assign kv
	emails["bob"] = "bob@gmail.com"

	//value, ok := emails["bob"]

	//emails := (map[string]string {"bob": "bob@gmail.com"})

	fmt.Println(emails["bob"])

	//elements in a map cannot be in the same order
	
}