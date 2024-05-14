package main

import "fmt"

func main() {
	var N, x, y, z int
	fmt.Scan(&N)
	vetor := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
		if vetor[i] > y {
			y = vetor[i]
			z = i

		}
	}
	for i := 0; i < N; i++ {
		if vetor[i] == vetor[N-1] {
			x++
		}
	}
	fmt.Printf("Nota %d, %d vezes\nNota %d, indice %d\n", vetor[N-1], x, y, z)
}
