package main

import "fmt"

func usedDefer(n int) int {
    defer fmt.Println("The end")
    if n >= 50 {
        fmt.Println("GT 50")
        return 50
    } else {
        fmt.Println("LT 50")
        return n
    }
}

func main() {
    r1 := usedDefer(100)
    fmt.Println()
    r2 := usedDefer(30)
    fmt.Println()
    fmt.Println(r1)
    fmt.Println(r2)
}