package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return;
		case ch <- fmt.Sprintf("time started at %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pingerChannel := make(chan string)
	done := make(chan struct{})

	go ping(ctx, pingerChannel)

	go func() {
		timeout := time.After(5 * time.Second)
		defer cancel()
		for {
			select {
			case <-timeout:
				fmt.Println("operation completed")
				close(pingerChannel)
				done <-struct{}{}
				return
			case msg := <-pingerChannel:
				fmt.Println("msg", msg)
			}
		}
	}()


	<- done
}