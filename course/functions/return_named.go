package main

import "fmt"

func change(p1 int, p2 int) (first int, second int) {
    first = p1
    second = p2
    return // return clean because already attr valur in variables
}


func main() {
    x, y := change(1, 2)
    fmt.Println(x, y)
}