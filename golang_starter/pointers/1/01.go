package main

import "fmt"

//here we modify using value
func changeValue(value int) {
	value = 12
}


func main() {
	value := 10

	// this holds the pointer to the variable(stores the memory address of the variable)
	pointerValue := &value //0x1459000..

	//the dereference(we go to where the memory address of the value is stored and read the value)
	fmt.Println("pointe", *pointerValue) 

	changeValue(*pointerValue)

	fmt.Println("new_value", value) // still remains 10

}