package main

import "fmt"

type Pessoa struct {
	nome string
}

func (p Pessoa) andou() {
	p.nome += " Stein"
	fmt.Printf("A pessoa %v andou\n", p.nome)
}

func (p *Pessoa) andou_com_mudanca_nome() {
	p.nome += " Stein"
	fmt.Printf("A pessoa %v andou\n", p.nome)
}

func main() {

	victor := Pessoa{
		nome: "Victor",
	}

	victor.andou()
	fmt.Printf("%v\n", victor.nome)

	victor.andou_com_mudanca_nome()
	fmt.Printf("%v\n", victor.nome)

}
