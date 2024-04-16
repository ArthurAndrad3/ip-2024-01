package main

import "fmt"

func main() {
	var A, B, C float64
	fmt.Printf("Digite o valor de A: ")
	fmt.Scan(&A)
	fmt.Printf("Digite o valor de B: ")
	fmt.Scan(&B)
	fmt.Printf("Digite o valor de C: ")
	fmt.Scan(&C)
	delta := B*B - 4*A*C
	fmt.Printf("O valor de delta é %.2f\n", delta)
}
