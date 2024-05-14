package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for {
		if N < 1 || N > 5000 {
			fmt.Println("Valor inválido")
			fmt.Scan(&N)
		} else {
			break
		}
	}
	vetor := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}
	for i := N - 1; i >= 0; i-- {
		fmt.Printf("%d ", vetor[i])
	}
}
