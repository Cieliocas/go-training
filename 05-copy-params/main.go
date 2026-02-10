package main

import "fmt"

func alterarCopia(x int ) {
	x = x * 2
}

func main() {
	numero := 10
	fmt.Println("Definido como:", numero)
	alterarCopia(numero)
	fmt.Println("Fora da função:", numero)
}