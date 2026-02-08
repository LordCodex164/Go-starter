package main

import "fmt"


func main() {
	fmt.Println("-------loop--------")
	counter := 0
	for {
		fmt.Println("Printing counter", counter)
		counter++
		if counter >= 6 {
			fmt.Println("the loops stop at count 5")
			break;
		}
		continue
	}
	fmt.Println("-------skipping--------")

	for i := 0; i< 5; i++ {
		if i == 1 || i== 4 {
			continue;
		}
		fmt.Println("print i", i)
	}
}
