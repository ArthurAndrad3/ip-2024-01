package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)

	if N < 1 || N > 1000000 {
		return
	}

	freq := make(map[int]int)
	var frequencia, valor int

	for i := 0; i < N; i++ {
		var N int
		fmt.Scanln(&N)
		freq[N]++

		if freq[N] > frequencia || (freq[N] == frequencia && N < valor) {
			frequencia = freq[N]
			valor = N
		}
	}

	fmt.Println(valor)
	fmt.Println(frequencia)
}
