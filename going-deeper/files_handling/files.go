package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Criando um arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escrevendo no arquivo
	tamanho, err := f.WriteString("Hello World!")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho do arquivo: %d bytes \n", tamanho)
	f.Close()

	// Lendo arquivo
	// Lê o arquivo completo
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// Para leitura gradual do arquivo
	arquivo_2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo_2)
	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

}
