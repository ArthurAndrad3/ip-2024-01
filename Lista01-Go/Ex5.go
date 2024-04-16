package main

import (
	"fmt"
)

func main() {
	var gasto, Vconta float64
	var conta int
	var consumidor string
	fmt.Printf("Digite o valor da conta: ")
	fmt.Scanln(&conta)
	fmt.Printf("Digite o consumo de água: ")
	fmt.Scanln(&gasto)
	fmt.Printf("Digite a inicial do tipo de consumidor: ")
	fmt.Scanln(&consumidor)
	if consumidor == "R" {
		Vconta = 5 + gasto*0.05
	}
	if consumidor == "C" {
		Vconta = 500 + (gasto-80)*0.25
	}
	if consumidor == "I" {
		Vconta = 800 + (gasto-100)*0.04
	}
	fmt.Printf("Conta: %d\n", conta)
	fmt.Printf("Valor da conta: %.2f\n", Vconta)
}
