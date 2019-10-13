package main

import (
	"fmt"
)

type sportive interface {
    turboOn()
}

type lux interface {
    makeBaliza()
}

type sportiveLux interface {
    sportive
    lux
    // you can add more methods
}

type BMW7 struct {}


func (b BMW7) turboOn() {
    fmt.Println("Turbo...")
}

func (b BMW7) makeBaliza() {
    fmt.Println("Baliza...")
}

func main() {
   var b sportiveLux = BMW7{}
   b.turboOn()
   b.makeBaliza()
}