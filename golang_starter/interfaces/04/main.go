package main

import "fmt"

type Person interface {
	String() string
}

type Human struct {
	name string
	age  int
	is_read bool
}

func(h Human) String() string {
	return fmt.Sprintln(h.name + " is the name of this human")
}

type ID int

func (idx ID) String() string {
	return fmt.Sprintf("%d is the ID", idx)
}

func main(){

	human := Human{
		name: "adam mockler",  
	}


	fmt.Println(human)

	id := ID(10)

	fmt.Println(id)
}