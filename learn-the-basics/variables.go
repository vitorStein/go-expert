package main

// Constante
const a = "Hello, World!"

// Declaração de um única variável
var b bool

// Declarar várias variáveis de escopo global
var (
	d bool    // default = false
	e int     // default = 0
	f string  // default = ""
	g float64 // default = 0.0000...
)

func main() {

	println("Constante A:", a)

	println("B:", b)
	b = true
	println("B:", b)

	println("D:", d)
	println("E:", e)
	println("F:", f)
	println("G:", g)

	// Variáveis de escopo local
	// Declarando e atribuindo uma variável
	var h string = "Variável Local"
	println("H:", h)

	// Tudo no go é fortemente tipado
	// não podemos atribuir um inteiro a variável h que é do tipo string
	// ERRO: h = 2

	// var i string
	// Se uma variável é declarada em um escopo local e não é usada
	// o compilador exibe erro de "i declared but not used"
	// Quando o escopo é global não tem problema

	// Outra forma "abreviada" de declarar variáveis
	// O próprio compilador faz a inferência de tipos
	j := "Forma Inferida"
	println("J:", j)

	// Criando tipos
	type ID int

	var k ID
	println("K:", k)

}
