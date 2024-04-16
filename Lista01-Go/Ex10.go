package main

import "fmt"

func main() {
	var A, B, C, D float64
	fmt.Printf("Digite o valor de A: ")
	fmt.Scan(&A)
	fmt.Printf("Digite o valor de B: ")
	fmt.Scan(&B)
	fmt.Printf("Digite o valor de C: ")
	fmt.Scan(&C)
	fmt.Printf("Digite o valor de D: ")
	fmt.Scan(&D)
	determinante := A*D - B*C
	fmt.Printf("O valor do determinante é %.2f\n", determinante)
}
