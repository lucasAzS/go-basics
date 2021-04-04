package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) setNome(nome string){
	p.Nome = nome
	fmt.Println(p.Nome)
}

func main() {
	pessoa := Pessoa{
		Nome: "Lucas",
		Idade: 2,
	}

	pessoa.setNome("Otávio")
}