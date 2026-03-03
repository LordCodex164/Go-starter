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
	totalJobs := 5

	for i := 1; i < totalJobs; i++ {
		wg.Add(1)
		greet(fmt.Sprintf("%d: say hello", i), &wg)
	}
	//communicaates with the main thread to wait before exiting finally
	wg.Wait()
	fmt.Println("end")
}