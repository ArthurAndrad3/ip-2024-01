package main

import "fmt"

func main() {
	var N, K, x int
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
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}
	fmt.Scan(&K)
	for i := 0; i < N; i++ {
		if vetor[i] >= K {
			x++
		}
	}
	fmt.Printf("%d\n", x)

}
