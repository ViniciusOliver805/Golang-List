package main

import "fmt"

func main() {
	var valor1 float32
	var valor2 float32
	var valor3 float32
	var total float32

	fmt.Print("PRIMEIRO valor: ")
	fmt.Scanln(&valor1)
	fmt.Print("SEGUNDO valor: ")
	fmt.Scanln(&valor2)
	fmt.Print("TERCEIRO valor: ")
	fmt.Scanln(&valor3)
	fmt.Println(" ")
	total = valor1 * valor2 * valor3
	
	fmt.Printf("TOTAL DA MULTIPLICAÇÃO ENTRE OS 3 NÚMEROS: %.2f \n", total)
}