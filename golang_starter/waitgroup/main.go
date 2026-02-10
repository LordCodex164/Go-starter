package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchAPI(d time.Duration, wg *sync.WaitGroup) {
	fmt.Println("Fetching Data...")
	time.Sleep(time.Second * d)
	fmt.Println("Data fetched Successfully...")
	wg.Done()
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	go fetchAPI(5, &wg)
	go fetchAPI(5, &wg)
	wg.Wait()
	// before running the main function
	fmt.Printf("the process took %d microseconds", time.Since(start))
}