package main

import (
	"fmt"
)

type note float64

func (n  note) entry(start, end float64) bool {
    return float64(n) >= start && float64(n) <= end
}

func noteByConcept(n note) string {
    switch {
    case n.entry(9.0, 10.0):
        return "A"
    case n.entry(7.0, 8.99):
        return "B"
    case n.entry(5.0, 7.99):
        return "C"
    case n.entry(3.0, 4.99):
        return "D"
    default:
        return "E"
    }
}

func main() {
    fmt.Println(noteByConcept(9.8))
    fmt.Println(noteByConcept(6.8))
    fmt.Println(noteByConcept(1.8))
}