package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, y = 2.4, 2
	fmt.Println(x / float64(y))

	value := 6.9
	valueFinesh := int(value)
	fmt.Println(valueFinesh)

	// Here it use table unicode string(97) == 'a'
	fmt.Println("Trying convert ", string(97))

	// int to string
	fmt.Println("Convert int to string", strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// 1 == true
	// != 1 is equal false
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Was converted to", b)
	}
}