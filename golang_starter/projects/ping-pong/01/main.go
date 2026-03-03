package main

import (
	"context"
	"fmt"
	"time"
)

func fetchAPI(ctx context.Context, dataCh chan string) {
	for {
		select {
		case <- ctx.Done():
			return
		case dataCh <- fmt.Sprintf("data 1 is being fetched at: %v every 3 seconds", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func readData (ctx context.Context, dataCh chan string) {
	for {
		select {
		case <- ctx.Done():
			return;
		case data := <-dataCh:
			fmt.Println("reading", data)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dataChan := make(chan string)
	done := make(chan bool)
	//here we fetch data from an api
	go fetchAPI(ctx, dataChan)
	//we read the data
	go readData(ctx, dataChan)

	//here the goroutine helps to cancel the context only after 9 seconds to signal to 
	// the fetchAPI goroutine to stop running, fetches 4 messages from the channel
	go func ()  {
		defer cancel()
		timeout := time.After(9 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("operation completed")
				done <-true
				close(dataChan) // close to avoid leaks
				return
			case msg := <-dataChan:
				fmt.Println(msg)
			}
		}
	}()

	<-done
}
