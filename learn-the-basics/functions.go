package main

import "errors"

func main() {

	println(sum(1, 2))
	println(sum_a(1, 2))
	println(sum_b(1, 2))
	println(sum_c(1, 2))
	println(sum_d(1, 2, 3, 4, 5, 6, 7))
	println(sum_e(6, 7))

}

// Diferentes formas de declarar funções

func sum(a int, b int) int {
	return a + b
}

func sum_a(a, b int) int {
	return a + b
}

// Funções no Go podem retornar mais de um valor
func sum_b(a, b int) (int, bool) {
	return a + b, true
}

// Isso é muito usado em erros
func sum_c(a, b int) (int, error) {
	return a + b, errors.New("Exemplo de erro que pode ser retornado")
}

// Funções variádicas
// Não sei a quantidade de parâmetros que a função irá receber
func sum_d(a ...int) int {
	total := 0

	for _, values := range a {
		total += values
	}

	return total
}

// Funções anonimas ou closures
func sum_e(a, b int) int {

	vezes_2 := func() int {
		return sum_d(a, b) * 2
	}()

	return vezes_2
}
