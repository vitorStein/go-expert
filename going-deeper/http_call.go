package main

import (
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	// Defer deixa a execução do comando para o final da execução
	// usado para finalizar conexões de arquivos, bd, ...
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))
}
