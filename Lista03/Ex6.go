package main

import "fmt"

func main() {
	var N, soma int
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
	for i := 0; i < N; i++ {
		soma += vetor[i]
	}
	fmt.Printf("%d\n", soma)

}
