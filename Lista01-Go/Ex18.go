package main

import "fmt"

func main() {
	var a1, r, n, soma int
	fmt.Printf("Digite o valor do primeiro termo da PA: ")
	fmt.Scan(&a1)
	fmt.Printf("Digite o valor da razão da PA: ")
	fmt.Scan(&r)
	fmt.Printf("Digite o valor de n: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		soma += a1 + i*r
	}
	fmt.Printf("A soma dos %d primeiros termos da PA é %d\n", n, soma)

}
