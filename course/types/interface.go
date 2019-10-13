package main

import (
	"fmt"
)

type print interface {
    toString() string
}

type people struct {
    name string
    lastName string
}

type product struct {
    name string
    price float64
}

func (p people) toString() string {
    return p.name + " " + p.lastName
}

func (p product) toString() string {
    return fmt.Sprintf("%s - %.2f", p.name, p.price)
}

func printer(x print) {
    fmt.Println(x.toString())
}

func main() {
    var thing print = people{"Renan", "Bastos"}
    fmt.Println(thing.toString())
    printer(thing)
    
    thing = product{"Pen", 0.97}
    fmt.Println(thing.toString())
    printer(thing)

    prod := product{"Trash", 99.33}
    printer(prod)
}