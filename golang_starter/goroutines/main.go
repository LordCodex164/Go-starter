package main

import (
	"fmt"
	"log"
	"sync"
	//"time"
)

//the roles and method for the cargo interface
type Cargo interface {
	updateContainer() error
	updateVessels() error
	updateName()error
	updateTracking() error
}

//an interface for storing data
type Container struct {
	id int
	width int
	height int
}

type Vessels struct {
	id int
	name string
	trackingId int
}

func (c *Container) updateContainer() error {
	c.height = 0
	return nil
}

func (c *Container) updateVessels() error {
	fmt.Println("loading data")
	c.height += 23
	return nil
}

func (c *Container) updateName() error {
	c.height = 0
	return nil
}

func (c *Container) updateTracking() error {
	fmt.Println("loading data")
	c.height += 23
	return nil
}

func (c *Vessels) updateContainer() error {
	c.name = "updated vessels"
	return nil
}

func (c *Vessels) updateVessels() error {
	fmt.Println("loading data")
	c.trackingId += 23
	return nil
}

func (c *Vessels) updateName() error {
	c.name = "updated vessels"
	return nil
}

func (c *Vessels) updateTracking() error {
	fmt.Println("loading data")
	c.trackingId += 23
	return nil
}

//the goal is to reduce 4secs to 1sec 

func processCargo (cargo Cargo) error {
	fmt.Printf("start processing cargo %+v \n", cargo)
	//simulate a processing time
	//time.Sleep(time.Second)
	err := cargo.updateContainer()
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}
	fmt.Printf("Finished processing cargo: %+v \n", cargo)
	return nil
}

func processFleet (cargos []Cargo) error {
	var wg sync.WaitGroup
	for _, c := range cargos {
		//specify how many goroutines you have
		wg.Add(1)
		go func(c Cargo) {
		  if err := processCargo(c); err != nil {
			log.Println(err)
		  } 
		  wg.Done()
		}(c)
	}
	wg.Wait()
	return nil
}

func main() {
	load := []Cargo{
		&Container{id: 1, width: 20, height: 300},
		&Vessels{id: 2, name: "test", trackingId: 100},
		&Container{id: 1, width: 20, height: 300},
		&Vessels{id: 2, name: "test", trackingId: 100},
	}

	if err := processFleet(load); err != nil {
		fmt.Printf("error processing load: %d", err)
	}
	fmt.Println("Processed load successfully")
}