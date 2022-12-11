package main

import "fmt"

func main() {
	var valor1 float32
	var valor2 float32
	var operacao int
	
	fmt.Print("escolha a operação desejada:\n ")
	fmt.Print("escolha o numero 1 para SOMA \n")
	fmt.Print("escolha o numero 2 para SUBTRAÇÃO \n ")
	fmt.Print("escolha o numero 3 para MULTIPLICAÇÃO \n")
	fmt.Print("escolha o numero 4 para DIVISÃO \n")
	fmt.Scanln(&operacao)

	if operacao == 1 {
		fmt.Print("Digite valor1: ")
		fmt.Scanln(&valor1)
		fmt.Print("Digite valor2: ")
		fmt.Scanln(&valor2)
		fmt.Printf("resultado da SOMA é: %.2f \n", valor1 + valor2)
		}
	if operacao == 2 {
		fmt.Print("Digite valor1: ")
		fmt.Scanln(&valor1)
		fmt.Print("Digite valor2: ")
		fmt.Scanln(&valor2)
		fmt.Printf("resultado da SUBTRAÇÃO é: %.2f \n", valor1 - valor2)
		}
	if operacao == 3 {
		fmt.Print("Digite valor1: ")
		fmt.Scanln(&valor1)
		fmt.Print("Digite valor2: ")
		fmt.Scanln(&valor2)
		fmt.Printf("resultado da MULTIPLICAÇÃO é: %.2f \n", valor1 * valor2)
		}
	if operacao == 4 {
		fmt.Print("Digite valor1: ")
		fmt.Scanln(&valor1)
		fmt.Print("Digite valor2: ")
		fmt.Scanln(&valor2)
		fmt.Printf("resultado da DIVISÃO é: %.2f \n", valor1 / valor2)
		}
	
	
	
}