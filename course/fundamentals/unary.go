package main

import (
	"fmt"
)


func main() {
	x, y := 1, 2

	// only postfix
	x++ // x += 1 or x = x + 1
	fmt.Println(x)

	y-- // y -= 1 or y = y - 1
	fmt.Println(y)
}