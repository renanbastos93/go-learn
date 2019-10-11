package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func exec(fun func(int, int) int, p1, p2 int) int {
	return fun(p1, p2)
}

func main() {
    fmt.Println(exec(mult, 7, 5))
}