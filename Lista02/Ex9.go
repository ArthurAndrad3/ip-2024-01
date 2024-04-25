package main

import "fmt"

func main() {
	var N, i int
	fmt.Print("Digite um número: ")
	fmt.Scan(&N)
	if N == 2 {
		fmt.Println("PRIMO")
	} else {
		for i = 2; i < N; i++ {
			if N%i == 0 {
				fmt.Printf("NAO PRIMO\n")
				break
			}
		}
	}
	if i == N {
		fmt.Println("PRIMO")
	}
}
