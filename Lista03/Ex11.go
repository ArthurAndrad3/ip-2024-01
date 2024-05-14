package main

import "fmt"

func main() {
	var N, menor, maior int
	fmt.Scan(&N)
	for {
		if N < 1 || N > 1000 {
			fmt.Println("Valor inválido")
			fmt.Scan(&N)
		} else {
			break
		}
	}
	vetorV := make([]int, N)
	vetorW := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetorV[i])
		if vetorV[i] > maior {
			maior = vetorV[i]
		}
		if i == 0 {
			menor = vetorV[i]
		}
		if vetorV[i] < menor {
			menor = vetorV[i]
		}
	}
	for i := N - 1; i >= 0; i-- {
		vetorW[N-1-i] = vetorV[i]
	}
	fmt.Println(vetorV)
	fmt.Println(vetorW)
	fmt.Printf("%d\n%d\n", maior, menor)

}
