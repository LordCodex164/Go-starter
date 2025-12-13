package main

import (
	"fmt"
	"log"
	"math/rand"
)

type Ability interface {
	ThrowIn()
}

type Player struct {
	stamina int
	power int
}

func (p *Player) ThrowIn() {
	shot := p.power + p.stamina
	fmt.Println("I am kicking the ball", shot)
}

/* 

[{stamina: 10, power: 20}]

*/

func main () {
	team := make([]Ability, 5)
	for i := 0; i < len(team); i++ {
		team[i] = &Player{
			stamina: rand.Intn(10),
			power: rand.Intn(10),
		}
		//team[i].PlayBall()
	}
	log.Println(team)
	// for i :=0; i < len(team); i++ {

	// }
}
