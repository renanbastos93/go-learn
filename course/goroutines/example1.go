package main

import (
	"fmt"
	"time"
)

func speak(people, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interetor %d)\n", people, text, i+1)
	}
}

func main() {
	// speak("Maria", "Why you do not speak to me?", 3)
	// speak("Joao", "I just speak after you", 1)

	// go speak("Maria", "Ei...", 500)
	// go speak("Joao", "Ops...", 500)
	// time.Sleep(500 * time.Second)
	// fmt.Println("End")

	go speak("Maria", "I understood", 10)
	speak("Joao", "Congrats", 5)
}