package main

import "fmt"

func printGraduates(graduate ...string) {
    for i, g := range graduate {
        fmt.Printf("%d) %s\n", i+1, g)
    }
}

func main() {
    graduates := []string{"Renan", "Maria", "Jose"}
    printGraduates(graduates...)
}