package main

import "fmt"

func fatorial(n int) (int, error) {
    switch {
	case n < 0:
		return -1, fmt.Errorf(" number is invalid to calc: %d", n)
	case n == 0:
		return 1, nil
	default:
		lastN, _ := fatorial(n - 1)
		return n * lastN, nil
	}
}

func fatImproved(n uint) (uint, error) {
    if n == 0 {
        return 1, nil
    }
    lastN, _ := fatImproved(n - 1)
    return n * lastN, nil
}

func main() {
    fat, _ := fatorial(5)
    fmt.Println(fat)

    _, err := fatorial(-7)
    if err != nil {
        fmt.Println(err)
    }

    // a better solution
    resImproved, _ := fatImproved(5)
    fmt.Println(resImproved)
}