package main

import "fmt"

func main() {

	var years int
	var days int
	var hours int
	var months int
	var minutes int 
	var seconds int

	fmt.Printf("Enter Here The Desired Number Of Years For Converting Minutes and Seconds: ")
	fmt.Scanf("%d", &years)

	months = years * 12
	days = years * 365
	hours = days * 24
	minutes = hours * 60
	seconds = minutes * 60

	fmt.Printf("\n%d Conversion to \n%d min, \n%d sec.\n",years, minutes, seconds, months)
}