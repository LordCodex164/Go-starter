package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Start")

	jobs := make(chan int, 5)
	
	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			j, ok := <-jobs
			if !ok {
				//there are no more jobs
				fmt.Println("Channel Closed")
				return
			}
			fmt.Println("processing job", j)
		}
	}(&wg)

	totalJobs := 3

	for i := 0; i < totalJobs; i++ {
		jobs <- i
		fmt.Println("sending", i)
	}

	close(jobs) //close the channel after it is done processing

	wg.Wait() //wait to receive from the channel

	fmt.Println("End")
}
