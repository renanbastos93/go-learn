package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// Using `for` equals while
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// `For` traditional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("Verify pair or ")
	for i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Print(" pair ")
		} else {
			fmt.Print(" odd ")
		}
	}

	// Infinite Loop
	for {
		fmt.Println("Forever...")
		time.Sleep(time.Second) // time.Second -> 1s
	}

}