package main

import (
	"fmt"
)

type course struct {
    name string
}

func main() {
    var thing interface{} // equals ANY
    fmt.Println(thing)

    thing = 3
    fmt.Println(thing)

    thing = "New course"
    fmt.Println(thing)

    thing = 1.3
    fmt.Println(thing)

    thing = true
    fmt.Println(thing)

    type dynamic interface{}

    var thing2 dynamic = "Wow"
    fmt.Println(thing2)

    thing2 = course{"Golang course"}
    fmt.Println(thing2)

}