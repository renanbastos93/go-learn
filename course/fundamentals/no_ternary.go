package main

import "fmt"

// Haven't operator ternary
func getResult(value float64) string {
	if value >= 6 {
		return "Win"
	}
	return "Loss"
}

func main() {
	fmt.Println(getResult(6.2))
	fmt.Println(getResult(3.2))
}