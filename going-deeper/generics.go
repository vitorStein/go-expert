package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, value := range m {
		soma += value
	}

	return soma
}

func SomaFloat(m map[string]float32) float32 {
	var soma float32
	for _, value := range m {
		soma += value
	}

	return soma
}

// T pode ser inteiro ou float32
func SomaGenerics[T int | float32](m map[string]T) T {
	var soma T
	for _, value := range m {
		soma += value
	}

	return soma
}

type Number interface {
	int | float32
}

func SomaGenericsConstrain[T Number](m map[string]T) T {
	var soma T
	for _, value := range m {
		soma += value
	}

	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {

	// Realiza a soma dos valores
	// Somente pode ser usado para valores inteiros
	m := map[string]int{"Funcionario 1": 1000, "Funcionario 2": 2000, "Funcionario 3": 1500}
	println(SomaInteiro(m))

	// Para podermos somar com float, teríamos que criar outra função -> SomaFloat
	m2 := map[string]float32{"Funcionario 1": 1000.2, "Funcionario 2": 2000.3, "Funcionario 3": 1500.4}
	println(SomaFloat(m2))

	// Os generics entram justamente para contornar esse problema
	println(SomaGenerics(m))
	println(SomaGenerics(m2))

	// Constraints
	println(SomaGenericsConstrain(m))
	println(SomaGenericsConstrain(m2))

	// Utilizando comparable
	println(Compara(10, 100))
	println(Compara(10, 10))

}
