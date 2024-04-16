package main

import "fmt"

func main() {
	var nota float64
	var conceito string
	fmt.Printf("Digite a nota: ")
	fmt.Scan(&nota)
	for {
		fmt.Printf("Digite a nota (entre 0 e 10): ")
		fmt.Scan(&nota)

		if nota >= 0 && nota <= 10 {
			break
		} else {
			fmt.Println("Nota inválida. Por favor, insira uma nota entre 0 e 10.")
		}
	}
	if nota >= 9 {
		conceito = "A"
	}
	if nota >= 7.5 && nota < 9 {
		conceito = "B"
	}
	if nota >= 6 && nota < 7.5 {
		conceito = "C"
	}
	if nota < 6 {
		conceito = "D"
	}
	fmt.Printf("Nota: %.2f\n", nota)
	fmt.Printf("Conceito: %s\n", conceito)
}
