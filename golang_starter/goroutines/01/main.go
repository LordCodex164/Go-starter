package main

import (
	"fmt"
	"sync"
)

//take it in the original waitgroup(reference)
func greet(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

 




func main() {
	fmt.Println("Start")
	var wg sync.WaitGroup
	wg.Add(3)
	go greet("say hello", &wg)
	go greet("say hello", &wg)
	go greet("say hello", &wg)
	//communicaates with the main thread to wait before exiting finally
	wg.Wait()
	fmt.Println("end")
}