package main

import "fmt"

func main() {
	var N1, N2, N3 int
	fmt.Printf("Digite o numero 1: ")
	fmt.Scanln(&N1)
	fmt.Printf("Digite o numero 2: ")
	fmt.Scanln(&N2)
	fmt.Printf("Digite o numero 3: ")
	fmt.Scanln(&N3)
	N1 = N1 * 100
	N2 = N2 * 10
	concatenacao := N1 + N2 + N3
	quadrado := concatenacao * concatenacao
	if N1 < 0 || N2 < 0 || N3 < 0 || N1 > 900 || N2 > 90 || N3 > 9 {
		fmt.Println("Digito invalido")
	} else {
		fmt.Printf("Concatenacao: %d\n", concatenacao)
		fmt.Printf("Quadrado: %d\n", quadrado)
	}
}
