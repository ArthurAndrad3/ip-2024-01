package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, y, x int
	fmt.Scan(&N)
	vetor := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}
	sort.Ints(vetor)
	fmt.Println(vetor)
	for i := 0; i < N; i++ {
		if vetor[i]%2 == 0 {
			vetor[y] = vetor[i]
			y++
		}
		if vetor[i]%2 != 0 {
			vetor[N-1-x] = vetor[i]
			x++
		}
	}
	fmt.Println(vetor)
	for i := 0; i < N; i++ {
		fmt.Printf("%d\n", vetor[i])
	}
}
