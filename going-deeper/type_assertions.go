package main

func main() {

	var a interface{} = "Interface Vazia"
	println(a)          // Imprime (0x45c500,0x47b860)
	println(a.(string)) // Imprime Interface Vazia

	// Verificar a o assertion deu certo
	res, ok := a.(int)
	println("RES:", res)
	println("OK:", ok)
	// Caso não funcione o ok vai retornar false

}
