package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	for i, numero := range numeros {
		fmt.Printf("Índice: %d, Valor: %d, Endereço: %p\n", i, numero, &numeros[i])
	}
}