package main

type Livro struct {
	nome      string
	n_paginas int
	ano_pub   int
	autor     string
	ativo     bool
}

func main() {

	livro_1 := Livro{
		nome:  "Livro 1",
		ativo: true,
	}
	println(livro_1.nome)
	println(livro_1.ativo)

	// Podemos mudar os valores de uma struct
	livro_1.nome = "Livro 2"
	println(livro_1.nome)

	livro_1.Desativar()
	println(livro_1.ativo)

}

func (l Livro) Desativar() {
	l.ativo = false
	println("Desativando:", l.ativo)
}
