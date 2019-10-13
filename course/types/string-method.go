package main

import (
	"fmt"
    str "strings"
)

type people struct {
    name string
    lastName string
}

func (p people) getFullName() string {
    return p.name + " " + p.lastName 
}

func (p *people) setFullName(fullName string) {
    sliceName := str.Split(fullName, " ")
    p.name = sliceName[0]
    p.lastName = sliceName[1]
}

func main() {
    p1 := people{"Renan", "Bastos"}
    fmt.Println(p1.getFullName())

    p1.setFullName("Elvis Bastos")
    fmt.Println(p1.getFullName())
}