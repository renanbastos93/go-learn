package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numRandom() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numRandom(); i > 5 { // supported also with switch
		fmt.Println("Rand:", i)
	} else {
		fmt.Println("Not exist rand:", nil)
	}
}