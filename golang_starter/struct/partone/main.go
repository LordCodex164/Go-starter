package main

import (
	"encoding/json"
	"fmt"
)

type Agent struct {
	Name   string `json:"name"`
    Status string `json:"status"`
}

type Player struct {
	Position string `json:"position"` //specify the tags
	Age      int    `json:"age"`
	Name     string `json:"name"`
	Agent    
}

//in this a is the identifier to the struct Agent
//here we modify the struct by reference
func (a *Agent) getAgent() *Agent {
	return a
}

//modifying by value


//modifying by reference

func main() {

	agent := Agent{
		Name: "jorge menedes",
		Status: "senior_level",
	}

	player := Player{
		Position: "Left Back",
		Age:      23,
		Name:     "Guler",
		Agent: 	  agent,
	}
	jsonData, _ := json.Marshal(player)
	fmt.Println("json", string(jsonData))
	fmt.Print(">>print agent", player.getAgent())
}