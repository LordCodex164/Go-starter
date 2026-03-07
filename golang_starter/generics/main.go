package main

import "fmt"


type Number interface {
	int | float64
}

func sum[T int | float64](numbers ...T) T {
	var total  T
	for _, number :=  range(numbers) {
		total += number
	}
	return total
}

func fetchData[T string | float64](values T){
	fmt.Printf("fetching data with this key value: %v\n", values)
}

func main(){
	fetchData(12.56)
	sum(23.4, 45)
}