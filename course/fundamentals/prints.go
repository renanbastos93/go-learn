package main

import "fmt"

func main() {
	// Method Print show message in same line
	fmt.Print("First Print...")
	fmt.Print("second Print...\n")

	// Method Println show message in a new line
	fmt.Println("First line")
	fmt.Println("A new line")

	// Concat value with message
	x := 3.14151515151515
	xs := fmt.Sprint(x) // Convert to string

	fmt.Println("The value X is " + xs)

	fmt.Println("The value X is ", x)
	fmt.Printf("The value X is %f", x)
	fmt.Printf("\nThe value X is %.2f", x)

	nInt := 1
	nFloat := 1.999
	vBool := false
	vString := "Something"
	// Print -> 1   1.999000   1.1   false   something
	fmt.Printf("\n%d %f %.1f %t %s", nInt, nFloat, nFloat, vBool, vString)
	fmt.Printf("\n%v %v %v %v", nInt, nFloat, vBool, vString)
}
