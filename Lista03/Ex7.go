package main

import (
	"fmt"
)

func main() {
	var N, x, y int
	for {
		fmt.Scan(&N)
		for {
			if N < 0 || N > 10000 {
				fmt.Println("Valor inválido")
				fmt.Scan(&N)
			} else {
				break
			}
		}
		if N == 0 {
			break
		}
		vetor := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&vetor[i])
		}
		y = 0
		for i := 0; i < N; i++ {
			if vetor[i] > y {
				y = vetor[i]
			}
		}
		x = 0
		for i := 0; i <= y; i++ {
			for j := 0; j < N; j++ {
				if i >= vetor[j] {
					x++
				}

			}
			fmt.Printf("(%d) %d\n ", i, x)
			x = 0

		}

	}
}
