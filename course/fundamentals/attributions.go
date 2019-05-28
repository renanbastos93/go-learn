package main

import (
	"fmt"
)

func main() {
	var b byte = 3 // byte is equal uint8
	fmt.Println(b)

	i := 3 // -> var i int = 3
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x // attribution new values
	fmt.Println(x, y)

}