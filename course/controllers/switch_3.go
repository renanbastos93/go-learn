package main


import (
	"fmt"
	"time"
)

func typeVar(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Float/ Double"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "do not know"
	}
}

func main() {
	fmt.Println(typeVar(123))
	fmt.Println(typeVar(1.351))
	fmt.Println(typeVar("123"))
	fmt.Println(typeVar(func(){}))
	fmt.Println(typeVar(time.Now()))
}