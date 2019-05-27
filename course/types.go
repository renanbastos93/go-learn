package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Numerics int
	fmt.Println(1, 2, 1000)
	fmt.Println("Is int ", reflect.TypeOf(3600))

	// Not signal is positive uint8 uint16 uint32 uint64
	var b byte = 255 // Alias to `uint8`
	fmt.Println("The byte is ", reflect.TypeOf(b))

	// with signal uint8 uint16 uint32 uint64
	firstInt := math.MaxInt64
	fmt.Println("This value max is ", firstInt)
	fmt.Println("The value has type ", reflect.TypeOf(firstInt))

	// By pattern is reference a unicode table return an uint32
	var someRune rune = 'a'
	fmt.Println("What's type? is ", reflect.TypeOf(someRune))
	fmt.Println(someRune)

	// Is normally when use floating point is `float32`
	var fl32 float32 = 49.99
	fmt.Println("Type of fl32 is ", reflect.TypeOf(fl32))
	fmt.Println("Type fl32 is equall 49.99 -> ", reflect.TypeOf(49.99))

	// Boolean
	var isBool bool = true
	fmt.Println("It is a type boolean? is ", reflect.TypeOf(isBool))
	fmt.Println(!isBool)

	// String's
	var someStr string = "Hi, my names is Renan"
	fmt.Println(someStr + "!")
	fmt.Println("Size of", len(someStr))

	otherStr := `Hi
	My name is
	Renan
	`
	fmt.Println("Size of another str", len(otherStr))

	// Char is equal int32
	// var c char = 'a' -> not exist in GO
	char := 'a'
	fmt.Println("Type of char:", reflect.TypeOf(char))
	fmt.Println(char)
}
