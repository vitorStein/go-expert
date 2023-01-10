package main

import "fmt"

func main() {

	// Definindo slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Tamanho do slice
	fmt.Printf("Tamanho do slice: %d\n", len(slice))

	// Capacidade
	fmt.Printf("Capacidade do slice: %d\n", cap(slice))

	// "Cortando" os slice
	slice_2 := slice[0:5]
	fmt.Printf("Tamanho do slice 2: %d\n", len(slice_2))
	fmt.Printf("Capacidade do slice 2: %d\n", cap(slice_2))

	// Aumentando a capacidade do slice
	slice = append(slice, 11)
	fmt.Printf("Tamanho do slice: %d\n", len(slice))
	fmt.Printf("Capacidade do slice: %d\n", cap(slice))
	// Mesmo com 11 elementos o slice fica com capacidade de 20
	// Go redimensiona o slice automaticamente

}
