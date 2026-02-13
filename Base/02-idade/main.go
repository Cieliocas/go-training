package main

import "fmt"

func main() {
	var age int
	fmt.Print("Digite a sua idade: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	
	}
	fmt.Printf("A idade digitada é %d\n", age)
}