package main

import "fmt"

func MainPlus() {
	idade := 30
	ponteiroIdade := &idade

	fmt.Println("Valor da idade:", idade)
	fmt.Println("Endereço de memória da idade:", &idade)
	fmt.Println("Valor do ponteiro para idade:", ponteiroIdade)
	fmt.Println("Valor apontado pelo ponteiro:", *ponteiroIdade)
}

func main() {
	numero := 42
	ponteiro := &numero
	fmt.Println("Antes da alteração:", numero)
	*ponteiro = 99
	fmt.Println("Depois da alteração:", numero)
}