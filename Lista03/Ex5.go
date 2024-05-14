package main

import "fmt"

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
		for i := 0; i < N; i++ {
			if vetor[i] > y {
				y = vetor[i]
				x = i
			}
		}
		fmt.Printf("%d %d \n", x, y)
		y = 0
	}
}
