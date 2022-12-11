package main

import "fmt"

func main() {

	var seconds int
	var minutes int
	var sec int
	var sobra int

	fmt.Printf("Digite aqui a quantidade de segundos desejada para a conversao: ")
	fmt.Scanf("%d", &seconds)

	sobra = seconds % 3600
	minutes = seconds / 60
	sec = sobra % 60
	fmt.Printf("%d segundos desejados para conversao: = %d minuto e %d segundo\n", seconds, minutes, sec)
}