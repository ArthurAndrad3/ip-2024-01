package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)

	mmc := 1
	primo := 2

	// Função para verificar se um número é primo
	ehPrimo := func(num int) bool {
		if num <= 1 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	for n1 > 1 || n2 > 1 || n3 > 1 {
		if ehPrimo(primo) {
			for n1%primo == 0 || n2%primo == 0 || n3%primo == 0 {
				mmc *= primo
				fmt.Printf("%d %d %d: %d\n", n1, n2, n3, primo)
				if n1%primo == 0 {
					n1 /= primo
				}
				if n2%primo == 0 {
					n2 /= primo
				}
				if n3%primo == 0 {
					n3 /= primo
				}

			}
		}
		primo++
	}

	fmt.Printf("MMC: %d\n", mmc)
}
