package main

import "fmt"

func avg(n ...float64) float64 {
    total := 0.0
    for _, num := range n {
        total += num
    }
    return total / float64(len(n))
}

func main() {
    fmt.Printf("Avg: %.2f", avg(1.2, 3.7, 5.9, 9.9))
}