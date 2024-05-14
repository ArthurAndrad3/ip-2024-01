package main

import "fmt"

func busca(vetor []int, x int) int {
	for i := 0; i < len(vetor); i++ {
		if x == vetor[i] {
			return i
		}
	}
	return -1
}

func main() {
	var N int
	fmt.Scan(&N)
	vetor := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}

	var Bu int
	fmt.Scan(&Bu)

	posicao := busca(vetor, Bu)
	fmt.Print(posicao)
}
