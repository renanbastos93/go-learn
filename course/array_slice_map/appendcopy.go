package main

import (
	"fmt"
)

func main() {
	arr1 := [3]int{1,2,3}
	var slice1 []int

	// arr1 = append(arr1, 4,5,6) -> not possible
	slice1 = append(slice1, 4,5,6)
	fmt.Println(arr1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)

}
