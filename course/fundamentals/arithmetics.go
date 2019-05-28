package main

import (
	"fmt"
	"math"
)

func main() {
	n1, n2 := 3, 2

	fmt.Println("Summation ->", n1 + n2)
	fmt.Println("Subtraction ->", n1 - n2)
	fmt.Println("Division ->", n1 / n2)
	fmt.Println("Multiplication ->", n1 * n2)
	fmt.Println("Module ->", n1 % n2)

	// Bitwise
	fmt.Println("And ->", n1 & n2) // 11 & 10 = 10 = 2
	fmt.Println("Or ->", n1 | n2) // 11 | 10 = 11 = 3
	fmt.Println("Xor ->", n1 ^ n2) // 11 ^ 10 = 10 = 01

	n3, n4 := 3.0, 2.0

	// Other operations using math...
	fmt.Println("Major ->", math.Max(float64(n1), float64(n2))  )
	fmt.Println("Minor ->", math.Min(n3, n4))
	fmt.Println("Exponential ->", math.Pow(n3, n4))
}