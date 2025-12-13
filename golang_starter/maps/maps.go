package main

import "fmt"

func main () {
	// emails := make(map[string]string)

	// //assign kv
	// emails["bob"] = "bob@gmail.com"

	emails := (map[string]string {"bob": "bob@gmail.com"})

	fmt.Println(emails["bob"])
	
}