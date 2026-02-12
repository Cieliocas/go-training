package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
	Maior bool
}

func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá, meu nome é %s, tenho %d anos.", p.Nome, p.Idade)
}

func (p *Pessoa) FazerAniversario() {
	p.Idade++
}

func main() {
	pessoa := Pessoa{"Lucas", 30, true}
	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Idade:", pessoa.Idade)
	fmt.Println("Pessoa:", pessoa)
	fmt.Println("Pessoa &:", &pessoa)
	fmt.Println("Pessoa Nome &:", &pessoa.Nome)
	fmt.Println("Pessoa Idade &:", &pessoa.Idade)
	fmt.Println("Pessoa Maior &:", &pessoa.Maior)
	fmt.Println("Lucas diz:", pessoa.Apresentar())
	pessoa.FazerAniversario()
	fmt.Println("Lucas fez aniversário! Agora tem", pessoa.Idade)
}