package main

func main() {

	// Definindo map
	maps := map[string]int{"xpto": 25, "xpta": 20}

	println(maps["xpto"])

	// Adicionando chaves
	maps["xptb"] = 15

	println(maps["xptb"])

	// Deletando chaves
	delete(maps, "xpto")

	// Podemos iniciar o map vazio
	map_vazio := map[string]int{}
	map_vazio["xpto"] = 10

	// Percorrendo map
	for key, value := range maps {
		println(key, value)
	}

	// No go temos os blank identifiers, como no python
	for _, value := range maps {
		println(value)
	}

}
