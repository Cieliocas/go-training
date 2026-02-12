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

func main() {
	calc := Calculadora{10, 5}
	resultado, err := calc.Dividir()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Resultado: ", resultado)
}