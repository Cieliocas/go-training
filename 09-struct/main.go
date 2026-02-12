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
	fmt.Println("Pessoa:", pessoa)
	fmt.Println("Pessoa &:", &pessoa)
	fmt.Println("Pessoa Nome &:", &pessoa.Nome)
	fmt.Println("Pessoa Idade &:", &pessoa.Idade)

}