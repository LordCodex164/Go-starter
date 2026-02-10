package main

import (
	"fmt"
	"sort"
	"strings"
)

//since the variable is declared outside the application, use var as the keyword

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

var revArray = []string{"running", "jogging", "walking"}

func longestCommonSuffix(strs []string) string {
	reversed_copy := []string{}
	for _, str := range(strs) {
		fmt.Println("len", len(strs))
		r := []rune(str) 
		fmt.Println("last_str", len(str) - 1)
		for i, j:= 0, len(str)-1; i < j; i, j = i+1, j-1 {
			fmt.Println("str", string(r[i]), string(r[j]))
			strings.TrimSpace(string(r))
			r[i], r[j] = r[j], r[i]
		}
		reversed_copy = append(reversed_copy, string(r))
	}
	prefix := longestCommonPrefix(reversed_copy)
	rPrefix := []rune(prefix)
	for i, j:= 0, len(prefix) - 1; i < j; i,j = i + 1, j-1 {
		rPrefix[i], rPrefix[j] = rPrefix[j], rPrefix[i]
	}
	return string(rPrefix)
}

 var strs = []string{"interview", "interrupt", "integrate"}

func shareCommonPrefix(strs []string, prefix string) bool {
	isCommon := true
	for _, str := range(strs) {
		if strings.HasPrefix(str, prefix) {
			isCommon = true
		} else {
            isCommon = false
		}
	}
	return isCommon
}

// the time_complexity is 0(nlog n + m)

func main() {
	fmt.Println("starting the application")
	//fmt.Println("prefix", longestCommonPrefix(strArray))
	fmt.Println("reverse", longestCommonSuffix(revArray))
	fmt.Println("shareCommon", shareCommonPrefix(strs, "int"))
}