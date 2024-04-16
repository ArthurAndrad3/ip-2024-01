package main

import "fmt"

func main() {
	var N int
	fmt.Printf("Digite um número inteiro: ")
	fmt.Scan(&N)
	if N%3 == 0 && N%5 == 0 {
		fmt.Printf("O número %d é divisível por 3 e por 5\n", N)
	} else {
		fmt.Printf("O número %d não é divisível por 3 e por 5\n", N)
	}
}
