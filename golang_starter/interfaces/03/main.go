package main

import "fmt"

type Modes interface {
	//Playing() error
	Reading() *Human
}

type Person interface {
	Age() int
}

type Human struct {
	name string
	age  int
	is_read bool
}

type SuperHuman struct {
	name string
	age  int
	is_read bool
}

func (h Human) Reading() *Human {
	h.is_read = true
	return &h
}

func (h Human) Age() int {
	return h.age
}

func (h SuperHuman) Age() int {
	return h.age
}

//in this case, h can be human or superman
func DisplayHuman(h Person) int {
	return h.Age()
}



type Sport interface {
	Stadium() 
	//Fans()	  error
	//Player()   error
}

type Football struct {
	player string
}

type Basketball struct {
	player string
}

func (f Football) Stadium() {
	fmt.Println(f.player + " is playing at surelere football stadium")
}

func (b Basketball) Stadium() {
	fmt.Println(b.player + " is playing at Hawks Basketall Academy")
}


func DisplaySport(s Sport) {
	s.Stadium()
}

 
func main(){
	human := Human{
		name: "james",
		age: 2,
		is_read: false,
	}
	
	fmt.Printf("human age: %v\n", DisplayHuman(human))

	superhuman := SuperHuman{
		name: "james",
		age: 2,
		is_read: false,
	}

	fmt.Printf("human age: %v\n", DisplayHuman(superhuman))

	foot := Football{
		player: "yamal",
	}
	
	DisplaySport(foot)

	bskBall := Basketball{
		player: "lebron james",
	}
	
	DisplaySport(bskBall)

}
  