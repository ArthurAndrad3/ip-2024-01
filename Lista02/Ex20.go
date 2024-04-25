package main

import "fmt"

func main() {
	var N, soma int
	fmt.Scanln(&N)
	fmt.Printf("%d = ", N)
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			if i != 1 && i != N {
				fmt.Printf("+ ")
			}
			if i != N {
				soma += i
				fmt.Printf("%d ", i)
			}
		}

	}
	fmt.Printf("= %d ", soma)
	if soma == N {
		fmt.Println("(NUMERO PERFEITO)")
	} else {
		fmt.Println("(NUMERO NAO E PERFEITO)")
	}
}
