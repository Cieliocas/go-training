package main

import "fmt"

func alterarReferencia(x *int) {
	*x = *x * 2
}

func main() {
	numero := 10
	fmt.Println("Definido como:", numero)
	alterarReferencia(&numero)
	fmt.Println("Fora da função:", numero)
}
