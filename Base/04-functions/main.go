package main

import "fmt"

func dobrar(numero int) (int, int, error) {
	if numero < 0 {
		return 0, 0, fmt.Errorf("O número fornecido %d é negativo", numero)
	}
	return numero * 2, numero * 2, nil
}

func main() {
	resultado, _, err := dobrar(10)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
