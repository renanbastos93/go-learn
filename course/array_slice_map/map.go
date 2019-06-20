package main

import (
	"fmt"
)


func main() {
	// Maps must be initialized
	// var approved map[int]string
	approved := make(map[int]string)

	approved[12313213] = "Maria"
	approved[23423424] = "Pedro"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approved[23423424])
	delete(approved, 23423424)
	fmt.Println(approved[23423424])


	// Map2
	funcSalary := map[string]float64{
		"Renan Bastos": 10000.01,
		"Fulano Bastos": 5000.91,
		"Beltrano Bastos": 7800.38,
	}

	funcSalary["Other people"] = 1200.32
	delete(funcSalary, "Not Exists")

	for name, salary := range funcSalary {
		fmt.Println(name, salary)
	}

	
	// Map Nested
	funcByLetter := map[string]map[string]float64{
		"G": {
			"Renan": 1234.66,
			"Bastos": 2639.09,
		},
		"Q": {
			"Maria": 9887.66,
			"Joana": 150036.09,
		},
	}

	delete(funcByLetter, "Q")

	for letter, funcs := range funcByLetter {
		fmt.Println(letter, funcs)
		for name, salary := range funcs {
			fmt.Println(name, salary)
		}
	}

}
