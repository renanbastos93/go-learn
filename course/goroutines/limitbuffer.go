package main

import (
	"fmt"
	"time"
)

func init() {
	message := `
=====================================================
Using a channel with goroutine and this channel have
a limit buffer
=====================================================
`
	fmt.Println(message)
}

func routine(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println("Exec") // dont show message because gt buffer size
	c <- 5
	c <- 6
	c <- 7
}

func main() {
	ch := make(chan int, 3) // without define a buffer
	go routine(ch)
	
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}