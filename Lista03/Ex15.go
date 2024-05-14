package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var T int
	fmt.Scan(&T)

	if T <= 1 || T > 10 {
		fmt.Println("O número de testes precisa estar entre 2 e 10.")
		return
	}

	for t := 0; t < T; t++ {
		var n int
		fmt.Scan(&n)

		if n < 2 || n > 1000 {
			fmt.Println("O tamanho do vetor precisa estar entre 2 e 1000.")
			return
		}

		vetor := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}

		sort.Ints(vetor)

		menorDistancia := math.Abs(float64(vetor[1] - vetor[0]))
		for i := 2; i < n; i++ {
			distancia := math.Abs(float64(vetor[i] - vetor[i-1]))
			if distancia < menorDistancia {
				menorDistancia = distancia
			}
		}

		numComparacoes := n * (n - 1) / 2
		fmt.Printf("%d %d\n", int(menorDistancia), numComparacoes)
	}
}
