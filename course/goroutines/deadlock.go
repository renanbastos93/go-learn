package main

import (
	"fmt"
	"time"
)

func init() {
	message := `
=====================================================
Using a channel with goroutine and we go to simulate
a DeadLock that happens when already done all routines
=====================================================
`
	fmt.Println(message)
}

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 2
}

func main() {
	ch := make(chan int) // without define a buffer
	go routine(ch)
	
	// a := <-ch
	fmt.Println(<-ch)
	fmt.Println("was read")
	fmt.Println(<-ch) // DeadLock
}