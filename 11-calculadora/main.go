package main

import (
	"errors"
	"fmt"
)

type Calculadora struct {
	Operando1 float64
	Operando2 float64
}

func (c Calculadora) Dividir() (float64, error) {
	if c.Operando2 == 0 {
		return 0, errors.New("erro: divisão por zero não é permitida")
	}
	return c.Operando1 / c.Operando2, nil
}

func (c Calculadora) Somar() float64 {
	return c.Operando1 + c.Operando2
}

func (c Calculadora) Subtrair() float64	{
	return c.Operando1 - c.Operando2
}

func (c Calculadora) Multiplicar() float64 {
	return c.Operando1 * c.Operando2
}

type Runner struct {
	calculadora *Calculadora
	valor1	  float64
	valor2	  float64
	resultado float64
	operacao  string
}

func (r *Runner) SolicitarValores() error {
	fmt.Println("Digite o primeiro número: ")
	if _, err := fmt.Scanln(&r.calculadora.Operando1); err != nil {
		return errors.New("erro: entrada inválida para o primeiro número")
	}
	fmt.Println("Digite o segundo número")
	if _, err := fmt.Scanln(&r.calculadora.Operando2); err != nil {
		return errors.New("erro: entrada inválida para o segundo número")
	}
	return nil
}

func (r *Runner) SolicitarOperacao() (error) {
	var operacao string
	fmt.Println("Escolha a operação (+, -, *, /):")
	if _, err := fmt.Scanln(&operacao); err != nil {
		return errors.New("erro: entrada inválida para a operação")
	}
	switch operacao {
		case "+", "-", "*", "/":
			r.operacao = operacao
			return nil
		default:
			return errors.New("operação inválida")
	}
}

func (r *Runner) ExecutarOperacao() {
	switch r.operacao {
		case "+":
			r.resultado = r.calculadora.Somar()
		case "-":
			r.resultado = r.calculadora.Subtrair()
		case "*":
			r.resultado = r.calculadora.Multiplicar()
		case "/":
			resultado, err := r.calculadora.Dividir()
			if err != nil {
				fmt.Println("Erro: ", err)
			} else {
				fmt.Println("Resultado: ", resultado)
			}
	}
}

func (r *Runner) Execute() {
	for {
		if err := r.SolicitarValores(); err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		err := r.SolicitarOperacao()
		if err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		r.ExecutarOperacao()
	}
}

func NewRunner(c *Calculadora) *Runner {
	return &Runner(calculadora: calculadora)
	runner := NewRunner(calculadora)
	runner.Execute()
}