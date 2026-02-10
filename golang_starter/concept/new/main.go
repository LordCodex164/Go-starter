package main

import "fmt"

type CInterface interface {
	Increment() error
}

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count ++
}

//using the new keyword

func NewCounter() *Counter {
	return new(Counter)
}

func main() {
	counter := NewCounter()  // it returns the allocated address of the struct type(essentially the original) and makes it look as if c.count is 0 when initialized
	counter.Increment() //which is why when the count is modifed using the method, the original struct gets modified
	fmt.Println("count", counter.count)
}
