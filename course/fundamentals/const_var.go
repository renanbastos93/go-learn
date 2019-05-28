package main

import (
	"fmt"
	"math"
	// m "math" -> For to create alias to package math
)

func main() {
	const PI = 3.1415
	var raio = 3.2 // The compiler will identify automatically what type variable to declared by the value

	// A variable declared should always be used
	area := PI * math.Pow(raio, 2)

	fmt.Println("This result is: ", area)

	// Declare block variables
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// Declare with reduce
	var e, f bool = false, true
	g, h, i := false, 1.3, "Hello..."

	fmt.Println(e, f)
	fmt.Println(g, h, i)
}
