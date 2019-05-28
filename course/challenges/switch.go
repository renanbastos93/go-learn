/*
 * Change code this file ./controllers/ifelseif.go
 * For using the switch
*/

package main

import (
	"fmt"
)

func conceptOfNote(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
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