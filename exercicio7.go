package main

import "fmt"

func looping(media, maior int) (int, int) {
	var num int
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		if num > maior {
			maior = num
		}
		media += num
	}
	media = media / 10
	return media, maior
}

func main() {
	var media, maior int
	fmt.Println("Por favor, digite aqui os 10 valores")
	media, maior = looping(media, maior)
	fmt.Println("a média dos valores é:", media)
	fmt.Println("o maior numero entres os dez é:", maior)
}