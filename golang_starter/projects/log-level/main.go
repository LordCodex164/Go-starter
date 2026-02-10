package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelWarn
	LevelError 
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "unknwown"
	}
	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.String())
}

type PlayerNumber int

const (
	Messi PlayerNumber = iota
	Ronaldo
	Neymar
	Zidane
)

var playerNames = []string{"messi", "ronaldo", "neymar", "zidane"}

func (p PlayerNumber) getPlayerNames() string {
	if p < Messi || p > Zidane {
		return "unknown"
	}
	return playerNames[p]
}

func PrintPlayerNumber(p PlayerNumber) {
	//we print the player number with the name
	fmt.Printf("printing player number %d with name: %s\n", p, p.getPlayerNames())
}

func main() {
	PrintPlayerNumber(Messi)
	PrintPlayerNumber(Ronaldo)
	PrintPlayerNumber(Neymar)
	PrintPlayerNumber(Zidane)
}


