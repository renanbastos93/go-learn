package main

import "fmt"

func inc1(n int) {
    n++
}

func inc2(n *int) {
    *n++
}

func main() {
    n := 1

    inc1(n)
    fmt.Println(n)

    // & is used to get reference of the variable
    inc2(&n)
    fmt.Println(n)
}