package main

import "fmt"

func main() {
	var salario, kw float64
	fmt.Printf("Digite o salario minimo: ")
	fmt.Scanln(&salario)
	fmt.Printf("Digite a quantidade de kw gasta: ")
	fmt.Scanln(&kw)
	valorKw := salario * 0.7 / 100
	valorConta := kw * valorKw
	valorContaDesconto := valorConta * 0.9
	fmt.Printf("Valor de cada kw: R$ %.2f\n", valorKw)
	fmt.Printf("Valor da conta: R$ %.2f\n", valorConta)
	fmt.Printf("Valor da conta com desconto: R$ %.2f\n", valorContaDesconto)
}
