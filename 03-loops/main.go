package main

import "fmt"

func Main1() {
	for i:=1; i<=10; i++ {
		fmt.Printf("Esta é a iteração número %d\n", i)
	}
}

func main() {
	var age int
	for {
		fmt.Printf("Por favor, digite a sua idade: (obrigatorio ser maior de 18)\n")
		_, err := fmt.Scanln(&age)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, insira um número inteiro.")
			continue
		}

		if age < 18 { 
			fmt.Println("Você deve ser maior de idade para continuar.")
			continue
		}

		break
	}
	fmt.Println("Obrigado por fornecer sua idade.")
}
