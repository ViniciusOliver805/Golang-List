package main

import "fmt"

func main() {
	var value1, value2, value3, value4, value5 float32
	var notes int

	fmt.Print("Please enter the number of notes you want averaged. Maximum 5: ")
	fmt.Scanln(&notes)
	if  notes == 5 {
		
		fmt.Print("First Note: ")
		fmt.Scanln(&value1)
		fmt.Print("Second Note: ")
		fmt.Scanln(&value2)
		fmt.Print("Third Note: ")
		fmt.Scanln(&value3)
		fmt.Print("Fourth Note: ")
		fmt.Scanln(&value4)
		fmt.Print("Fifth Note: ")
		fmt.Scanln(&value5)
		var sumofthevalues float32 = value1 + value2 + value3 + value4 + value5 
		var averageofvalues float32 = sumofthevalues / 5
		fmt.Printf("average: %.2f \n", averageofvalues)

	}
	if 	notes == 4 {

		fmt.Print("First Note: ")
		fmt.Scanln(&value1)
		fmt.Print("Second Note: ")
		fmt.Scanln(&value2)
		fmt.Print("Third Note: ")
		fmt.Scanln(&value3)
		fmt.Print("Fourth Note: ")
		fmt.Scanln(&value4)
		var sumofthevalues float32 = value1 + value2 + value3 + value4  
		var averageofvalues float32 = sumofthevalues / 4
		fmt.Printf("average: %.2f \n", averageofvalues)

	}
	if  notes == 3	{

		fmt.Print("First Note: ")
		fmt.Scanln(&value1)
		fmt.Print("Second Note: ")
		fmt.Scanln(&value2)
		fmt.Print("Third Note: ")
		fmt.Scanln(&value3)
		var sumofthevalues float32 = value1 + value2 + value3 
		var averageofvalues float32 = sumofthevalues / 3
		fmt.Printf("average: %.2f \n", averageofvalues)

	}
	if  notes == 2 {

		fmt.Print("First Note: ")
		fmt.Scanln(&value1)
		fmt.Print("Second Note: ")
		fmt.Scanln(&value2)
		var sumofthevalues float32 = value1 + value2 + value3 
		var averageofvalues float32 = sumofthevalues / 3
		fmt.Printf("Average: %.2f \n", averageofvalues)

	}

	
}