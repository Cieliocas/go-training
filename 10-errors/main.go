package main

import "fmt"

type Calculadora struct {
	Operando1 float64
	Operando2 float64
}

func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, fmt.Errorf("erro: divisão por zero")
	}
	return c.Operando1 / c.Operando2, nil
}

func (c Calculadora) Somar() float64 {
	return c.Operando1 + c.Operando2
}

func (c Calculadora) Subtrair() float64 {
	return c.Operando1 - c.Operando2
}

func (c Calculadora) Multiplicar() float64 {
	return c.Operando1 * c.Operando2
}

func main() {
	var op1, op2 float64
	var operacao string
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&op1)
	fmt.Println("Digite o segundo número: ")
	fmt.Scanln(&op2)
	fmt.Println("Escolha a operação (+, -, *, /):")
	fmt.Scanln(&operacao)
	calc := Calculadora{op1, op2}
	switch operacao {
		case "+":
			fmt.Println("Resultado: ", calc.Somar())
		case "-":
			fmt.Println("Resultado: ", calc.Subtrair())
		case "*":
			fmt.Println("Resultado: ", calc.Multiplicar())
		case "/":
			resultado, err := calc.Dividir()
			if err != nil {
				fmt.Println("Erro: ", err)
			} else {
				fmt.Println("Resultado: ", resultado)
			}
		default:
			fmt.Println("Operação inválida!")	
	}
}