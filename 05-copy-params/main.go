package main

import "fmt"

func alterarCopia(x int) {
	fmt.Println("Recebido como:", x)
	x = x * 2
	fmt.Println("atualizado para:", x)
}

func main() {
	numero := 10
	fmt.Println("Definido como:", numero)
	alterarCopia(numero)
	fmt.Println("Fora da função:", numero)
}