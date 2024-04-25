package main

import "fmt"

func main() {
	var n, fatorial float64
	fmt.Printf("Digite um número: ")
	fmt.Scan(&n)
	fatorial = n
	if n == 0 {
		fmt.Printf("0! = 1\n")
	} else {
		for i := 1; i < int(n); i++ {
			fatorial *= (n - float64(i))
		}
		fmt.Printf("%.0f! = %.0f", n, fatorial)
	}

}
