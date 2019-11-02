package main

import (
	"fmt"
	"time"
)

func init() {
	message := `==================================================
Using a channel with goroutine.
Channel is a way to communicate between GoRoutines
It is a type, e.g. type int, string, etc...
==================================================`
	fmt.Println(message)
}

func mult(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)
	go mult(2, ch)
	
	a, b, c := <-ch, <-ch, <-ch
	fmt.Println(a, b, c)
}