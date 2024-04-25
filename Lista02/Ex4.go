package main

import "fmt"

func main() {
	var n, i, K, s float64
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)
	fmt.Printf("Digite um número inicial: ")
	fmt.Scan(&i)
	x := i
	fmt.Printf("Digite o número de operações: ")
	fmt.Scan(&K)
	fmt.Printf("Digite o valor a ser somado: ")
	fmt.Scan(&s)
	fmt.Printf("Tabuada de soma:\n")
	for j := 1; j <= int(K); j++ {
		fmt.Printf("%.2f + %.2f = %.2f\n", n, x, n+x)
		x += s
	}
	fmt.Printf("Tabuada de subtração:\n")
	x = i
	for j := 1; j <= int(K); j++ {
		fmt.Printf("%.2f - %.2f = %.2f\n", n, x, n-x)
		x += s
	}
	fmt.Printf("Tabuada de multiplicação:\n")
	x = i
	for j := 1; j <= int(K); j++ {
		fmt.Printf("%.2f * %.2f = %.2f\n", n, x, n*x)
		x += s
	}
	fmt.Printf("Tabuada de divisão:\n")
	x = i
	for j := 1; j <= int(K); j++ {
		fmt.Printf("%.2f / %.2f = %.2f\n", n, x, n/x)
		x += s
	}
}
