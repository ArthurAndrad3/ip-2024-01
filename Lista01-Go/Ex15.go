package main

import "fmt"

func main() {
	var N int
	fmt.Printf("Digite um número inteiro: ")
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ^ 2 = %d\n", i, i*i)
		}
	}
}
