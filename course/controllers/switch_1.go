package main


import (
	"fmt"
)

func conceptNote(n float64) string {
	var note = int(n)

	switch note {
		case 10:
			fallthrough // next step of the switch
		case 9:
			return "A"
		case 8, 7:
			return "B"
		case 6, 5:
			return "C"
		case 4, 3:
			return "D"	
		case 2, 1, 0:
			return "F"
		default:
			return "Invalid note" 	
	}

}


func main() {
	fmt.Println(conceptNote(10))
	fmt.Println(conceptNote(9))
	fmt.Println(conceptNote(8))
	fmt.Println(conceptNote(5.4))
	fmt.Println(conceptNote(3))
	fmt.Println(conceptNote(1))
	fmt.Println(conceptNote(11))
}