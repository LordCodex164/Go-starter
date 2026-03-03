package main

import "fmt"

func main() {
	//create a buffered channel with a capacity of 3
	messages := make(chan string, 3)

	messages <- "this is the first message"
	messages <- "this is the second message"
	messages <- "this is the third message"
	//here the buffered channel gets blocked when it becomes full
	messages <- "this is the fourth message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
