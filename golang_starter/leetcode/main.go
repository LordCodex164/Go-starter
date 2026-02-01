package main

import (
	"fmt"
	"sort"
)

//since it is declared outside the application

var strArray = []string{"hello", "hell"}

func longestCommonPrefix(strs []string) string {
	//arrange in alphabetical order
	sort.Strings(strs)
	//compare the first and last string

	first := strs[0]
	last := strs[len(strs) - 1]

	//matching the character 

	i := 0

	for i < len(first) && first[i] == last[i] {
		i++
	}

	return first[:i]

}

func main() {
	fmt.Println("starting the application")
	fmt.Println("prefix", longestCommonPrefix(strArray))
}