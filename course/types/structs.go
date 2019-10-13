package main

import "fmt"

type product struct {
    name string
    price float64
    discount float64
}

func (p product) discountPrice() float64 {
    return p.price * (1 - p.discount)
}

func main() {
    p := product {
        name: "Pen",
        price: 1.79,
        discount: 0.75,
    }
    finalPrice := p.discountPrice()
    fmt.Println(p, finalPrice)
}