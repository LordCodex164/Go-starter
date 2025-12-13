package main

import "fmt"

//defining the methods type for the data
type M_TRUCK interface {
	trackIC() error
}

//here implement the methods and we pass the argument a Truck argument to store data
func (t *Truck1) trackIC() error {
	t.spec = "turboshe"
	t.IC_number = 2345
	return nil
}

func (t *Truck2) trackIC() error {
	t.spec = "turboshe"
	t.IC_number = 2345
	return nil
}

//creating a struct for storing the data 
type Truck1 struct {
	id string
	spec string
	IC_number int64 
}

type Truck2 struct {
	id string
	spec string
	IC_number int64 
}

func triggerProcess(m M_TRUCK) error {
	fmt.Println("pleaase hold on!!!")
	err := m.trackIC()
	if err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}
    return nil
}

func main () {
	fmt.Println("PROCESSING TRUCK DATA")
	truck := []Truck1{
		{id: "12", spec: "puma", IC_number: 3421},
	}
	fmt.Println(truck)
	for _, t := range truck {
		triggerProcess(&t)
	}
	fmt.Println(truck)
}
