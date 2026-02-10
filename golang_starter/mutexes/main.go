package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Player struct {
	health int32
	mu     sync.RWMutex
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func (p *Player) getHealth() int32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.health
}

func (p *Player) damagePlayer(value int32) {
	health := p.getHealth()
	fmt.Println("health", health, value)
	atomic.StoreInt32(&p.health, int32(health - value))
}
 
func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println("reading player health", p.getHealth())
		<- ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		//damage the player health
		p.damagePlayer(int32(rand.Intn(40)))
		fmt.Println("final_health>>>", p.health)
		if p.health <= 0 {
			fmt.Println("GAME OVER Bitch!!!")
			break;
		}
		<- ticker.C
	}
}

func main() {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
	fmt.Println("Hello, World!")
}

/* 
The problem with race condition is when multiple goroutine read and write a shared state, a
state caan be perceived as something while we don't know it is updated with another oeration

*/