package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func print(value int) {
	fmt.Println(value)
}

func main() {
	result := sum(3, 4)
	print(result)
}