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

func main() {
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	mySlice := arr1[2:5]

	fmt.Println("s1", mySlice)
	fmt.Println("len(s1)", len(mySlice))
	fmt.Println("cap(s1)", cap(mySlice))
  }

/* 
3. creating a slice With The make() Function
*/

// func makeSlice() {
// 	data := []int{1, 2, 3, 4 , 5}
// 	for index, value := range data {
// 		fmt.Printf("Index: %d, Value: %d\n", index, value)
// 	}
// }

// func main () {
// 	makeSlice()
// 	slice1 := make([]int, 5, 10)
// 	//it is important to know the capacity doubles when there is a new memory allocation
// 	fmt.Println(slice1, "slice1")
// 	fmt.Println("cap", cap(slice1))

// }
