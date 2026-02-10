package main

import (
	"fmt"
)

//here we change using address
func changeValuePointer (value *string) {
	//since value here is the memory address 
	//to modify, dereference
	*value = "new-bar"
}


func main(){

	var foo = "bar"

	var fooAdr = &foo // the memory address //0x14811234..

	changeValuePointer(fooAdr)

	fmt.Println("foo", foo) //"new-bar"
}