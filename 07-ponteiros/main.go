package main

import "fmt"

func main() {
	idade := 30
	ponteiroIdade := &idade

	fmt.Println("Valor da idade:", idade)
	fmt.Println("Endereço de memória da idade:", &idade)
	fmt.Println("Valor do ponteiro para idade:", ponteiroIdade)
	fmt.Println("Valor apontado pelo ponteiro:", *ponteiroIdade)
}