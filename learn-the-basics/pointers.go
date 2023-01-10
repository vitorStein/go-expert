package main

func main() {

	// Memória -> Endereço -> Valor
	var a int = 10

	// Imprimindo variável a
	println(a)

	// Imprimindo endereço
	println(&a)

	// Criando ponteiro
	var ponteiro_para_a *int = &a
	println(ponteiro_para_a)

	// Se mudarmos o valor do ponteiro_para_a
	// mudamos o valor da variável a
	*ponteiro_para_a = 20
	println(a)

	// Variável b apontando para o valor de memória de a
	b := &a
	// Mostrando o valor que está no endereço da memória
	println(*b)

	// Exemplo de uso de ponteiros
	c := 10
	d := 30
	soma(c, d)
	println("C: ", c)
	println("D: ", d)

	// Se alterarmos o valor do parâmetro dentro da função
	// Não alteramos o valor da variável c e d
	// o Go cria passa o valor que está no endereço de memória de c e d
	soma_b(c, d)
	println("C_b: ", c)
	println("D_b: ", d)

	// Se quisermos mudar o valor de c e d pela função soma
	// podemos passar o endereço da memória para a função
	soma_c(&c, &d)
	println("C_c: ", c)
	println("D_c: ", d)
	// c e d foram alterados pela função soma_c

}

func soma(a, b int) int {

	return a + b
}

func soma_b(a, b int) int {

	a = 50
	return a + b
}

func soma_c(a, b *int) int {

	*a = 50
	*b = 50

	return *a + *b
}
