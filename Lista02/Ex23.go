package main

import "fmt"

func main() {
	somaDivisores := func(n int) int {
		soma := 0
		for i := 1; i <= n/2; i++ {
			if n%i == 0 {
				soma += i
			}
		}
		return soma
	}

	var N int
	fmt.Scan(&N)

	paresEncontrados := 0
	numero := 2

	for paresEncontrados < N {
		soma := somaDivisores(numero)
		if soma > numero && somaDivisores(soma) == numero {
			fmt.Printf("(%d,%d)\n", numero, soma)
			paresEncontrados++
		}
		numero++
	}
}
