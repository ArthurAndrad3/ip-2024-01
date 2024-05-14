package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	if N < 1 || N > 1000 {
		return
	}
	vetor := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}
	sort.Ints(vetor)
	for i := 0; i < N; i++ {
		if i == 0 {
			fmt.Printf("%d\n", vetor[i])
		} else {

			if vetor[i] != vetor[i-1] {
				fmt.Printf("%d\n", vetor[i])
			}
		}

	}
}
