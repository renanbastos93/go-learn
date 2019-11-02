package main

import (
	"fmt"
)

func init() {
	fmt.Println("Class channels")
}

func main() {
	ch := make(chan int, 1)
	ch <- 7 // Send data to this channel
	valChannel := <-ch // Receive value this channel
	fmt.Println(valChannel) 
}