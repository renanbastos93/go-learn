package main

import (
	"fmt"
)

func main() {
	i := 1
	
	var p *int = nil
	p = &i // get the reference of memory of the `i` attribute to `p`
	*p++ // remove reference to get the value
	i++

	// Golang hasn't arithmetic of pointer
	// p++

	fmt.Println(p, *p, i, &i)
}