package main

import "fmt"

func main() {
	// Array is homogeneous(same type) and static (size fix)
	var note [3]float64
	fmt.Println(note)

	note[0], note[1], note[2] = 7.8, 4.3, 9.1
	fmt.Println(note)

	total := 0.0
	for i := 0; i < len(note); i++ {
		total += note[i]
	}

	avg := total / float64(len(note))
	fmt.Printf("Avarage: %.2f", avg)
}
