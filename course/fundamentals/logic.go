package main

import (
	"fmt"
)

/*
	j1, j2 := "Job1" , "Job2"
	If j1 and j2 jobs succeed -> buy tv 50"
	If j1 is different j2 -> buy tv 32"
	If j1 job or j2 job succeed -> buy ice cream"
*/
func purchase(j1, j2 bool) (bool, bool, bool) {
	buyTv50 := j1 && j2
	buyTv32 := j1 != j2
	buyIceCream := j1 || j2

	return buyTv50, buyTv32, buyIceCream
}

func main() {
	tv50, tv32, iceCream := purchase(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Ice Cream: %t, Healthy: %v", tv50, tv32, iceCream, !iceCream)
}