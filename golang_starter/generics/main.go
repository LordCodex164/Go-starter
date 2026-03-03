package main

import "fmt"

func fetchData[T string | float64](values T){
	fmt.Printf("fetching data with this key value: %v\n", values)
}

func main(){
	fetchData(12.56)
}