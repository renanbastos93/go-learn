package main

import (
	"fmt"
)

func main() {
	sliceMake := make([]int, 10)
	sliceMake[9] = 12
	fmt.Println(sliceMake)
	
	sliceMake = make([]int, 10, 20)
	// len == length
	// cap == capacity
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake = append(sliceMake, 1,2,3,4,5,6,7,8,9,0)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake = append(sliceMake, 1)
	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))
}
