package main

import "fmt"

func main() {
	var n int
	var soma float64
	fmt.Printf("Digite o valor de n: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Printf("O valor de n deve ser maior que 0\n")
	} else {
		for i := 1; i <= n; i++ {
			soma += 1 / float64(i)
		}
		fmt.Printf("A soma dos %d primeiros termos da série harmônica é %f\n", n, soma)
	}
}
