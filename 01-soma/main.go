package main

import "fmt"

func sumFn(n1, n2 int) int {
	return n1 + n2
}

func main() {
	var num1, num2 int
	fmt.Print("Digite um numero ")
	fmt.Scanln(&num1)

	fmt.Print("Digite outro numero ")
	fmt.Scanln(&num2)

	sum := sumFn(num1, num2)
	fmt.Printf("A soma de %d e %d é %d\n", num1, num2, sum)
}