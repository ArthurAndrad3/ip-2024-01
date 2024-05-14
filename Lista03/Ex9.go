package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N < 2 || N > 10000 {
		fmt.Println("Número de pontos inválido")
		return
	}

	var distancia float64
	vetor1 := make([]float64, 3)
	fmt.Scan(&vetor1[0], &vetor1[1], &vetor1[2])

	vetor := make([]float64, N-1)
	for i := 0; i < N-1; i++ {
		vetor2 := make([]float64, 3)
		fmt.Scan(&vetor2[0], &vetor2[1], &vetor2[2])

		distancia = math.Sqrt(math.Pow(vetor2[0]-vetor1[0], 2) + math.Pow(vetor2[1]-vetor1[1], 2) + math.Pow(vetor2[2]-vetor1[2], 2))
		vetor1 = vetor2
		vetor[i] = distancia
	}

	for i := 0; i < N-1; i++ {
		fmt.Printf("%.2f\n", vetor[i])
	}
}
