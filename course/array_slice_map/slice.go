package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1 := [3]int{1,2,3}
	slice1 := []int{1,2,3}

	fmt.Println(arr1, slice1)
	fmt.Println(reflect.TypeOf(arr1), reflect.TypeOf(slice1))
	
	arr2 := [5]int{1,2,3,4,5}

	// Slice is not an array, slice defines a piece of an array
	slice2 := arr2[1:3]
	fmt.Println(arr2, slice2)

	slice3 := arr2[:2] // new slice, but is a pointer to the same array
	fmt.Println(arr2, slice3)

	// you can imagine a slice how: size a pointer to an element of the array
	slice4 := slice2[:1]
	fmt.Println(slice2, slice4)
}
