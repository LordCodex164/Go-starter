package main

import "fmt"

type Cargo interface {
	updateContainer() error
	updateVessels() error
}

type Container struct {
	id int
	width int
	height int
}

type Vessels struct {
	id string
	name string
	trackingId int
}

func updateContainer (c *Container) {
	c.height = 1
}

func (v *Vessels) updateVessels() {
	v.name = "test"
}

func main () {
	load := Container{id: 1, width: 20, height: 40}
	updateContainer(&load)
	fmt.Printf("height: %d\n", load.height)
}
