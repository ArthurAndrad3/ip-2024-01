package main

import "fmt"

func main() {
	var N, k int
	fmt.Print("Digite o número de times no campeonato: ")
	fmt.Scan(&N)
	for i := 1; i < N; i++ {

		for j := i + 1; j <= N; j++ {
			k++
			fmt.Printf("Final %d: %d x %d\n", k, i, j)
		}
	}
}
