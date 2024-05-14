package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N < 2 || N > 10000 {
		fmt.Println("Número de pontos inválido")
		return
	}

	vetor1 := make([]float64, 3)
	fmt.Scan(&vetor1[0], &vetor1[1], &vetor1[2])

	vetor := make([]float64, N-1)
	for i := 0; i < N-1; i++ {
		vetor2 := make([]float64, 3)
		fmt.Scan(&vetor2[0], &vetor2[1], &vetor2[2])
		distancia := make([]float64, 3)
		for j := 0; j < 3; j++ {
			distancia[j] = math.Abs(vetor2[j] - vetor1[j])
		}
		sort.Float64s(distancia)
		fmt.Println(distancia)
		vetor1 = vetor2
		vetor[i] = distancia[2]
	}

	for i := 0; i < N-1; i++ {
		fmt.Printf("%.2f\n", vetor[i])
	}
}
