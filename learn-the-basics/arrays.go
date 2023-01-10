package main

import "fmt"

func main() {

	// Definindo array
	var array [3]int

	// Atribuindo valor no array
	array[0] = 1
	fmt.Printf("Na posição 0 temos o valor: %v\n", array[0])

	// Caso não tenha sido atribuído, pega valor default do tipo
	// int = 0
	fmt.Printf("Na posição 1 temos o valor: %v\n", array[1])

	// Tamanho do array
	fmt.Printf("Tamanho do array: %v\n", len(array))

	// Percorrendo o array
	for i, v := range array {
		fmt.Printf("Index: %v Valor: %v\n", i, v)
	}

}
