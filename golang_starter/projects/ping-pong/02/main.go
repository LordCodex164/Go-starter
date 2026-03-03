package main

import (
	"context"
	"fmt"
	//"sync"
	"time"
)

func ping(ctx context.Context, dataCh chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ping is done")
			return;
		case dataCh <- fmt.Sprintln("writing ping ping:", time.Now()):
			time.Sleep(2 * time.Second)

		}
	}
}


func pong(ctx context.Context, dataCh chan string){
	for {
		select {
		case <-ctx.Done():
			fmt.Println("pong is done")
			return;
		case dataCh <- fmt.Sprintln("writing pong pong:", time.Now()):
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	dataChan := make(chan string)
	done := make(chan bool)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go ping(ctx, dataChan)
	go pong(ctx, dataChan)

	//func to handle the messages from them

	go func(){
		timeout := time.After(time.Second * 5)

		for {
			select{
			case <-timeout:
				fmt.Println("operation completed")
				done <- true
				close(dataChan)
				
				//close tha channels
			case data := <- dataChan:
				fmt.Println(data, ">>>>")
			}
		}
	}()

	<-done
	fmt.Println("end")
}