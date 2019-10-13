package main

import (
	"fmt"
)

type sportive interface {
    turboOn()
}

type ferrari struct {
    model string
    turboOn2 bool
    currentVelocity int
}

func (f *ferrari) turboOn() {
    f.turboOn2 = true
}

func main() {
    ferrari1 := ferrari{
        "F 40",
        false,
        0,
    }
    ferrari1.turboOn()

    var ferrari2 = &ferrari{"F 40", false, 0}
    ferrari2.turboOn()

    
    fmt.Println(ferrari1, ferrari2)
}