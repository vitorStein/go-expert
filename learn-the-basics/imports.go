package main

import "fmt"

var a string = "Golang!"

func main() {

	// %T busca o tipo da variável
	fmt.Printf("A variável A é do tipo: %T\n", a)

	// %v retorna o valor da variável
	fmt.Printf("O valor da variável A é: %v\n", a)

	// %T com tipo criado
	type ID int
	var b ID = 1
	fmt.Printf("A variável B é do tipo: %T\n", b)

}
