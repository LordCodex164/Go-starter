package main

import (
	"fmt"
	"time"
)

func main() {
	switch t := time.Duration(3 * time.Second); {
	case t < 1:
		fmt.Println("the duration is less than the actual duration")
	case t == 3:
		fmt.Println("the duration is the actual duration")
	default:
		fmt.Println("duration unknown")
	}

	checkType := func(value interface{}) {
		switch vType := value.(type) {
		case int:
			fmt.Println("Integer", vType)
		case string:
			fmt.Println("string", vType)
		case bool:
			fmt.Println("bool", vType)
		}
	}

	checkType(3)
	checkType(true)
}
 