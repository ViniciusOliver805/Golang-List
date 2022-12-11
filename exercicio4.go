package main

import "fmt"

func main() {
	var note1 float32 
	var note2 float32

	
	fmt.Print(" please enter the first note:	")
	fmt.Scanln(&note1)
	fmt.Print("please enter the second note: ")
	fmt.Scanln(&note2)
	fmt.Println(" ")
	var sumofthenotes float32 = note1 + note2


	fmt.Printf("the sum of the two notes are: %.2f \n", sumofthenotes)
}