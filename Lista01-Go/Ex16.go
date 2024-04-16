package main

import "fmt"

func main() {
	var salario float64
	fmt.Printf("Digite o salário: ")
	fmt.Scan(&salario)
	if salario <= 300 {
		salario += salario * 0.5
	} else {
		salario += salario * 0.3
	}
	fmt.Printf("Salário com reajuste: R$ %.2f\n", salario)
}
