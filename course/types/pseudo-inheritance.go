package main

import (
	"fmt"
)

// PSEUDO inheritance STRUCTS
// struct-inheritance.go

type car struct {
    name string
    currentVelocity int
}

type ferrari struct {
    car
    turboOn bool
}

func main() {
    f := ferrari{}
    f.name = "F-40"
    f.currentVelocity = 0
    f.turboOn = true
    fmt.Println("Ferrari: ", f.name)
    fmt.Println("Turbo On? ", f.turboOn)
}