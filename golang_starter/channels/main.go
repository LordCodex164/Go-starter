package main

import (
	"fmt"
	"time"
)

type cv struct {
	name string
	social_links string
}

func main() {
	//rules when dealing with channels
	//1. pulling in, the arrow is on the right
	//2. pulling out, the arrow is on the left
	messages := make(chan string)
	cvChannel := make(chan cv)

	go func(){
		//write or pull in to the messages channel
		messages <- "channel data"
	}()

	go func() {
		cvChannel <- cv{
			name: "daniel adeniran",
			social_links: "twitter.com",
		}
	}()
   time.Sleep(1 * time.Second)
   //here the messages is pulled out of the channel
	msg := <- messages
	fmt.Println("msg", msg)

	cvData := <- cvChannel

	fmt.Println("cv_data", cvData)
	
}