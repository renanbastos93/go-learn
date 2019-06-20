package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 10, 20)
	slice2 := append(slice, 1,2,3)
	fmt.Println(slice, slice2)

	slice[0] = 7
	fmt.Println(slice, slice2)
}

// Output
// slice				 slice2
// [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0 1 2 3]
// [7 0 0 0 0 0 0 0 0 0] [7 0 0 0 0 0 0 0 0 0 1 2 3]
