package main

import "fmt"

type LogLevel int

//iota is a way of declaring enums

const (
	LogError LogLevel = iota + 1
	LogWarn
	LogInfo
	LogFatal
	LogDebug
)

func main() {
	fmt.Println("1", LogError)
	fmt.Println("2", LogWarn)
}