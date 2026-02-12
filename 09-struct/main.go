package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{"Lucas", 30}
	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Idade:", pessoa.Idade)
}