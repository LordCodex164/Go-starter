package main

import (
	"fmt"
	"log"
)

func main () {
	// ids := []int{33, 76, 54, 23, 11, 2}

	// sum := 0

	// for i, id := range ids {
	// 	sum += id
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }
	// fmt.Printf("sum: %d\n", sum)

	// fmt.Println("")
	// emails := map[string]string{"daniel": "daniel@gmail.com", "james": "james@gmail.com"}

	// for k, v := range emails {
	// 	fmt.Printf("key: %s value: %s\n", k, v)
	// }

	person := make(map[string] interface{}, 0)

	person["age"] = 17
	person["name"] = "daniel"

	fmt.Printf("%s\n", person)

	name, exists := person["name"].(string)

	if !exists {
		log.Fatal("name does not exists")
		return;
	}
	fmt.Println(name)

}
