package main

import (
	"fmt"
)

func conceptOfNote(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	fmt.Println("Concept notes")
	fmt.Println(conceptOfNote(7.3))
	fmt.Println(conceptOfNote(2.1))
	fmt.Println(conceptOfNote(9.1))
	fmt.Println(conceptOfNote(8.1))
	fmt.Println(conceptOfNote(3.1))
}