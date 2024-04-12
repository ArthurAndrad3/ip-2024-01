package main

import (
	"fmt"
)

func main() {
	var nota1, nota2, nota3 float64

	fmt.Println("Digite a nota 1:")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a nota 2:")
	fmt.Scanln(&nota2)
	fmt.Println("Digite a nota 3:")
	fmt.Scanln(&nota3)

	media := (nota1 + nota2 + nota3) / 3

	fmt.Printf("Média: %.2f\n", media)

	if media > 6 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
