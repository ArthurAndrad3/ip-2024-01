package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	if N <= 0 || N > 1000000 {
		return
	}
	vetor := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}
	sort.Ints(vetor)
	mediana := 0.0
	if N%2 == 0 {
		mediana = float64(vetor[N/2-1]+vetor[N/2]) / 2.0
	} else {
		mediana = float64(vetor[N/2])
	}

	fmt.Printf("%.2f\n", mediana)
}
