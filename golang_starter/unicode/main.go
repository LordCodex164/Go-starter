package main

import  (
	"fmt"
)

//we only use range to handle data when dealing with unicode like this because we think in terms of bytes and not characters

func main() {
	username := "文字絵"
	for _, v := range(username) {
		fmt.Println(string(v))
	}

	//using rune

	data := []rune{'文', '字', '絵'}
	for _, v := range(data) {
		fmt.Println(string(v))
	}

}