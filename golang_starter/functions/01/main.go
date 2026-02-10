package main

import (
	"fmt"
	"strings"
)

var numbers []int

func init() {
	numbers = make([]int, 0)
}

func incrementVariable(a *int) func() int {
	return func() int {
		*a++
		return *a
	}
}

//("--------Variadic functions--------")

func sum(numbers ...int) int {
	total := 0
	if len(numbers) > 0 {
		for _, number := range numbers {
			fmt.Println(">>num", number)
			total += number
		}
		fmt.Println("total", total)
		return total
	} else {
		fmt.Println("total", total)
		return total
	}
}

//(-------multi-values--------)

func splitName(fullName string) (firstname, lastname string){
	split := strings.Split(fullName, " ")
	firstname = split[0]
	lastname = split[1]
	return
}

func main() {
	a := 1
	incremFunc := incrementVariable(&a)
	incremFunc()
	fmt.Println(">>", a)
	greeting := func(msg string) *string {
		return &msg
	}

	msg := greeting("hello world")

	fmt.Println("msg", *msg)

	numbers = []int{1, 2, 3, 4, 5}

	sum()

	fmt.Println(splitName("daniel wint"))

}
