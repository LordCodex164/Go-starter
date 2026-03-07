package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text1 := "Hello world!"

	regGo, err := regexp.Compile("Hello world")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//

	fmt.Printf("Text %s, matches 'world': %t\n", text1, regGo.Match([]byte(text1)))

	//find the nearest product code
	text2 := "Products Code: C789, E123, NP099"

	findProductCode := regexp.MustCompile(`E\d+`)

	nearestProduct := findProductCode.FindString(text2)

	fmt.Println(nearestProduct)
	
	//
}