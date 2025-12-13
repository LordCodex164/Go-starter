package main
import (
	"fmt"
)

//how to create a slice
/* 
1. using the []datatype{values}
*/

// func main () {
// 	slice1 := []int{}
// 	fmt.Println(slice1)
// 	slice2 := []string{"hello", "world"}
// 	fmt.Println(len(slice2))
// }


/* 
2. creating a slice from an array

var myarray = 
*/

// func main() {
// 	arr1 := [6]int{10, 11, 12, 13, 14,15}
// 	myslice := arr1[2:4]
  
// 	fmt.Printf("myslice = %v\n", myslice)
// 	fmt.Printf("length = %d\n", len(myslice))
// 	fmt.Printf("capacity = %d\n", cap(myslice))
//   }

/* 
3. creating a slice With The make() Function
*/

func main () {
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1, "slice1")
	fmt.Println("cap", cap(slice1))
}