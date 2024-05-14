package main

import "fmt"

func main() {
	var N, x int
	fmt.Scan(&N)
	if N > 5000 {
		return
	}
	vetor := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetor[i])
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if vetor[i] == vetor[j] && i != j {
				break
			}

			if j == N-1 && vetor[i] != vetor[j] && i != j {
				x++
			}
			if i == N-1 && j == N-1 {
				x++
			}

		}
	}
	fmt.Println(x)
}
