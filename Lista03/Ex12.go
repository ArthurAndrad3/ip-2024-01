package main

import "fmt"

func main() {
	var N, maior int
	fmt.Scan(&N)
	for {
		if N < 1 || N > 1000 {
			fmt.Println("Valor inválido")
			fmt.Scan(&N)
		} else {
			break
		}
	}
	vetor := make([]int, N)
	vetorW := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
		if vetor[i] > maior {
			maior = vetor[i]
		}
	}
	for i := 0; i < N; i++ {
		vetorW[i] = maior
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i > 0 {
				if vetorW[i-1] < vetor[j] {
					if vetorW[i] > vetor[j] {
						vetorW[i] = vetor[j]
					}
				}
			} else {
				if vetorW[i] > vetor[j] {
					vetorW[i] = vetor[j]
				}
			}
		}

	}
	for i := 0; i < N; i++ {
		fmt.Printf("%d\n", vetorW[i])
	}
}
